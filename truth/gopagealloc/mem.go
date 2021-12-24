package mem

import (
	"syscall"
	"unsafe"
)

func GetPageSize(npage, base uintptr) (ask, nBase uintptr) {
	ask = alignUp(npage, pallocChunkPages) * pageSize

	end := base + ask
	nBase = alignUp(end, physPageSize)
	return
}

func (p *pageAlloc) grow(base, size uintptr) {
	limit := alignUp(base+size, pallocChunkBytes)
	base = alignDown(base, pallocChunkBytes)
	p.sysGrow(base, limit)

	firstGrowth := p.start == 0
	start, end := chunkIndex(base), chunkIndex(limit)
	if firstGrowth || start < p.start {
		p.start = start
	}

	if end > p.end {
		p.end = end
	}

	p.inUse.add(makeAddrRange(base, limit))

	if b := (offAddr{base}); b.lessThan(p.searchAddr) {
		p.searchAddr = b
	}

	for c := chunkIndex(base); c < chunkIndex(limit); c++ {
		if p.chunks[c.l1()] == nil {
			// Create the necessary l2 entry.
			//
			// Store it atomically to avoid races with readers which
			// don't acquire the heap lock.
			r := sysAlloc(unsafe.Sizeof(*p.chunks[0]))
			if r == nil {
				throw("pageAlloc: out of memory")
			}
			// atomic.StorepNoWB(unsafe.Pointer(&p.chunks[c.l1()]), r)
			p.chunks[c.l1()] = (*[8192]pallocData)(r)
		}
		p.chunkOf(c).scavenged.setRange(0, pallocChunkPages)
		// fmt.Printf("set chunk %b \n", p.chunkOf(c).scavenged)
	}

	p.update(base, size/pageSize, true, false)
}

func (p *pageAlloc) alloc(npages uintptr) (addr uintptr, scav uintptr) {

	if chunkIndex(p.searchAddr.addr()) >= p.end {
		return 0, 0
	}

	// If npages has a chance of fitting in the chunk where the searchAddr is,
	// search it directly.
	searchAddr := minOffAddr
	if pallocChunkPages-chunkPageIndex(p.searchAddr.addr()) >= uint(npages) {
		// npages is guaranteed to be no greater than pallocChunkPages here.
		i := chunkIndex(p.searchAddr.addr())
		if max := p.summary[len(p.summary)-1][i].max(); max >= uint(npages) {
			j, searchIdx := p.chunkOf(i).find(npages, chunkPageIndex(p.searchAddr.addr()))
			if j == ^uint(0) {
				print("runtime: max = ", max, ", npages = ", npages, "\n")
				// print("runtime: searchIdx = ", chunkPageIndex(p.searchAddr.addr()), ", p.searchAddr = ", hex(p.searchAddr.addr()), "\n")
				throw("bad summary data")
			}
			addr = chunkBase(i) + uintptr(j)*pageSize
			searchAddr = offAddr{chunkBase(i) + uintptr(searchIdx)*pageSize}
			goto Found
		}
	}
	// We failed to use a searchAddr for one reason or another, so try
	// the slow path.
	addr, searchAddr = p.find(npages)
	if addr == 0 {
		if npages == 1 {
			// We failed to find a single free page, the smallest unit
			// of allocation. This means we know the heap is completely
			// exhausted. Otherwise, the heap still might have free
			// space in it, just not enough contiguous space to
			// accommodate npages.
			p.searchAddr = maxSearchAddr
		}
		return 0, 0
	}
Found:
	// Go ahead and actually mark the bits now that we have an address.
	scav = p.allocRange(addr, npages)

	// If we found a higher searchAddr, we know that all the
	// heap memory before that searchAddr in an offset address space is
	// allocated, so bump p.searchAddr up to the new one.
	if p.searchAddr.lessThan(searchAddr) {
		p.searchAddr = searchAddr
	}
	return addr, scav
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

func sysAlloc(n uintptr) unsafe.Pointer {
	// 140339408883712
	var fd int = -1
	var v uintptr = 1649267441664
	p, _, eno := syscall.Syscall6(syscall.SYS_MMAP, uintptr(v), n, syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_ANON|syscall.MAP_PRIVATE, uintptr(fd), 0)
	if eno != 0 {
		panic(eno)
	}

	// m[uintptr(p)] = &vma{uintptr(p), n}
	// fmt.Printf("malloced(%d) -end(%v)\n", uintptr(p), p+n*8)
	return unsafe.Pointer(p)
}

func sysMap(v unsafe.Pointer, n uintptr) unsafe.Pointer {
	var fd int = -1
	p, _, eno := syscall.Syscall6(syscall.SYS_MMAP, uintptr(v), n,
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_ANON|syscall.MAP_FIXED|syscall.MAP_PRIVATE,
		uintptr(fd), 0)
	if eno != 0 {
		panic(eno)
	}
	// fmt.Printf("want(%v) malloc(%v)\n", uintptr(v), p)
	// m[uintptr(v)] = &vma{p, n}
	return unsafe.Pointer(p)
}

func sysReserve(v unsafe.Pointer, n uintptr) unsafe.Pointer {
	var fd int = -1
	p, _, eno := syscall.Syscall6(syscall.SYS_MMAP, uintptr(v), n, syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_ANON|syscall.MAP_PRIVATE, uintptr(fd), 0)
	if eno != 0 {
		panic(eno)
	}

	// fmt.Printf("reserved(%d) -end(%v) n(%v)\n", uintptr(p), p+n*8, n)
	return unsafe.Pointer(p)
}
