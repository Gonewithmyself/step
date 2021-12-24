package mem

import "fmt"

const Goarch386 = 0
const GoarchAmd64 = 1
const GoarchAmd64p32 = 0
const GoarchArm = 0
const GoarchArmbe = 0
const GoarchArm64 = 0
const GoarchArm64be = 0
const GoarchPpc64 = 0
const GoarchPpc64le = 0
const GoarchMips = 0
const GoarchMipsle = 0
const GoarchMips64 = 0
const GoarchMips64le = 0
const GoarchMips64p32 = 0
const GoarchMips64p32le = 0
const GoarchPpc = 0
const GoarchRiscv = 0
const GoarchRiscv64 = 0
const GoarchS390 = 0
const GoarchS390x = 0
const GoarchSparc = 0
const GoarchSparc64 = 0
const GoarchWasm = 0

const GoosAix = 0
const GoosAndroid = 0
const GoosDarwin = 0
const GoosDragonfly = 0
const GoosFreebsd = 0
const GoosHurd = 0
const GoosIllumos = 0
const GoosIos = 0
const GoosJs = 0
const GoosLinux = 0
const GoosNacl = 0
const GoosNetbsd = 0
const GoosOpenbsd = 0
const GoosPlan9 = 0
const GoosSolaris = 0
const GoosWindows = 1
const GoosZos = 0

const (
	_MaxSmallSize     = 32768
	smallSizeDiv      = 8
	smallSizeMax      = 1024
	largeSizeDiv      = 128
	_NumSizeClasses   = 68
	_PageShift        = 13
	_64bit            = 1 << (^uintptr(0) >> 63) / 2
	arenaL1Bits       = 6 * (_64bit * GoosWindows)
	heapAddrBits      = (_64bit*(1-GoarchWasm)*(1-GoosIos*GoarchArm64))*48 + (1-_64bit+GoarchWasm)*(32-(GoarchMips+GoarchMipsle)) + 33*GoosIos*GoarchArm64
	logHeapArenaBytes = (6+20)*(_64bit*(1-GoosWindows)*(1-GoarchWasm)) + (2+20)*(_64bit*GoosWindows) + (2+20)*(1-_64bit) + (2+20)*GoarchWasm
	arenaL2Bits       = heapAddrBits - logHeapArenaBytes - arenaL1Bits

	_PageSize = 1 << _PageShift
	_PageMask = _PageSize - 1
)
const PtrSize = 4 << (^uintptr(0) >> 63)

const (
	pageShift      = _PageShift
	pageSize       = _PageSize
	pageMask       = _PageMask
	uintptrMask    = 1<<(8*PtrSize) - 1
	heapArenaBytes = 67108864

	pallocChunkPages    = 1 << logPallocChunkPages
	pallocChunkBytes    = pallocChunkPages * pageSize
	logPallocChunkPages = 9
	logPallocChunkBytes = logPallocChunkPages + pageShift
	arenaBaseOffset     = 0xffff800000000000*GoarchAmd64 + 0x0a00000000000000*GoosAix
)
const GOARCH = `amd64`
const GOOS = `linux`
const raceenabled = false

var class_to_size = [_NumSizeClasses]uint16{0, 8, 16, 24, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 208, 224, 240, 256, 288, 320, 352, 384, 416, 448, 480, 512, 576, 640, 704, 768, 896, 1024, 1152, 1280, 1408, 1536, 1792, 2048, 2304, 2688, 3072, 3200, 3456, 4096, 4864, 5376, 6144, 6528, 6784, 6912, 8192, 9472, 9728, 10240, 10880, 12288, 13568, 14336, 16384, 18432, 19072, 20480, 21760, 24576, 27264, 28672, 32768}
var class_to_allocnpages = [_NumSizeClasses]uint8{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 2, 3, 1, 3, 2, 3, 4, 5, 6, 1, 7, 6, 5, 4, 3, 5, 7, 2, 9, 7, 5, 8, 3, 10, 7, 4}

type divMagic struct {
	shift    uint8
	shift2   uint8
	mul      uint16
	baseMask uint16
}

