package sort

import "testing"

func Test_numK(t *testing.T) {
	data := []int32{13, 5, 8, 20, 7}
	numK(data, 3)
	t.Log(data)
}
