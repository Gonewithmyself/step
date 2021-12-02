package limiter

import "golang.org/x/time/rate"

type xtlimiter struct {
	*rate.Limiter
}

func newXtLimiter(count, extral int) *xtlimiter {
	return &xtlimiter{
		Limiter: rate.NewLimiter(rate.Limit(count), extral),
	}
}
