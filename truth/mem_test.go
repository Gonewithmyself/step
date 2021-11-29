package truth

import (
	"fmt"
	"testing"
)

func Test_class_to_size(t *testing.T) {
	var n int
	for i := range class_to_size {
		if class_to_size[i] != 0 {
			n++
		}
	}

	t.Log(len(class_to_size), n, class_to_size[67])

	l1, l2 := getarenasz()
	t.Log(l1, l2/1024)
}

func Test_allocPages(t *testing.T) {
	// 4M
	prev := uintptr(0)
	for i := 0; i < 1000; i++ {
		n := alignUp(uintptr(i), pallocChunkPages)
		if n != prev {
			t.Log(prev, n, i)
			prev = n
		}
	}

	const end = 512 * pageSize
	const physPageSize = 4096
	nBase := alignUp(end, physPageSize)
	t.Log(nBase, nBase/1024/1024, end/1024/1024)

	var base uintptr = 824633720832
	ask := alignUp(1, pallocChunkPages) * pageSize
	nBase = alignUp(base+ask, physPageSize)
	t.Log(nBase - uintptr(base))

	size := nBase - uintptr(base)
	limit := alignUp(base+size, pallocChunkBytes)
	v := base
	base = alignDown(base, pallocChunkBytes)
	t.Log(v == base, limit-base, size)
}

func TestTrunk(t *testing.T) {

	var size uintptr = 512 * pageSize
	var base uintptr = 824633720832 + size*10
	limit := alignUp(base+size, pallocChunkBytes)
	start, end := chunkIndex(base), chunkIndex(limit)

	t.Log(start, end, start.l1(), start.l2(), 824633720832/1024/1024/1024)

	// var stat sysMemStat
	// p := pageAlloc{}
	// p.sysInit()
	// p.inUse.init(&stat)

	// p.sysGrow(base, limit)
	// p.update(base, size/pageSize, true, false)
}

func TestArena(t *testing.T) {

}

func TestSummarize(t *testing.T) {
	var p pallocBits

	// 2251800887427584
	s := p.summarize()
	bs := fmt.Sprintf("%b", s)
	bst := fmt.Sprintf("%b", s.end())
	bsmax := fmt.Sprintf("%b", s.max())
	t.Logf("%b %b %v", s, s.max(), len(bs))
	t.Log(s.max(), s.start(), s.end(), len(bst), len(bsmax), tob(p))

	var npages uintptr = 3
	var base uintptr = 824633720832
	limit := base + npages*pageSize - 1
	// sc, ec := chunkIndex(base), chunkIndex(limit)
	si, ei := chunkPageIndex(base), chunkPageIndex(limit)
	p.allocRange(si, ei+1-si)
	s = p.summarize()
	t.Log(s.start(), s.max(), s.end(), tob(p))

	ac := &bitsAlloc{
		curr: 824633720832,
	}

	t.Log(ac)

	ac.alloc(1)
	t.Log(ac)

	ac.alloc(2)
	t.Log(ac)
}

type bitsAlloc struct {
	pallocBits
	curr uintptr
}

func (p *bitsAlloc) alloc(npages uintptr) {
	base := p.curr
	limit := base + npages*pageSize - 1

	si, ei := chunkPageIndex(base), chunkPageIndex(limit)
	p.allocRange(si, ei+1-si)
	p.curr += npages * pageSize
}

func (p bitsAlloc) String() string {
	s := p.summarize()
	return fmt.Sprintf("%v %v %v %v", s.start(), s.max(), s.end(), tob(p))
}

func tob(s interface{}) string {
	return fmt.Sprintf("%b", s)
}

func TestArenaHint(t *testing.T) {
	narena := 4194304
	arenasz := heapArenaBytes
	total := arenasz * narena
	t.Log(total, total/1024/1024/1024, total/heapArenaBytes)

	var base uintptr = 824633720832
	first := base + heapArenaBytes - 1
	t.Log(arenaIndex(base), arenaIndex(first))

	arenahint()

}

func (p *pageAlloc) allocRange(base, npages uintptr) uintptr {

	limit := base + npages*pageSize - 1
	sc, ec := chunkIndex(base), chunkIndex(limit)
	si, ei := chunkPageIndex(base), chunkPageIndex(limit)

	scav := uint(0)
	if sc == ec {
		// The range doesn't cross any chunk boundaries.
		chunk := p.chunkOf(sc)
		scav += chunk.scavenged.popcntRange(si, ei+1-si)
		chunk.allocRange(si, ei+1-si)
	} else {
		// The range crosses at least one chunk boundary.
		chunk := p.chunkOf(sc)
		scav += chunk.scavenged.popcntRange(si, pallocChunkPages-si)
		chunk.allocRange(si, pallocChunkPages-si)
		for c := sc + 1; c < ec; c++ {
			chunk := p.chunkOf(c)
			scav += chunk.scavenged.popcntRange(0, pallocChunkPages)
			chunk.allocAll()
		}
		chunk = p.chunkOf(ec)
		scav += chunk.scavenged.popcntRange(0, ei+1)
		chunk.allocRange(0, ei+1)
	}
	p.update(base, npages, true, true)
	return uintptr(scav) * pageSize
}

func arenahint() {
	var prev uintptr
	var n = 0
	for i := 0x7f; i >= 0; i-- {
		var p uintptr
		switch {
		case raceenabled:
			// The TSAN runtime requires the heap
			// to be in the range [0x00c000000000,
			// 0x00e000000000).
			p = uintptr(i)<<32 | uintptrMask&(0x00c0<<32)
			if p >= uintptrMask&0x00e000000000 {
				continue
			}
		case GOARCH == "arm64" && GOOS == "ios":
			p = uintptr(i)<<40 | uintptrMask&(0x0013<<28)
		case GOARCH == "arm64":
			p = uintptr(i)<<40 | uintptrMask&(0x0040<<32)
		case GOOS == "aix":
			if i == 0 {
				// We don't use addresses directly after 0x0A00000000000000
				// to avoid collisions with others mmaps done by non-go programs.
				continue
			}
			p = uintptr(i)<<40 | uintptrMask&(0xa0<<52)
		default:
			p = uintptr(i)<<40 | uintptrMask&(0x00c0<<32)
		}
		fmt.Println(prev, p, prev-p, arenaIndex(p))
		prev = p
		n++
	}
	fmt.Println(n)
}