var class_to_divmagic = [_NumSizeClasses]divMagic{{0, 0, 0, 0}, {3, 0, 1, 65528}, {4, 0, 1, 65520}, {3, 11, 683, 0}, {5, 0, 1, 65504}, {4, 11, 683, 0}, {6, 0, 1, 65472}, {4, 10, 205, 0}, {5, 9, 171, 0}, {4, 11, 293, 0}, {7, 0, 1, 65408}, {4, 13, 911, 0}, {5, 10, 205, 0}, {4, 12, 373, 0}, {6, 9, 171, 0}, {4, 13, 631, 0}, {5, 11, 293, 0}, {4, 13, 547, 0}, {8, 0, 1, 65280}, {5, 9, 57, 0}, {6, 9, 103, 0}, {5, 12, 373, 0}, {7, 7, 43, 0}, {5, 10, 79, 0}, {6, 10, 147, 0}, {5, 11, 137, 0}, {9, 0, 1, 65024}, {6, 9, 57, 0}, {7, 9, 103, 0}, {6, 11, 187, 0}, {8, 7, 43, 0}, {7, 8, 37, 0}, {10, 0, 1, 64512}, {7, 9, 57, 0}, {8, 6, 13, 0}, {7, 11, 187, 0}, {9, 5, 11, 0}, {8, 8, 37, 0}, {11, 0, 1, 63488}, {8, 9, 57, 0}, {7, 10, 49, 0}, {10, 5, 11, 0}, {7, 10, 41, 0}, {7, 9, 19, 0}, {12, 0, 1, 61440}, {8, 9, 27, 0}, {8, 10, 49, 0}, {11, 5, 11, 0}, {7, 13, 161, 0}, {7, 13, 155, 0}, {8, 9, 19, 0}, {13, 0, 1, 57344}, {8, 12, 111, 0}, {9, 9, 27, 0}, {11, 6, 13, 0}, {7, 14, 193, 0}, {12, 3, 3, 0}, {8, 13, 155, 0}, {11, 8, 37, 0}, {14, 0, 1, 49152}, {11, 8, 29, 0}, {7, 13, 55, 0}, {12, 5, 7, 0}, {8, 14, 193, 0}, {13, 3, 3, 0}, {7, 14, 77, 0}, {12, 7, 19, 0}, {15, 0, 1, 32768}}
var size_to_class8 = [smallSizeMax/smallSizeDiv + 1]uint8{0, 1, 2, 3, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10, 11, 11, 12, 12, 13, 13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 19, 19, 20, 20, 20, 20, 21, 21, 21, 21, 22, 22, 22, 22, 23, 23, 23, 23, 24, 24, 24, 24, 25, 25, 25, 25, 26, 26, 26, 26, 27, 27, 27, 27, 27, 27, 27, 27, 28, 28, 28, 28, 28, 28, 28, 28, 29, 29, 29, 29, 29, 29, 29, 29, 30, 30, 30, 30, 30, 30, 30, 30, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32}
var size_to_class128 = [(_MaxSmallSize-smallSizeMax)/largeSizeDiv + 1]uint8{32, 33, 34, 35, 36, 37, 37, 38, 38, 39, 39, 40, 40, 40, 41, 41, 41, 42, 43, 43, 44, 44, 44, 44, 44, 45, 45, 45, 45, 45, 45, 46, 46, 46, 46, 47, 47, 47, 47, 47, 47, 48, 48, 48, 49, 49, 50, 51, 51, 51, 51, 51, 51, 51, 51, 51, 51, 52, 52, 52, 52, 52, 52, 52, 52, 52, 52, 53, 53, 54, 54, 54, 54, 55, 55, 55, 55, 55, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 56, 57, 57, 57, 57, 57, 57, 57, 57, 57, 57, 58, 58, 58, 58, 58, 58, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 59, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 60, 61, 61, 61, 61, 61, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 62, 63, 63, 63, 63, 63, 63, 63, 63, 63, 63, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67, 67}

func roundupsize(size uintptr) uintptr {
	if size < _MaxSmallSize {
		if size <= smallSizeMax-8 {
			return uintptr(class_to_size[size_to_class8[divRoundUp(size, smallSizeDiv)]])
		} else {
			return uintptr(class_to_size[size_to_class128[divRoundUp(size-smallSizeMax, largeSizeDiv)]])
		}
	}
	if size+_PageSize < size {
		return size
	}
	return alignUp(size, _PageSize)
}

// alignUp rounds n up to a multiple of a. a must be a power of 2.
func alignUp(n, a uintptr) uintptr {
	return (n + a - 1) &^ (a - 1)
}

