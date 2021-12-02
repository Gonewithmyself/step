package limiter

import (
	"fmt"
	"testing"
	"time"
)

func Test_newXtLimiter(t *testing.T) {
	l := newXtLimiter(10, 20)

	try := func(i int) {
		start := time.Now()
		x := 0
		for time.Since(start).Seconds() < 1 {
			if l.Allow() {
				x++
				fmt.Println(time.Since(start).Milliseconds(), x, i)
			}
		}

	}
	for i := 0; i < 3; i++ {
		try(i)
	}

	time.Sleep(time.Second * 1)
	try(500)

}
