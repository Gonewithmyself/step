package mem

import (
	"fmt"
	"unsafe"
)

func (p *pageAlloc) scavengeStartGen() {
	p.inUse.cloneInto(&p.scav.inUse)

	// Pick the new starting address for the scavenger cycle.
	var startAddr offAddr
	if p.scav.scavLWM.lessThan(p.scav.freeHWM) {
		// The "free" high watermark exceeds the "scavenged" low watermark,
		// so there are free scavengable pages in parts of the address space
		// that the scavenger already searched, the high watermark being the
		// highest one. Pick that as our new starting point to ensure we
		// see those pages.
		startAddr = p.scav.freeHWM
	} else {
		// The "free" high watermark does not exceed the "scavenged" low
		// watermark. This means the allocator didn't free any memory in
		// the range we scavenged last cycle, so we might as well continue
		// scavenging from where we were.
		startAddr = p.scav.scavLWM
	}
	p.scav.inUse.removeGreaterEqual(startAddr.addr())

	// reservationBytes may be zero if p.inUse.totalBytes is small, or if
	// scavengeReservationShards is large. This case is fine as the scavenger
	// will simply be turned off, but it does mean that scavengeReservationShards,
	// in concert with pallocChunkBytes, dictates the minimum heap size at which
	// the scavenger triggers. In practice this minimum is generally less than an
	// arena in size, so virtually every heap has the scavenger on.
	p.scav.reservationBytes = alignUp(p.inUse.totalBytes, pallocChunkBytes) / scavengeReservationShards
	p.scav.gen++
	p.scav.released = 0
	p.scav.freeHWM = minOffAddr
	p.scav.scavLWM = maxOffAddr
}

func (p *pageAlloc) free(base, npages uintptr) {

	// If we're freeing pages below the p.searchAddr, update searchAddr.
	if b := (offAddr{base}); b.lessThan(p.searchAddr) {
		p.searchAddr = b
	}
	// Update the free high watermark for the scavenger.
	limit := base + npages*pageSize - 1
	if offLimit := (offAddr{limit}); p.scav.freeHWM.lessThan(offLimit) {
		p.scav.freeHWM = offLimit
	}
	if npages == 1 {
		// Fast path: we're clearing a single bit, and we know exactly
		// where it is, so mark it directly.
		i := chunkIndex(base)
		chunk := p.chunkOf(i)
		chunk.free1(chunkPageIndex(base))
		fmt.Println("free chunk", i)
	} else {
		// Slow path: we're clearing more bits so we may need to iterate.
		sc, ec := chunkIndex(base), chunkIndex(limit)
		si, ei := chunkPageIndex(base), chunkPageIndex(limit)

		if sc == ec {
			// The range doesn't cross any chunk boundaries.
			p.chunkOf(sc).free(si, ei+1-si)
		} else {
			// The range crosses at least one chunk boundary.
			p.chunkOf(sc).free(si, pallocChunkPages-si)
			for c := sc + 1; c < ec; c++ {
				p.chunkOf(c).freeAll()
			}
			p.chunkOf(ec).free(0, ei+1)
		}
	}
	p.update(base, npages, true, false)

	p.scavengeStartGen()
}

func (b *pallocBits) free1(i uint) {
	(*pageBits)(b).clear(i)
}

// free frees the range [i, i+n) of pages in the pallocBits.
func (b *pallocBits) free(i, n uint) {
	(*pageBits)(b).clearRange(i, n)
}

// freeAll frees all the bits of b.
func (b *pallocBits) freeAll() {
	(*pageBits)(b).clearAll()
}

func (p *pageAlloc) scavenge(nbytes uintptr, mayUnlock bool) uintptr {

	var (
		addrs addrRange
		gen   uint32
	)
	released := uintptr(0)
	for released < nbytes {
		if addrs.size() == 0 {
			if addrs, gen = p.scavengeReserve(); addrs.size() == 0 {
				break
			}
		}
		r, a := p.scavengeOne(addrs, nbytes-released, mayUnlock)
		released += r
		addrs = a
	}
	// Only unreserve the space which hasn't been scavenged or searched
	// to ensure we always make progress.
	p.scavengeUnreserve(addrs, gen)
	return released
}

