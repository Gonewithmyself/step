package breaker

import (
	"time"

	"github.com/go-kratos/aegis/circuitbreaker"
	"github.com/go-kratos/aegis/circuitbreaker/sre"
)

type sreBreaker struct {
	circuitbreaker.CircuitBreaker
}

func newSreBreaker() *sreBreaker {
	bk := sre.NewBreaker(
		sre.WithWindow(time.Second),
	)
	return &sreBreaker{CircuitBreaker: bk}
}
