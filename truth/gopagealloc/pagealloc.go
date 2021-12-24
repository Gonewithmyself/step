package mem

import (
	"fmt"
	"log"
	"unsafe"
)

const (
	summaryLevels    = 5
	summaryLevelBits = 3
	summaryL0Bits    = heapAddrBits - logPallocChunkBytes - (summaryLevels-1)*summaryLevelBits
)

type pageAlloc struct {
	start, end chunkIdx
	summary    [summaryLevels][]pallocSum
	inUse      addrRanges
	chunks     [1 << pallocChunksL1Bits]*[1 << pallocChunksL2Bits]pallocData
	searchAddr offAddr
}

type pallocData struct {
	pallocBits
	scavenged pageBits
}

type pallocBits pageBits

type pageBits [pallocChunkPages / 64]uint64

var deBruijn64tab = [64]byte{
	0, 1, 56, 2, 57, 49, 28, 3, 61, 58, 42, 50, 38, 29, 17, 4,
	62, 47, 59, 36, 45, 43, 51, 22, 53, 39, 33, 30, 24, 18, 12, 5,
	63, 55, 48, 27, 60, 41, 37, 16, 46, 35, 44, 21, 52, 32, 23, 11,
	54, 26, 40, 15, 34, 20, 31, 10, 25, 14, 19, 9, 13, 8, 7, 6,
}

const deBruijn64 = 0x03f79d71b4ca8b09

// TrailingZeros64 returns the number of trailing zero bits in x; the result is 64 for x == 0.
func TrailingZeros64(x uint64) int {
	if x == 0 {
		return 64
	}
	// If popcount is fast, replace code below with return popcount(^x & (x - 1)).
	//
	// x & -x leaves only the right-most bit set in the word. Let k be the
	// index of that bit. Since only a single bit is set, the value is two
	// to the power of k. Multiplying by a power of two is equivalent to
	// left shifting, in this case by k bits. The de Bruijn (64 bit) constant
	// is such that all six bit, consecutive substrings are distinct.
	// Therefore, if we have a left shifted version of this constant we can
	// find by how many bits it was shifted by looking at which six bit
	// substring ended up at the top of the word.
	// (Knuth, volume 4, section 7.3.1)
	return int(deBruijn64tab[(x&-x)*deBruijn64>>(64-6)])
}

func LeadingZeros64(x uint64) int { return 64 - Len64(x) }
func Len64(x uint64) (n int) {
	if x >= 1<<32 {
		x >>= 32
		n = 32
	}
	if x >= 1<<16 {
		x >>= 16
		n += 16
	}
	if x >= 1<<8 {
		x >>= 8
		n += 8
	}
	return n + int(len8tab[x])
}

var len8tab = [256]uint8{
	0x00, 0x01, 0x02, 0x02, 0x03, 0x03, 0x03, 0x03, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
	0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
	0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06,
	0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06,
	0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
	0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
	0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
	0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
	0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
}

// summarize returns a packed summary of the bitmap in pallocBits.
func (b *pallocBits) summarize() pallocSum {
	var start, max, cur uint
	const notSetYet = ^uint(0) // sentinel for start value
	start = notSetYet
	for i := 0; i < len(b); i++ {
		x := b[i]
		if x == 0 {
			cur += 64
			continue
		}
		t := uint(TrailingZeros64(x))
		l := uint(LeadingZeros64(x))

		// Finish any region spanning the uint64s
		cur += t
		if start == notSetYet {
			start = cur
		}
		if cur > max {
			max = cur
		}
		// Final region that might span to next uint64
		cur = l
	}
	if start == notSetYet {
		// Made it all the way through without finding a single 1 bit.
		const n = uint(64 * len(b))
		return packPallocSum(n, n, n)
	}
	if cur > max {
		max = cur
	}
	if max >= 64-2 {
		// There is no way an internal run of zeros could beat max.
		return packPallocSum(start, max, cur)
	}
	// Now look inside each uint64 for runs of zeros.
	// All uint64s must be nonzero, or we would have aborted above.
outer:
	for i := 0; i < len(b); i++ {
		x := b[i]

		// Look inside this uint64. We have a pattern like
		// 000000 1xxxxx1 000000
		// We need to look inside the 1xxxxx1 for any contiguous
		// region of zeros.

		// We already know the trailing zeros are no larger than max. Remove them.
		x >>= TrailingZeros64(x) & 63
		if x&(x+1) == 0 { // no more zeros (except at the top).
			continue
		}

		// Strategy: shrink all runs of zeros by max. If any runs of zero
		// remain, then we've identified a larger maxiumum zero run.
		p := max     // number of zeros we still need to shrink by.
		k := uint(1) // current minimum length of runs of ones in x.
		for {
			// Shrink all runs of zeros by p places (except the top zeros).
			for p > 0 {
				if p <= k {
					// Shift p ones down into the top of each run of zeros.
					x |= x >> (p & 63)
					if x&(x+1) == 0 { // no more zeros (except at the top).
						continue outer
					}
					break
				}
				// Shift k ones down into the top of each run of zeros.
				x |= x >> (k & 63)
				if x&(x+1) == 0 { // no more zeros (except at the top).
					continue outer
				}
				p -= k
				// We've just doubled the minimum length of 1-runs.
				// This allows us to shift farther in the next iteration.
				k *= 2
			}

			// The length of the lowest-order zero run is an increment to our maximum.
			j := uint(TrailingZeros64(^x)) // count contiguous trailing ones
			x >>= j & 63                   // remove trailing ones
			j = uint(TrailingZeros64(x))   // count contiguous trailing zeros
			x >>= j & 63                   // remove zeros
			max += j                       // we have a new maximum!
			if x&(x+1) == 0 {              // no more zeros (except at the top).
				continue outer
			}
			p = j // remove j more zeros from each zero run.
		}
	}
	return packPallocSum(start, max, cur)
}

