package truth

import (
	"fmt"
	"net"
	"runtime"
	"syscall"
	"time"
	"unsafe"
)

func syscalleg() {
	runtime.Gosched()

	runtime.GOMAXPROCS(1)

	fd, _, en := syscall.Syscall(syscall.SYS_SOCKET, syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if en != 0 {
		panic(en)
	}
	// syscall.SetNonblock(int(fd), true)
	syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.0:9999")
	if err != nil {
		panic(err)
	}
	sa := &syscall.SockaddrInet4{Port: tcpAddr.Port}
	copy(sa.Addr[:], tcpAddr.IP)
	er := syscall.Bind(int(fd), sa)
	if er != nil {
		panic(er)
	}

	er = syscall.Listen(int(fd), 10)
	if er != nil {
		panic(er)
	}

	ra := syscall.SockaddrInet4{}
	ralen := unsafe.Sizeof(ra)
	pra := unsafe.Pointer(&ra)
	plen := unsafe.Pointer(&ralen)

	go func() {
		<-time.After(time.Second)
		fmt.Println("time's up")
	}()

	//cfd, _, eno := syscall.Syscall(syscall.SYS_ACCEPT, fd, uintptr(pra), uintptr(plen))
	cfd, _, eno := syscall.RawSyscall(syscall.SYS_ACCEPT, fd, uintptr(pra), uintptr(plen))
	fmt.Println(cfd, eno)

	time.Sleep(time.Second * 60)
}