func (p *pageAlloc) scavengeUnreserve(r addrRange, gen uint32) {

	if r.size() == 0 || gen != p.scav.gen {
		return
	}
	if r.base.addr()%pallocChunkBytes != 0 {
		throw("unreserving unaligned region")
	}
	p.scav.inUse.add(r)
}

func (p *pageAlloc) scavengeReserve() (addrRange, uint32) {

	// Start by reserving the minimum.
	r := p.scav.inUse.removeLast(p.scav.reservationBytes)

	// Return early if the size is zero; we don't want to use
	// the bogus address below.
	if r.size() == 0 {
		return r, p.scav.gen
	}

	// The scavenger requires that base be aligned to a
	// palloc chunk because that's the unit of operation for
	// the scavenger, so align down, potentially extending
	// the range.
	newBase := alignDown(r.base.addr(), pallocChunkBytes)

	// Remove from inUse however much extra we just pulled out.
	p.scav.inUse.removeGreaterEqual(newBase)
	r.base = offAddr{newBase}
	return r, p.scav.gen
}

func Loadp(ptr unsafe.Pointer) unsafe.Pointer {
	return *(*unsafe.Pointer)(ptr)
}

func (m *pallocData) hasScavengeCandidate(min uintptr) bool {
	if min&(min-1) != 0 || min == 0 {
		print("runtime: min = ", min, "\n")
		throw("min must be a non-zero power of 2")
	} else if min > maxPagesPerPhysPage {
		print("runtime: min = ", min, "\n")
		throw("min too large")
	}

	// The goal of this search is to see if the chunk contains any free and unscavenged memory.
	for i := len(m.scavenged) - 1; i >= 0; i-- {
		// 1s are scavenged OR non-free => 0s are unscavenged AND free
		//
		// TODO(mknyszek): Consider splitting up fillAligned into two
		// functions, since here we technically could get by with just
		// the first half of its computation. It'll save a few instructions
		// but adds some additional code complexity.
		x := fillAligned(m.scavenged[i]|m.pallocBits[i], uint(min))

		// Quickly skip over chunks of non-free or scavenged pages.
		if x != ^uint64(0) {
			return true
		}
	}
	return false
}