var levelShift = [summaryLevels]uint{
	heapAddrBits - summaryL0Bits,
	heapAddrBits - summaryL0Bits - 1*summaryLevelBits,
	heapAddrBits - summaryL0Bits - 2*summaryLevelBits,
	heapAddrBits - summaryL0Bits - 3*summaryLevelBits,
	heapAddrBits - summaryL0Bits - 4*summaryLevelBits,
}

var levelBits = [summaryLevels]uint{
	summaryL0Bits,
	summaryLevelBits,
	summaryLevelBits,
	summaryLevelBits,
	summaryLevelBits,
}

func addrsToSummaryRange(level int, base, limit uintptr) (lo int, hi int) {
	// This is slightly more nuanced than just a shift for the exclusive
	// upper-bound. Note that the exclusive upper bound may be within a
	// summary at this level, meaning if we just do the obvious computation
	// hi will end up being an inclusive upper bound. Unfortunately, just
	// adding 1 to that is too broad since we might be on the very edge of
	// of a summary's max page count boundary for this level
	// (1 << levelLogPages[level]). So, make limit an inclusive upper bound
	// then shift, then add 1, so we get an exclusive upper bound at the end.
	lo = int((base - arenaBaseOffset) >> levelShift[level])
	hi = int(((limit-1)-arenaBaseOffset)>>levelShift[level]) + 1
	return
}

func blockAlignSummaryRange(level int, lo, hi int) (int, int) {
	e := uintptr(1) << levelBits[level]
	return int(alignDown(uintptr(lo), e)), int(alignUp(uintptr(hi), e))
}

var physPageSize uintptr = 4096

func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

func (p *pageAlloc) sysInit() {
	// Reserve memory for each level. This will get mapped in
	// as R/W by setArenas.
	for l, shift := range levelShift {
		entries := 1 << (heapAddrBits - shift)

		// Reserve b bytes of memory anywhere in the address space.
		b := alignUp(uintptr(entries)*pallocSumBytes, physPageSize)
		r := sysReserve(nil, b)
		if r == nil {
			throw("failed to reserve page summary memory")
		}

		// Put this reservation into a slice.
		sl := notInHeapSlice{(*notInHeap)(r), 0, entries}
		p.summary[l] = *(*[]pallocSum)(unsafe.Pointer(&sl))
	}
}

func (p *pageAlloc) init() {
	p.inUse.init(nil)

	p.sysInit()

	p.searchAddr = maxSearchAddr
}

