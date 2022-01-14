package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	fmt.Println("start.", debug.SetGCPercent(100))

	container := make([][]byte, 0, 8)
	fmt.Println("> loop.")
	for i := 0; ; i++ {
		if len(container) < 32 {
			container = append(container, make([]byte, (1<<20)*4))
		} else {
			tp := make([][]byte, len(container)-1)
			copy(tp, container)
			container = tp
		}

		time.Sleep(time.Microsecond * 100)
	}

	// for i := 0; i < 10; i++ {
	// 	runtime.GC()
	// }
	fmt.Println("< loop.")
}

func getgoal(last, start, end, goal float32) {
	goalrate := goal/last - 1
	triggerrate := start/last - 1
	realrate := end/last - 1
	delta := (goal - start) - (end - start)
	deltarate := (goalrate - triggerrate) - (realrate - triggerrate)
	nxrate := triggerrate + deltarate*0.5
	nxgoal := goal * 2
	fmt.Println(delta, deltarate, nxrate, nxgoal, nxgoal*nxrate)
}
