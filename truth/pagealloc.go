package truth

import (
	"fmt"
	"unsafe"
)

const (
	summaryLevels    = 5
	summaryLevelBits = 3
	summaryL0Bits    = heapAddrBits - logPallocChunkBytes - (summaryLevels-1)*summaryLevelBits
)

type pageAlloc struct {
	summary [summaryLevels][]pallocSum
	inUse   addrRanges
	chunks  [1 << pallocChunksL1Bits]*[1 << pallocChunksL2Bits]pallocData
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

var physPageSize uintptr = 4

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

func sysReserve(v unsafe.Pointer, n uintptr) unsafe.Pointer {
	buf := make([]byte, n)
	return unsafe.Pointer(&buf[0])
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
		fmt.Println(level, limitOffset, baseOffset)
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
		// var idx = []int{131072, 527424, 4218944, 33751104, 270008384}
		// d := p.summary[l][idx[l]]
		// _ = d
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

		xx := need.size()
		fmt.Println(l, xx, xx/1024/1024, need.base.addr() > base, need.limit.addr() > base)
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
	} else if contig {
		// Slow contiguous path: the allocation spans more than one chunk
		// and at least one summary is guaranteed to change.
		summary := p.summary[len(p.summary)-1]

		// Update the summary for chunk sc.
		summary[sc] = p.chunkOf(sc).summarize()

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
	} else {
		// Slow general path: the allocation spans more than one chunk
		// and at least one summary is guaranteed to change.
		//
		// We can't assume a contiguous allocation happened, so walk over
		// every chunk in the range and manually recompute the summary.
		summary := p.summary[len(p.summary)-1]
		for c := sc; c <= ec; c++ {
			summary[c] = p.chunkOf(c).summarize()
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
			}
		}
	}
}

type pallocSum uint64

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
	for i := 1; i < len(sums); i++ {
		// Merge in sums[i].
		si, mi, ei := sums[i].unpack()

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