func (p *pageAlloc) sysGrow(base, limit uintptr) {
	addrRangeToSummaryRange := func(level int, r addrRange) (int, int) {
		sumIdxBase, sumIdxLimit := addrsToSummaryRange(level, r.base.addr(), r.limit.addr())
		return blockAlignSummaryRange(level, sumIdxBase, sumIdxLimit)
	}

	// summaryRangeToSumAddrRange converts a range of indices in any
	// level of p.summary into page-aligned addresses which cover that
	// range of indices.
	summaryRangeToSumAddrRange := func(level, sumIdxBase, sumIdxLimit int) addrRange {
		baseOffset := alignDown(uintptr(sumIdxBase)*pallocSumBytes, physPageSize)
		limitOffset := alignUp(uintptr(sumIdxLimit)*pallocSumBytes, physPageSize)
		base := unsafe.Pointer(&p.summary[level][0])
		return addrRange{
			offAddr{uintptr(add(base, baseOffset))},
			offAddr{uintptr(add(base, limitOffset))},
		}
	}

	// addrRangeToSumAddrRange is a convienience function that converts
	// an address range r to the address range of the given summary level
	// that stores the summaries for r.
	addrRangeToSumAddrRange := func(level int, r addrRange) addrRange {
		sumIdxBase, sumIdxLimit := addrRangeToSummaryRange(level, r)
		return summaryRangeToSumAddrRange(level, sumIdxBase, sumIdxLimit)
	}

	// Find the first inUse index which is strictly greater than base.
	//
	// Because this function will never be asked remap the same memory
	// twice, this index is effectively the index at which we would insert
	// this new growth, and base will never overlap/be contained within
	// any existing range.
	//
	// This will be used to look at what memory in the summary array is already
	// mapped before and after this new range.
	inUseIndex := p.inUse.findSucc(base)

	// Walk up the radix tree and map summaries in as needed.
	for l := range p.summary {
		// Figure out what part of the summary array this new address space needs.
		needIdxBase, needIdxLimit := addrRangeToSummaryRange(l, makeAddrRange(base, limit))

		// Update the summary slices with a new upper-bound. This ensures
		// we get tight bounds checks on at least the top bound.
		//
		// We must do this regardless of whether we map new memory.
		if needIdxLimit > len(p.summary[l]) {
			p.summary[l] = p.summary[l][:needIdxLimit]
		}

		// Compute the needed address range in the summary array for level l.
		need := summaryRangeToSumAddrRange(l, needIdxBase, needIdxLimit)

		// Prune need down to what needs to be newly mapped. Some parts of it may
		// already be mapped by what inUse describes due to page alignment requirements
		// for mapping. prune's invariants are guaranteed by the fact that this
		// function will never be asked to remap the same memory twice.
		if inUseIndex > 0 {
			need = need.subtract(addrRangeToSumAddrRange(l, p.inUse.ranges[inUseIndex-1]))
		}
		if inUseIndex < len(p.inUse.ranges) {
			need = need.subtract(addrRangeToSumAddrRange(l, p.inUse.ranges[inUseIndex]))
		}
		// It's possible that after our pruning above, there's nothing new to map.
		if need.size() == 0 {
			continue
		}
		// fmt.Println(need.size(), need.base.addr(), needIdxBase, needIdxLimit)

		// Map and commit need.
		sysMap(unsafe.Pointer(need.base.addr()), need.size())
		// sysUsed(unsafe.Pointer(need.base.addr()), need.size())
	}
}

func (p *pageAlloc) chunkOf(ci chunkIdx) *pallocData {
	return &p.chunks[ci.l1()][ci.l2()]
}

var levelLogPages = [summaryLevels]uint{
	logPallocChunkPages + 4*summaryLevelBits,
	logPallocChunkPages + 3*summaryLevelBits,
	logPallocChunkPages + 2*summaryLevelBits,
	logPallocChunkPages + 1*summaryLevelBits,
	logPallocChunkPages,
}