// alignDown rounds n down to a multiple of a. a must be a power of 2.
func alignDown(n, a uintptr) uintptr {
	return n &^ (a - 1)
}

func divRoundUp(n, a uintptr) uintptr {
	// a is generally a power of two. This will get inlined and
	// the compiler will optimize the division.
	return (n + a - 1) / a
}

func getarenasz() (int, int) {
	return 1 << arenaL1Bits, 1 << arenaL2Bits
}

type arenaIdx uint

func arenaIndex(p uintptr) arenaIdx {
	return arenaIdx((p - arenaBaseOffset) / heapArenaBytes)
}

type chunkIdx uint

// chunkIndex returns the global index of the palloc chunk containing the
// pointer p.
func chunkIndex(p uintptr) chunkIdx {
	return chunkIdx((p - arenaBaseOffset) / pallocChunkBytes)
}

// chunkIndex returns the base address of the palloc chunk at index ci.
func chunkBase(ci chunkIdx) uintptr {
	return uintptr(ci)*pallocChunkBytes + arenaBaseOffset
}

// chunkPageIndex computes the index of the page that contains p,
// relative to the chunk which contains p.
func chunkPageIndex(p uintptr) uint {
	return uint(p % pallocChunkBytes / pageSize)
}

const (
	pallocChunksL1Bits  = 13
	pallocChunksL2Bits  = heapAddrBits - logPallocChunkBytes - pallocChunksL1Bits
	pallocChunksL1Shift = pallocChunksL2Bits
)

// l1 returns the index into the first level of (*pageAlloc).chunks.
func (i chunkIdx) l1() uint {
	if pallocChunksL1Bits == 0 {
		// Let the compiler optimize this away if there's no
		// L1 map.
		return 0
	} else {
		return uint(i) >> pallocChunksL1Shift
	}
}

// l2 returns the index into the second level of (*pageAlloc).chunks.
func (i chunkIdx) l2() uint {
	if pallocChunksL1Bits == 0 {
		return uint(i)
	} else {
		return uint(i) & (1<<pallocChunksL2Bits - 1)
	}
}
func OnesCount64(x uint64) int {
	// Implementation: Parallel summing of adjacent bits.
	// See "Hacker's Delight", Chap. 5: Counting Bits.
	// The following pattern shows the general approach:
	//
	//   x = x>>1&(m0&m) + x&(m0&m)
	//   x = x>>2&(m1&m) + x&(m1&m)
	//   x = x>>4&(m2&m) + x&(m2&m)
	//   x = x>>8&(m3&m) + x&(m3&m)
	//   x = x>>16&(m4&m) + x&(m4&m)
	//   x = x>>32&(m5&m) + x&(m5&m)
	//   return int(x)
	//
	// Masking (& operations) can be left away when there's no
	// danger that a field's sum will carry over into the next
	// field: Since the result cannot be > 64, 8 bits is enough
	// and we can ignore the masks for the shifts by 8 and up.
	// Per "Hacker's Delight", the first line can be simplified
	// more, but it saves at best one instruction, so we leave
	// it alone for clarity.
	const m0 = 0x5555555555555555 // 01010101 ...
	const m1 = 0x3333333333333333 // 00110011 ...
	const m2 = 0x0f0f0f0f0f0f0f0f // 00001111 ...
	const m = 1<<64 - 1
	x = x>>1&(m0&m) + x&(m0&m)
	x = x>>2&(m1&m) + x&(m1&m)
	x = (x>>4 + x) & (m2 & m)
	x += x >> 8
	x += x >> 16
	x += x >> 32
	return int(x) & (1<<7 - 1)
}

func (b *pageBits) popcntRange(i, n uint) (s uint) {
	if n == 1 {
		return uint((b[i/64] >> (i % 64)) & 1)
	}
	_ = b[i/64]
	j := i + n - 1
	if i/64 == j/64 {
		return uint(OnesCount64((b[i/64] >> (i % 64)) & ((1 << n) - 1)))
	}
	_ = b[j/64]
	s += uint(OnesCount64(b[i/64] >> (i % 64)))
	for k := i/64 + 1; k < j/64; k++ {
		s += uint(OnesCount64(b[k]))
	}
	s += uint(OnesCount64(b[j/64] & ((1 << (j%64 + 1)) - 1)))
	return
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
