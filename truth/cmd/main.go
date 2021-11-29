package main

import (
	"fmt"
	"unsafe"
)

func main() {
	buf := make([]byte, 10240)
	base := 824633720832
	start := uintptr(unsafe.Pointer(&buf[0]))
	fmt.Printf("%d diff %d \n", &buf[0], (start-uintptr(base))/1024)
	// time.Sleep(time.Minute)
}