func (p *pageAlloc) update(base, npages uintptr, contig, alloc bool) {

	// base, limit, start, and end are inclusive.
	limit := base + npages*pageSize - 1
	sc, ec := chunkIndex(base), chunkIndex(limit)

	// Handle updating the lowest level first.
	if sc == ec {
		// Fast path: the allocation doesn't span more than one chunk,
		// so update this one and if the summary didn't change, return.
		x := p.summary[len(p.summary)-1][sc]
		y := p.chunkOf(sc).summarize()
		if x == y {
			return
		}
		p.summary[len(p.summary)-1][sc] = y
		log.Printf("lv(%v) sidx(%v), changed(%v)", len(p.summary)-1, ec, p.summary[len(p.summary)-1][sc])
	} else if contig {
		// Slow contiguous path: the allocation spans more than one chunk
		// and at least one summary is guaranteed to change.
		summary := p.summary[len(p.summary)-1]

		// Update the summary for chunk sc.
		summary[sc] = p.chunkOf(sc).summarize()
		sy := summary[sc].String()
		_ = sy
		// Update the summaries for chunks in between, which are
		// either totally allocated or freed.
		whole := p.summary[len(p.summary)-1][sc+1 : ec]
		if alloc {
			// Should optimize into a memclr.
			for i := range whole {
				whole[i] = 0
			}
		} else {
			for i := range whole {
				whole[i] = freeChunkSum
			}
		}

		// Update the summary for chunk ec.
		summary[ec] = p.chunkOf(ec).summarize()
		log.Printf("lv(%v) sidx(%v), changed(%v)", len(p.summary)-1, ec, summary[ec])
	} else {
		// Slow general path: the allocation spans more than one chunk
		// and at least one summary is guaranteed to change.
		//
		// We can't assume a contiguous allocation happened, so walk over
		// every chunk in the range and manually recompute the summary.
		summary := p.summary[len(p.summary)-1]
		for c := sc; c <= ec; c++ {
			summary[c] = p.chunkOf(c).summarize()
			log.Printf("lv(%v) sidx(%v), changed(%v)", len(p.summary)-1, c, summary[c])
		}
	}

	// Walk up the radix tree and update the summaries appropriately.
	changed := true
	for l := len(p.summary) - 2; l >= 0 && changed; l-- {
		// Update summaries at level l from summaries at level l+1.
		changed = false

		// "Constants" for the previous level which we
		// need to compute the summary from that level.
		logEntriesPerBlock := levelBits[l+1]
		logMaxPages := levelLogPages[l+1]

		// lo and hi describe all the parts of the level we need to look at.
		lo, hi := addrsToSummaryRange(l, base, limit+1)

		// Iterate over each block, updating the corresponding summary in the less-granular level.
		for i := lo; i < hi; i++ {
			children := p.summary[l+1][i<<logEntriesPerBlock : (i+1)<<logEntriesPerBlock]
			sum := mergeSummaries(children, logMaxPages)
			old := p.summary[l][i]
			if old != sum {
				changed = true
				p.summary[l][i] = sum
				log.Printf("lv(%v) sidx(%v), changed(%v)", l, i, p.summary[l][i])
			}
		}
	}
}

type pallocSum uint64

func (x pallocSum) String() string {
	return fmt.Sprintf("s(%v)m(%v)e(%v), %b ",
		x.start(), x.max(), x.end(), x)
}

func (x pallocSum) Show() (uint, uint, uint) {
	return x.start(), x.max(), x.end()
}

const (
	pallocSumBytes = unsafe.Sizeof(pallocSum(0))

	// maxPackedValue is the maximum value that any of the three fields in
	// the pallocSum may take on.
	maxPackedValue    = 1 << logMaxPackedValue
	logMaxPackedValue = logPallocChunkPages + (summaryLevels-1)*summaryLevelBits

	freeChunkSum = pallocSum(uint64(pallocChunkPages) |
		uint64(pallocChunkPages<<logMaxPackedValue) |
		uint64(pallocChunkPages<<(2*logMaxPackedValue)))
)

// packPallocSum takes a start, max, and end value and produces a pallocSum.
func packPallocSum(start, max, end uint) pallocSum {
	if max == maxPackedValue {
		return pallocSum(uint64(1 << 63))
	}
	return pallocSum((uint64(start) & (maxPackedValue - 1)) |
		((uint64(max) & (maxPackedValue - 1)) << logMaxPackedValue) |
		((uint64(end) & (maxPackedValue - 1)) << (2 * logMaxPackedValue)))
}

// start extracts the start value from a packed sum.
func (p pallocSum) start() uint {
	if uint64(p)&uint64(1<<63) != 0 {
		return maxPackedValue
	}
	return uint(uint64(p) & (maxPackedValue - 1))
}

// max extracts the max value from a packed sum.
func (p pallocSum) max() uint {
	if uint64(p)&uint64(1<<63) != 0 {
		return maxPackedValue
	}
	return uint((uint64(p) >> logMaxPackedValue) & (maxPackedValue - 1))
}

// end extracts the end value from a packed sum.
func (p pallocSum) end() uint {
	if uint64(p)&uint64(1<<63) != 0 {
		return maxPackedValue
	}
	return uint((uint64(p) >> (2 * logMaxPackedValue)) & (maxPackedValue - 1))
}

// unpack unpacks all three values from the summary.
func (p pallocSum) unpack() (uint, uint, uint) {
	if uint64(p)&uint64(1<<63) != 0 {
		return maxPackedValue, maxPackedValue, maxPackedValue
	}
	return uint(uint64(p) & (maxPackedValue - 1)),
		uint((uint64(p) >> logMaxPackedValue) & (maxPackedValue - 1)),
		uint((uint64(p) >> (2 * logMaxPackedValue)) & (maxPackedValue - 1))
}