func (p *pageAlloc) scavengeOne(work addrRange, max uintptr, mayUnlock bool) (uintptr, addrRange) {

	// Defensively check if we've received an empty address range.
	// If so, just return.
	if work.size() == 0 {
		// Nothing to do.
		return 0, work
	}
	// Check the prerequisites of work.
	if work.base.addr()%pallocChunkBytes != 0 {
		throw("scavengeOne called with unaligned work region")
	}
	// Calculate the maximum number of pages to scavenge.
	//
	// This should be alignUp(max, pageSize) / pageSize but max can and will
	// be ^uintptr(0), so we need to be very careful not to overflow here.
	// Rather than use alignUp, calculate the number of pages rounded down
	// first, then add back one if necessary.
	maxPages := max / pageSize
	if max%pageSize != 0 {
		maxPages++
	}

	// Calculate the minimum number of pages we can scavenge.
	//
	// Because we can only scavenge whole physical pages, we must
	// ensure that we scavenge at least minPages each time, aligned
	// to minPages*pageSize.
	minPages := physPageSize / pageSize
	if minPages < 1 {
		minPages = 1
	}

	// Helpers for locking and unlocking only if mayUnlock == true.
	lockHeap := func() {
		if mayUnlock {

		}
	}
	unlockHeap := func() {
		if mayUnlock {

		}
	}

	// Fast path: check the chunk containing the top-most address in work,
	// starting at that address's page index in the chunk.
	//
	// Note that work.end() is exclusive, so get the chunk we care about
	// by subtracting 1.
	maxAddr := work.limit.addr() - 1
	maxChunk := chunkIndex(maxAddr)
	if p.summary[len(p.summary)-1][maxChunk].max() >= uint(minPages) {
		// We only bother looking for a candidate if there at least
		// minPages free pages at all.
		base, npages := p.chunkOf(maxChunk).findScavengeCandidate(chunkPageIndex(maxAddr), minPages, maxPages)

		// If we found something, scavenge it and return!
		if npages != 0 {
			work.limit = offAddr{p.scavengeRangeLocked(maxChunk, base, npages)}

			return uintptr(npages) * pageSize, work
		}
	}
	// Update the limit to reflect the fact that we checked maxChunk already.
	work.limit = offAddr{chunkBase(maxChunk)}

	// findCandidate finds the next scavenge candidate in work optimistically.
	//
	// Returns the candidate chunk index and true on success, and false on failure.
	//
	// The heap need not be locked.
	findCandidate := func(work addrRange) (chunkIdx, bool) {
		// Iterate over this work's chunks.
		for i := chunkIndex(work.limit.addr() - 1); i >= chunkIndex(work.base.addr()); i-- {
			// If this chunk is totally in-use or has no unscavenged pages, don't bother
			// doing a more sophisticated check.
			//
			// Note we're accessing the summary and the chunks without a lock, but
			// that's fine. We're being optimistic anyway.

			// Check quickly if there are enough free pages at all.
			if p.summary[len(p.summary)-1][i].max() < uint(minPages) {
				continue
			}

			// Run over the chunk looking harder for a candidate. Again, we could
			// race with a lot of different pieces of code, but we're just being
			// optimistic. Make sure we load the l2 pointer atomically though, to
			// avoid races with heap growth. It may or may not be possible to also
			// see a nil pointer in this case if we do race with heap growth, but
			// just defensively ignore the nils. This operation is optimistic anyway.
			l2 := (*[1 << pallocChunksL2Bits]pallocData)(Loadp(unsafe.Pointer(&p.chunks[i.l1()])))
			if l2 != nil && l2[i.l2()].hasScavengeCandidate(minPages) {
				return i, true
			}
		}
		return 0, false
	}

	// Slow path: iterate optimistically over the in-use address space
	// looking for any free and unscavenged page. If we think we see something,
	// lock and verify it!
	for work.size() != 0 {
		unlockHeap()

		// Search for the candidate.
		candidateChunkIdx, ok := findCandidate(work)

		// Lock the heap. We need to do this now if we found a candidate or not.
		// If we did, we'll verify it. If not, we need to lock before returning
		// anyway.
		lockHeap()

		if !ok {
			// We didn't find a candidate, so we're done.
			work.limit = work.base
			break
		}

		// Find, verify, and scavenge if we can.
		chunk := p.chunkOf(candidateChunkIdx)
		base, npages := chunk.findScavengeCandidate(pallocChunkPages-1, minPages, maxPages)
		if npages > 0 {
			work.limit = offAddr{p.scavengeRangeLocked(candidateChunkIdx, base, npages)}

			return uintptr(npages) * pageSize, work
		}

		// We were fooled, so let's continue from where we left off.
		work.limit = offAddr{chunkBase(candidateChunkIdx)}
	}

	return 0, work
}

