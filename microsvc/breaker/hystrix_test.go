package breaker

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

func Test_newHystrixBreaker(t *testing.T) {
	bk := newHystrixBreaker("test")

	once := func() {
		for i := 0; i < 3000; i++ {
			i := i
			er := bk.Do(func() error {
				e := normal(i)
				if e == nil {
					log.Println(i, "normal")
				}
				return e
			}, func(e error) error {
				er := onBreakerOpen(i, e)
				if e == hystrix.ErrCircuitOpen {
					time.Sleep(time.Millisecond * 100)
				}
				return er
			})
			if er != nil {
				log.Println(er)
			}
		}
	}

	once()

	time.Sleep(time.Second)
}

var lastNormal = time.Now()
var normalCnt int64

func normal(i int) error {
	// 10 qps
	if time.Since(lastNormal) < time.Second {
		if normalCnt < 10 {
			normalCnt++
			return nil
		}
		return fmt.Errorf("cnt exceed during 1s, %v", i)
	}

	lastNormal = time.Now()
	normalCnt = 1
	return nil
}

func onBreakerOpen(i int, er error) error {
	log.Println(i, er)
	return nil
}
