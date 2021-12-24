package mem

import (
	"testing"
	"unsafe"
)

const (
	pageBytes  = 1024 * 8        // 8K
	chunkBytes = pageBytes * 512 // 4M
)

func TestPageAlloc(t *testing.T) {
	p := new(pageAlloc)
	p.init()

	// ck := unsafe.Sizeof(*p.chunks[0])
	// t.Log(ck, ck/8192, unsafe.Sizeof(pallocData{}))

	// 33751040 33751048 824633720832
	var base uintptr = 0
	var size uintptr = chunkBytes * 4
	p.grow(base, size) // 往树里装n个trunk的内存

	p.alloc(1) // 分配x页

	// p.update()
}

func TestGetPageSize(t *testing.T) {
	// 		   824637915136
	// base := 824633720832
	// 140299046318080
	// 140299045269504
	// 140649912590336
	// 140514732510208

	// 139677629743104 139678243291136
	// 139677629739008 139678242769920
	// 1649267441664
	var base uintptr
	ask, nbase := GetPageSize(1, base)
	t.Log(ask, nbase, toM(nbase-base))
	dt := 140299046318080 - 140299045269504
	dt1 := 140299045269504 - 140299036880896
	dt3 := 140298969772032 - 140298432901120
	t.Log(toM(uintptr(dt)), toM(uintptr(dt1)), toM(uintptr(dt3)))

	base = 140021188333568
	ask, nbase = GetPageSize(512, base)
	t.Log(ask, nbase, toM(nbase-base), unsafe.Sizeof(int(1)))
}

func TestSum(t *testing.T) {
	t.Logf("%v", freeChunkSum)
	l := 2
	children := newChildren(8, l)
	logMaxPages := levelLogPages[l+1]
	sum := mergeSummaries(children, logMaxPages)

	t.Logf("%v %d", sum, sum)

}

const (
	maxSum = 1152922054362923008
)

var maxSums = [...]uint64{
	9223372036854775808,  // lv0
	1152922054362923008,  // lv1
	144115256795365376,   // lv2
	18014407099420672,    // lv3
	uint64(freeChunkSum), // lv4
}

func TestTrunk(t *testing.T) {
	var base uintptr = 824633720832
	saddr := offAddr{base}
	cpi := chunkPageIndex(saddr.addr())
	ci := chunkIndex(saddr.addr())

	t.Log(cpi, ci)

}

// func getChunk(i chunkIdx) *pallocData {
// 	p := new(pageAlloc)
// 	r := sysAlloc(unsafe.Sizeof(*p.chunks[0]))
// 	l1 := (*[8192]pallocData)(r)
// 	return &*l1[i.l2()]
// }

func newChildren(n, lv int) []pallocSum {
	s := make([]pallocSum, n)
	for i := range s {
		s[i] = pallocSum(maxSums[lv+1])
	}
	// s[1] = 0
	return s
}

func toK(v uintptr) uintptr {
	return v / 1024
}

func toM(v uintptr) uintptr {
	return v / 1024 / 1024
}