func mergeSummaries(sums []pallocSum, logMaxPagesPerSum uint) pallocSum {
	// Merge the summaries in sums into one.
	//
	// We do this by keeping a running summary representing the merged
	// summaries of sums[:i] in start, max, and end.
	start, max, end := sums[0].unpack()
	//start, max, end := (*(*pallocSum)(getaddr(unsafe.Pointer(&sums[0]), 0))).unpack()
	for i := 1; i < len(sums); i++ {
		// Merge in sums[i].
		si, mi, ei := sums[i].unpack()
		// si, mi, ei := (*(*pallocSum)(getaddr(unsafe.Pointer(&sums[0]), i))).unpack()

		// Merge in sums[i].start only if the running summary is
		// completely free, otherwise this summary's start
		// plays no role in the combined sum.
		if start == uint(i)<<logMaxPagesPerSum {
			start += si
		}

		// Recompute the max value of the running sum by looking
		// across the boundary between the running sum and sums[i]
		// and at the max sums[i], taking the greatest of those two
		// and the max of the running sum.
		if end+si > max {
			max = end + si
		}
		if mi > max {
			max = mi
		}

		// Merge in end by checking if this new summary is totally
		// free. If it is, then we want to extend the running sum's
		// end by the new summary. If not, then we have some alloc'd
		// pages in there and we just want to take the end value in
		// sums[i].
		if ei == 1<<logMaxPagesPerSum {
			end += 1 << logMaxPagesPerSum
		} else {
			end = ei
		}
	}
	return packPallocSum(start, max, end)
}

func offAddrToLevelIndex(level int, addr offAddr) int {
	return int((addr.a - arenaBaseOffset) >> levelShift[level])
}

func levelIndexToOffAddr(level, idx int) offAddr {
	return offAddr{(uintptr(idx) << levelShift[level]) + arenaBaseOffset}
}

var maxSearchAddr = maxOffAddr

func (p *pageAlloc) findMappedAddr(addr offAddr) offAddr {

	// If we're not in a test, validate first by checking mheap_.arenas.
	// This is a fast path which is only safe to use outside of testing.

	return addr
}