func (p *pageAlloc) scavengeRangeLocked(ci chunkIdx, base, npages uint) uintptr {

	p.chunkOf(ci).scavenged.setRange(base, npages)

	// Compute the full address for the start of the range.
	addr := chunkBase(ci) + uintptr(base)*pageSize

	// Update the scavenge low watermark.
	if oAddr := (offAddr{addr}); oAddr.lessThan(p.scav.scavLWM) {
		p.scav.scavLWM = oAddr
	}

	// Only perform the actual scavenging if we're not in a test.
	// It's dangerous to do so otherwise.
	// if p.test {
	// 	return addr
	// }
	// sysUnused(unsafe.Pointer(addr), uintptr(npages)*pageSize)

	// // Update global accounting only when not in test, otherwise
	// // the runtime's accounting will be wrong.
	// nbytes := int64(npages) * pageSize
	// atomic.Xadd64(&memstats.heap_released, nbytes)

	// // Update consistent accounting too.
	// stats := memstats.heapStats.acquire()
	// atomic.Xaddint64(&stats.committed, -nbytes)
	// atomic.Xaddint64(&stats.released, nbytes)
	// memstats.heapStats.release()

	return addr
}

const (
	scavengePercent = 1 // 1%

	// retainExtraPercent represents the amount of memory over the heap goal
	// that the scavenger should keep as a buffer space for the allocator.
	//
	// The purpose of maintaining this overhead is to have a greater pool of
	// unscavenged memory available for allocation (since using scavenged memory
	// incurs an additional cost), to account for heap fragmentation and
	// the ever-changing layout of the heap.
	retainExtraPercent = 10

	maxPhysPageSize = 512 << 10

	// maxPagesPerPhysPage is the maximum number of supported runtime pages per
	// physical page, based on maxPhysPageSize.
	maxPagesPerPhysPage = maxPhysPageSize / pageSize

	// scavengeCostRatio is the approximate ratio between the costs of using previously
	// scavenged memory and scavenging memory.
	//
	// For most systems the cost of scavenging greatly outweighs the costs
	// associated with using scavenged memory, making this constant 0. On other systems
	// (especially ones where "sysUsed" is not just a no-op) this cost is non-trivial.
	//
	// This ratio is used as part of multiplicative factor to help the scavenger account
	// for the additional costs of using scavenged memory in its pacing.
	scavengeCostRatio = 0.7 * (GoosDarwin + GoosIos)

	// scavengeReservationShards determines the amount of memory the scavenger
	// should reserve for scavenging at a time. Specifically, the amount of
	// memory reserved is (heap size in bytes) / scavengeReservationShards.
	scavengeReservationShards = 64
)

const (
	physHugePageSize = 2097152
)

