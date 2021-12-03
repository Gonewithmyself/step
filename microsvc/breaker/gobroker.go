package breaker

import (
	"log"
	"time"

	"github.com/sony/gobreaker"
)

type sonyBreaker struct {
	*gobreaker.CircuitBreaker
}

func newSonyBreaker(name string) *sonyBreaker {
	opts := gobreaker.Settings{
		Name:        name,
		MaxRequests: 3,
		Interval:    time.Second * 10,
		Timeout:     time.Millisecond * 500,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			var rate float64
			defer func() {
				log.Printf("counts: total(%v) fail(%v) rate(%v)\n",
					counts.Requests, counts.TotalFailures, rate)
			}()
			total := counts.Requests
			if total < 10 {
				return true
			}
			rate = float64(counts.TotalFailures) / float64(total)
			return rate-0.2 < 0
		},
		OnStateChange: func(name string, from, to gobreaker.State) {
			log.Printf("state change from(%v) to(%v)\n", from, to)
		},
	}

	bk := gobreaker.NewCircuitBreaker(opts)
	return &sonyBreaker{
		CircuitBreaker: bk,
	}
}

func (b *sonyBreaker) Do(fn func() error, onBreakerOpen func(error) error) error {
	_, er := b.Execute(func() (interface{}, error) {
		er := fn()
		return nil, er
	})
	if er != nil {
		onBreakerOpen(er)
		return nil
	}

	return nil
}
