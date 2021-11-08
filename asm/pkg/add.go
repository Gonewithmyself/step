package pkg

func Add(a, b int64) int64 {
	e := a - b
	if a < b {
		e = b - a
	}
	return e + a
}

// func foo(a, b int64) (c, d int64) {
// 	c = Add(a, b)
// 	d = a + b
// 	return
// }

// func swap(a, b int64) (c, d int64) {
// 	return b, a
// }