func (p *pageAlloc) find(npages uintptr) (uintptr, offAddr) {
	// Search algorithm.
	//
	// This algorithm walks each level l of the radix tree from the root level
	// to the leaf level. It iterates over at most 1 << levelBits[l] of entries
	// in a given level in the radix tree, and uses the summary information to
	// find either:
	//  1) That a given subtree contains a large enough contiguous region, at
	//     which point it continues iterating on the next level, or
	//  2) That there are enough contiguous boundary-crossing bits to satisfy
	//     the allocation, at which point it knows exactly where to start
	//     allocating from.
	//
	// i tracks the index into the current level l's structure for the
	// contiguous 1 << levelBits[l] entries we're actually interested in.
	//
	// NOTE: Technically this search could allocate a region which crosses
	// the arenaBaseOffset boundary, which when arenaBaseOffset != 0, is
	// a discontinuity. However, the only way this could happen is if the
	// page at the zero address is mapped, and this is impossible on
	// every system we support where arenaBaseOffset != 0. So, the
	// discontinuity is already encoded in the fact that the OS will never
	// map the zero page for us, and this function doesn't try to handle
	// this case in any way.

	// i is the beginning of the block of entries we're searching at the
	// current level.
	i := 0

	// firstFree is the region of address space that we are certain to
	// find the first free page in the heap. base and bound are the inclusive
	// bounds of this window, and both are addresses in the linearized, contiguous
	// view of the address space (with arenaBaseOffset pre-added). At each level,
	// this window is narrowed as we find the memory region containing the
	// first free page of memory. To begin with, the range reflects the
	// full process address space.
	//
	// firstFree is updated by calling foundFree each time free space in the
	// heap is discovered.
	//
	// At the end of the search, base.addr() is the best new
	// searchAddr we could deduce in this search.
	firstFree := struct {
		base, bound offAddr
	}{
		base:  minOffAddr,
		bound: maxOffAddr,
	}
	// foundFree takes the given address range [addr, addr+size) and
	// updates firstFree if it is a narrower range. The input range must
	// either be fully contained within firstFree or not overlap with it
	// at all.
	//
	// This way, we'll record the first summary we find with any free
	// pages on the root level and narrow that down if we descend into
	// that summary. But as soon as we need to iterate beyond that summary
	// in a level to find a large enough range, we'll stop narrowing.
	foundFree := func(addr offAddr, size uintptr) {
		if firstFree.base.lessEqual(addr) && addr.add(size-1).lessEqual(firstFree.bound) {
			// This range fits within the current firstFree window, so narrow
			// down the firstFree window to the base and bound of this range.
			firstFree.base = addr
			firstFree.bound = addr.add(size - 1)
		} else if !(addr.add(size-1).lessThan(firstFree.base) || firstFree.bound.lessThan(addr)) {
			// This range only partially overlaps with the firstFree range,
			// so throw.
			// print("runtime: addr = ", hex(addr.addr()), ", size = ", size, "\n")
			// print("runtime: base = ", hex(firstFree.base.addr()), ", bound = ", hex(firstFree.bound.addr()), "\n")
			throw("range partially overlaps")
		}
	}

	// lastSum is the summary which we saw on the previous level that made us
	// move on to the next level. Used to print additional information in the
	// case of a catastrophic failure.
	// lastSumIdx is that summary's index in the previous level.
	lastSum := packPallocSum(0, 0, 0)
	lastSumIdx := -1
	_ = lastSum
	_ = lastSumIdx

nextLevel:
	for l := 0; l < len(p.summary); l++ {
		// For the root level, entriesPerBlock is the whole level.
		entriesPerBlock := 1 << levelBits[l]
		logMaxPages := levelLogPages[l]

		// We've moved into a new level, so let's update i to our new
		// starting index. This is a no-op for level 0.
		i <<= levelBits[l]

		// Slice out the block of entries we care about.
		entries := p.summary[l][i : i+entriesPerBlock]

		// Determine j0, the first index we should start iterating from.
		// The searchAddr may help us eliminate iterations if we followed the
		// searchAddr on the previous level or we're on the root leve, in which
		// case the searchAddr should be the same as i after levelShift.
		j0 := 0
		if searchIdx := offAddrToLevelIndex(l, p.searchAddr); searchIdx&^(entriesPerBlock-1) == i {
			j0 = searchIdx & (entriesPerBlock - 1)
		}

		// Run over the level entries looking for
		// a contiguous run of at least npages either
		// within an entry or across entries.
		//
		// base contains the page index (relative to
		// the first entry's first page) of the currently
		// considered run of consecutive pages.
		//
		// size contains the size of the currently considered
		// run of consecutive pages.
		var base, size uint
		for j := j0; j < len(entries); j++ {
			sum := entries[j]
			if sum == 0 {
				// A full entry means we broke any streak and
				// that we should skip it altogether.
				size = 0
				continue
			}

			// We've encountered a non-zero summary which means
			// free memory, so update firstFree.
			foundFree(levelIndexToOffAddr(l, i+j), (uintptr(1)<<logMaxPages)*pageSize)

			s := sum.start()
			if size+s >= uint(npages) {
				// If size == 0 we don't have a run yet,
				// which means base isn't valid. So, set
				// base to the first page in this block.
				if size == 0 {
					base = uint(j) << logMaxPages
				}
				// We hit npages; we're done!
				size += s
				break
			}
			if sum.max() >= uint(npages) {
				// The entry itself contains npages contiguous
				// free pages, so continue on the next level
				// to find that run.
				i += j
				lastSumIdx = i
				lastSum = sum
				continue nextLevel
			}
			if size == 0 || s < 1<<logMaxPages {
				// We either don't have a current run started, or this entry
				// isn't totally free (meaning we can't continue the current
				// one), so try to begin a new run by setting size and base
				// based on sum.end.
				size = sum.end()
				base = uint(j+1)<<logMaxPages - size
				continue
			}
			// The entry is completely free, so continue the run.
			size += 1 << logMaxPages
		}
		if size >= uint(npages) {
			// We found a sufficiently large run of free pages straddling
			// some boundary, so compute the address and return it.
			addr := levelIndexToOffAddr(l, i).add(uintptr(base) * pageSize).addr()
			return addr, p.findMappedAddr(firstFree.base)
		}
		if l == 0 {
			// We're at level zero, so that means we've exhausted our search.
			return 0, maxSearchAddr
		}

		// We're not at level zero, and we exhausted the level we were looking in.
		// This means that either our calculations were wrong or the level above
		// lied to us. In either case, dump some useful state and throw.
		// print("runtime: summary[", l-1, "][", lastSumIdx, "] = ", lastSum.start(), ", ", lastSum.max(), ", ", lastSum.end(), "\n")
		// print("runtime: level = ", l, ", npages = ", npages, ", j0 = ", j0, "\n")
		// print("runtime: p.searchAddr = ", hex(p.searchAddr.addr()), ", i = ", i, "\n")
		// print("runtime: levelShift[level] = ", levelShift[l], ", levelBits[level] = ", levelBits[l], "\n")
		// for j := 0; j < len(entries); j++ {
		// 	sum := entries[j]
		// 	print("runtime: summary[", l, "][", i+j, "] = (", sum.start(), ", ", sum.max(), ", ", sum.end(), ")\n")
		// }
		// throw("bad summary data")
	}

	// Since we've gotten to this point, that means we haven't found a
	// sufficiently-sized free region straddling some boundary (chunk or larger).
	// This means the last summary we inspected must have had a large enough "max"
	// value, so look inside the chunk to find a suitable run.
	//
	// After iterating over all levels, i must contain a chunk index which
	// is what the final level represents.
	ci := chunkIdx(i)
	j, searchIdx := p.chunkOf(ci).find(npages, 0)
	if j == ^uint(0) {
		// We couldn't find any space in this chunk despite the summaries telling
		// us it should be there. There's likely a bug, so dump some state and throw.
		sum := p.summary[len(p.summary)-1][i]
		print("runtime: summary[", len(p.summary)-1, "][", i, "] = (", sum.start(), ", ", sum.max(), ", ", sum.end(), ")\n")
		print("runtime: npages = ", npages, "\n")
		throw("bad summary data")
	}

	// Compute the address at which the free space starts.
	addr := chunkBase(ci) + uintptr(j)*pageSize

	// Since we actually searched the chunk, we may have
	// found an even narrower free window.
	searchAddr := chunkBase(ci) + uintptr(searchIdx)*pageSize
	foundFree(offAddr{searchAddr}, chunkBase(ci+1)-searchAddr)
	return addr, p.findMappedAddr(firstFree.base)
}