func (m *pallocData) findScavengeCandidate(searchIdx uint, min, max uintptr) (uint, uint) {
	if min&(min-1) != 0 || min == 0 {
		print("runtime: min = ", min, "\n")
		throw("min must be a non-zero power of 2")
	} else if min > maxPagesPerPhysPage {
		print("runtime: min = ", min, "\n")
		throw("min too large")
	}
	// max may not be min-aligned, so we might accidentally truncate to
	// a max value which causes us to return a non-min-aligned value.
	// To prevent this, align max up to a multiple of min (which is always
	// a power of 2). This also prevents max from ever being less than
	// min, unless it's zero, so handle that explicitly.
	if max == 0 {
		max = min
	} else {
		max = alignUp(max, min)
	}

	i := int(searchIdx / 64)
	// Start by quickly skipping over blocks of non-free or scavenged pages.
	for ; i >= 0; i-- {
		// 1s are scavenged OR non-free => 0s are unscavenged AND free
		x := fillAligned(m.scavenged[i]|m.pallocBits[i], uint(min))
		if x != ^uint64(0) {
			break
		}
	}
	if i < 0 {
		// Failed to find any free/unscavenged pages.
		return 0, 0
	}
	// We have something in the 64-bit chunk at i, but it could
	// extend further. Loop until we find the extent of it.

	// 1s are scavenged OR non-free => 0s are unscavenged AND free
	x := fillAligned(m.scavenged[i]|m.pallocBits[i], uint(min))
	z1 := uint(LeadingZeros64(^x))
	run, end := uint(0), uint(i)*64+(64-z1)
	if x<<z1 != 0 {
		// After shifting out z1 bits, we still have 1s,
		// so the run ends inside this word.
		run = uint(LeadingZeros64(x << z1))
	} else {
		// After shifting out z1 bits, we have no more 1s.
		// This means the run extends to the bottom of the
		// word so it may extend into further words.
		run = 64 - z1
		for j := i - 1; j >= 0; j-- {
			x := fillAligned(m.scavenged[j]|m.pallocBits[j], uint(min))
			run += uint(LeadingZeros64(x))
			if x != 0 {
				// The run stopped in this word.
				break
			}
		}
	}

	// Split the run we found if it's larger than max but hold on to
	// our original length, since we may need it later.
	size := run
	if size > uint(max) {
		size = uint(max)
	}
	start := end - size

	// Each huge page is guaranteed to fit in a single palloc chunk.
	//
	// TODO(mknyszek): Support larger huge page sizes.
	// TODO(mknyszek): Consider taking pages-per-huge-page as a parameter
	// so we can write tests for this.
	if physHugePageSize > pageSize && physHugePageSize > physPageSize {
		// We have huge pages, so let's ensure we don't break one by scavenging
		// over a huge page boundary. If the range [start, start+size) overlaps with
		// a free-and-unscavenged huge page, we want to grow the region we scavenge
		// to include that huge page.

		// Compute the huge page boundary above our candidate.
		pagesPerHugePage := uintptr(physHugePageSize / pageSize)
		hugePageAbove := uint(alignUp(uintptr(start), pagesPerHugePage))

		// If that boundary is within our current candidate, then we may be breaking
		// a huge page.
		if hugePageAbove <= end {
			// Compute the huge page boundary below our candidate.
			hugePageBelow := uint(alignDown(uintptr(start), pagesPerHugePage))

			if hugePageBelow >= end-run {
				// We're in danger of breaking apart a huge page since start+size crosses
				// a huge page boundary and rounding down start to the nearest huge
				// page boundary is included in the full run we found. Include the entire
				// huge page in the bound by rounding down to the huge page size.
				size = size + (start - hugePageBelow)
				start = hugePageBelow
			}
		}
	}
	return start, size
}

func fillAligned(x uint64, m uint) uint64 {
	apply := func(x uint64, c uint64) uint64 {
		// The technique used it here is derived from
		// https://graphics.stanford.edu/~seander/bithacks.html#ZeroInWord
		// and extended for more than just bytes (like nibbles
		// and uint16s) by using an appropriate constant.
		//
		// To summarize the technique, quoting from that page:
		// "[It] works by first zeroing the high bits of the [8]
		// bytes in the word. Subsequently, it adds a number that
		// will result in an overflow to the high bit of a byte if
		// any of the low bits were initially set. Next the high
		// bits of the original word are ORed with these values;
		// thus, the high bit of a byte is set iff any bit in the
		// byte was set. Finally, we determine if any of these high
		// bits are zero by ORing with ones everywhere except the
		// high bits and inverting the result."
		return ^((((x & c) + c) | x) | c)
	}
	// Transform x to contain a 1 bit at the top of each m-aligned
	// group of m zero bits.
	switch m {
	case 1:
		return x
	case 2:
		x = apply(x, 0x5555555555555555)
	case 4:
		x = apply(x, 0x7777777777777777)
	case 8:
		x = apply(x, 0x7f7f7f7f7f7f7f7f)
	case 16:
		x = apply(x, 0x7fff7fff7fff7fff)
	case 32:
		x = apply(x, 0x7fffffff7fffffff)
	case 64: // == maxPagesPerPhysPage
		x = apply(x, 0x7fffffffffffffff)
	default:
		throw("bad m value")
	}
	// Now, the top bit of each m-aligned group in x is set
	// that group was all zero in the original x.

	// From each group of m bits subtract 1.
	// Because we know only the top bits of each
	// m-aligned group are set, we know this will
	// set each group to have all the bits set except
	// the top bit, so just OR with the original
	// result to set all the bits.
	return ^((x - (x >> (m - 1))) | x)
}
