package breaker

import (
	"log"
	"testing"
	"time"

	"github.com/go-kratos/aegis/circuitbreaker"
)

func Test_newSreBreaker(t *testing.T) {
	bk := newSreBreaker("test")

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
				if e == circuitbreaker.ErrNotAllowed {
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
