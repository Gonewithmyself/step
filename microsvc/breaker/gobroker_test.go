package breaker

import (
	"log"
	"testing"
	"time"

	"github.com/sony/gobreaker"
)

func Test_newSonyBreaker(t *testing.T) {
	bk := newSonyBreaker("test")

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
				if e == gobreaker.ErrOpenState {
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
