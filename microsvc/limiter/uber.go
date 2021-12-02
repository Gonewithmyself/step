package limiter

import "go.uber.org/ratelimit"

type uber struct {
	ratelimit.Limiter
}

func newUber(rate, cap int) *uber {
	return &uber{
		Limiter: ratelimit.New(rate, ratelimit.WithSlack(cap)),
	}
}

func (l *uber) Allow() bool {
	l.Take()
	return true
}
