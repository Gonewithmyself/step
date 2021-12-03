package truth

// go tool compile -S fncall.go

func Add(x, y int) int {
	z := y + 3
	return x + z
}

func Sub(x, y int) int {
	return x - y
}

func A(x, y int) int {
	z := Add(x, y)
	return Sub(z, y)
}
