package main

import (
	"sync"
)

func TestAppend() (result []int) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		v := i
		wg.Add(1)
		go func() {
			// other logic
			result = append(result, v)
			wg.Done()
		}()
	}

	wg.Wait()
	//fmt.Printf("%v\n", len(result))
	return result
}

type baseInfo struct {
	head int32
	lv   int32
}
type user struct {
	id   int64
	base *baseInfo
}
type player struct {
	id   int64
	base baseInfo
}

func main() {
	for a := 0; a < 100000; a++ {
		res := TestAppend()
		println("len(res):", len(res))
	}
}
