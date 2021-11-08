package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func Test_Sort(t *testing.T) {
	fns := []func([]int){
		insertSort,
	}

	for i := 0; i < 10; i++ {
		arr := genarray()
		for j := range fns {
			tosort := make([]int, len(arr))
			copy(tosort, arr)
			cost := wrappSort(tosort, fns[j])
			if !sort.IntsAreSorted(tosort) {
				t.Error(reflect.TypeOf(fns[j]).Name())
			}
			t.Log(reflect.TypeOf(fns[j]).Name(), cost, len(tosort))
		}
	}
}

func wrappSort(arr []int, fn func(arr []int)) int64 {
	start := time.Now().UnixNano()
	fn(arr)
	return time.Now().UnixNano() - start
}

func genarray() []int {
	n := rand.Intn(5000) + 3000
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}
	return arr
}

func init() {
	rand.Seed(time.Now().Unix())
}