func (m *pallocData) allocRange(i, n uint) {
	// Clear the scavenged bits when we alloc the range.
	m.pallocBits.allocRange(i, n)
	m.scavenged.clearRange(i, n)
}

// allocRange allocates the range [i, i+n).
func (b *pallocBits) allocRange(i, n uint) {
	(*pageBits)(b).setRange(i, n)
}

// allocAll allocates all the bits of b.
func (b *pallocBits) allocAll() {
	(*pageBits)(b).setAll()
}

func (b *pallocBits) find(npages uintptr, searchIdx uint) (uint, uint) {
	if npages == 1 {
		addr := b.find1(searchIdx)
		return addr, addr
	} else if npages <= 64 {
		return b.findSmallN(npages, searchIdx)
	}
	return b.findLargeN(npages, searchIdx)
}

// find1 is a helper for find which searches for a single free page
// in the pallocBits and returns the index.
//
// See find for an explanation of the searchIdx parameter.
func (b *pallocBits) find1(searchIdx uint) uint {
	_ = b[0] // lift nil check out of loop
	for i := searchIdx / 64; i < uint(len(b)); i++ {
		x := b[i]
		if ^x == 0 {
			continue
		}
		return i*64 + uint(TrailingZeros64(^x))
	}
	return ^uint(0)
}

func findBitRange64(c uint64, n uint) uint {
	// This implementation is based on shrinking the length of
	// runs of contiguous 1 bits. We remove the top n-1 1 bits
	// from each run of 1s, then look for the first remaining 1 bit.
	p := n - 1   // number of 1s we want to remove.
	k := uint(1) // current minimum width of runs of 0 in c.
	for p > 0 {
		if p <= k {
			// Shift p 0s down into the top of each run of 1s.
			c &= c >> (p & 63)
			break
		}
		// Shift k 0s down into the top of each run of 1s.
		c &= c >> (k & 63)
		if c == 0 {
			return 64
		}
		p -= k
		// We've just doubled the minimum length of 0-runs.
		// This allows us to shift farther in the next iteration.
		k *= 2
	}
	// Find first remaining 1.
	// Since we shrunk from the top down, the first 1 is in
	// its correct original position.
	return uint(TrailingZeros64(c))
}

