package main

import (
	"fmt"
	"step/asm/pkg"
)

func main() {
	a, b := pkg.Swap(1, 2)
	fmt.Println(pkg.Num, pkg.Str, a, b)
}
