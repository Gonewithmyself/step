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
		bubbleSort,
		quickSort,
		quickSortTwoPointer,
		mergeSort,
		heapSort,
		countSort,
	}

	for i := 0; i < 10; i++ {
		arr := genarray()
		for j := range fns {
			tosort := make([]int, len(arr))
			copy(tosort, arr)
			cost := wrappSort(tosort, fns[j])
			if !sort.IntsAreSorted(tosort) {
				t.Error(j, reflect.TypeOf(fns[j]).Name())
			}
			_ = cost
			// t.Log(reflect.TypeOf(fns[j]).Name(), cost, len(tosort))
		}
	}
}

func wrappSort(arr []int, fn func(arr []int)) int64 {
	start := time.Now().UnixNano()
	fn(arr)
	return time.Now().UnixNano() - start
}

func genarray() []int {
	n := rand.Intn(10) + 3
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}
	return arr
}

func init() {
	rand.Seed(time.Now().Unix())
}

func Test_qsort(t *testing.T) {
	src := []int{8, 10, 2, 3, 6, 1, 5}
	countSort(src)
	t.Log(src)
}

func BenchmarkQsort(b *testing.B) {
	src := []int{8, 10, 2, 3, 6, 1, 5}
	for i := 0; i < b.N; i++ {
		quickSort(src)
	}
}

func BenchmarkQsortTwo(b *testing.B) {
	src := []int{8, 10, 2, 3, 6, 1, 5}
	for i := 0; i < b.N; i++ {
		quickSortTwoPointer(src)
	}
}