func (b *pallocBits) findSmallN(npages uintptr, searchIdx uint) (uint, uint) {
	end, newSearchIdx := uint(0), ^uint(0)
	for i := searchIdx / 64; i < uint(len(b)); i++ {
		bi := b[i]
		if ^bi == 0 {
			end = 0
			continue
		}
		// First see if we can pack our allocation in the trailing
		// zeros plus the end of the last 64 bits.
		if newSearchIdx == ^uint(0) {
			// The new searchIdx is going to be at these 64 bits after any
			// 1s we file, so count trailing 1s.
			newSearchIdx = i*64 + uint(TrailingZeros64(^bi))
		}
		start := uint(TrailingZeros64(bi))
		if end+start >= uint(npages) {
			return i*64 - end, newSearchIdx
		}
		// Next, check the interior of the 64-bit chunk.
		j := findBitRange64(^bi, uint(npages))
		if j < 64 {
			return i*64 + j, newSearchIdx
		}
		end = uint(LeadingZeros64(bi))
	}
	return ^uint(0), newSearchIdx
}

// findLargeN is a helper for find which searches for npages contiguous free pages
// in this pallocBits and returns the index where that run starts, as well as the
// index of the first free page it found it its search.
//
// See alloc for an explanation of the searchIdx parameter.
//
// Returns a ^uint(0) index on failure and the new searchIdx should be ignored.
//
// findLargeN assumes npages > 64, where any such run of free pages
// crosses at least one aligned 64-bit boundary in the bits.
func (b *pallocBits) findLargeN(npages uintptr, searchIdx uint) (uint, uint) {
	start, size, newSearchIdx := ^uint(0), uint(0), ^uint(0)
	for i := searchIdx / 64; i < uint(len(b)); i++ {
		x := b[i]
		if x == ^uint64(0) {
			size = 0
			continue
		}
		if newSearchIdx == ^uint(0) {
			// The new searchIdx is going to be at these 64 bits after any
			// 1s we file, so count trailing 1s.
			newSearchIdx = i*64 + uint(TrailingZeros64(^x))
		}
		if size == 0 {
			size = uint(LeadingZeros64(x))
			start = i*64 + 64 - size
			continue
		}
		s := uint(TrailingZeros64(x))
		if s+size >= uint(npages) {
			size += s
			return start, newSearchIdx
		}
		if s < 64 {
			size = uint(LeadingZeros64(x))
			start = i*64 + 64 - size
			continue
		}
		size += 64
	}
	if size < uint(npages) {
		return ^uint(0), newSearchIdx
	}
	return start, newSearchIdx
}

func (b *pageBits) get(i uint) uint {
	return uint((b[i/64] >> (i % 64)) & 1)
}

// block64 returns the 64-bit aligned block of bits containing the i'th bit.
func (b *pageBits) block64(i uint) uint64 {
	return b[i/64]
}

// set sets bit i of pageBits.
func (b *pageBits) set(i uint) {
	b[i/64] |= 1 << (i % 64)
}

// setRange sets bits in the range [i, i+n).
func (b *pageBits) setRange(i, n uint) {
	_ = b[i/64]
	if n == 1 {
		// Fast path for the n == 1 case.
		b.set(i)
		return
	}
	// Set bits [i, j].
	j := i + n - 1
	if i/64 == j/64 {
		b[i/64] |= ((uint64(1) << n) - 1) << (i % 64)
		return
	}
	_ = b[j/64]
	// Set leading bits.
	b[i/64] |= ^uint64(0) << (i % 64)
	for k := i/64 + 1; k < j/64; k++ {
		b[k] = ^uint64(0)
	}
	// Set trailing bits.
	b[j/64] |= (uint64(1) << (j%64 + 1)) - 1
}

// setAll sets all the bits of b.
func (b *pageBits) setAll() {
	for i := range b {
		b[i] = ^uint64(0)
	}
}

// clear clears bit i of pageBits.
func (b *pageBits) clear(i uint) {
	b[i/64] &^= 1 << (i % 64)
}

// clearRange clears bits in the range [i, i+n).
func (b *pageBits) clearRange(i, n uint) {
	_ = b[i/64]
	if n == 1 {
		// Fast path for the n == 1 case.
		b.clear(i)
		return
	}
	// Clear bits [i, j].
	j := i + n - 1
	if i/64 == j/64 {
		b[i/64] &^= ((uint64(1) << n) - 1) << (i % 64)
		return
	}
	_ = b[j/64]
	// Clear leading bits.
	b[i/64] &^= ^uint64(0) << (i % 64)
	for k := i/64 + 1; k < j/64; k++ {
		b[k] = 0
	}
	// Clear trailing bits.
	b[j/64] &^= (uint64(1) << (j%64 + 1)) - 1
}

// clearAll frees all the bits of b.
func (b *pageBits) clearAll() {
	for i := range b {
		b[i] = 0
	}
}
