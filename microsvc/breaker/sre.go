package breaker

import (
	"time"

	"github.com/go-kratos/aegis/circuitbreaker"
	"github.com/go-kratos/aegis/circuitbreaker/sre"
)

type sreBreaker struct {
	circuitbreaker.CircuitBreaker
}

func newSreBreaker(name string) *sreBreaker {
	bk := sre.NewBreaker(
		sre.WithWindow(time.Second * 10),
	)
	return &sreBreaker{CircuitBreaker: bk}
}

func (b *sreBreaker) Do(fn func() error, onBreakerOpen func(error) error) error {
	er := b.Allow()

	if er == nil {
		er = fn()
		if er == nil {
			b.MarkSuccess()
			return nil
		}
		b.MarkFailed()
		return er
	}

	return onBreakerOpen(er)
}
