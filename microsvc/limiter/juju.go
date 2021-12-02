package limiter

import "github.com/juju/ratelimit"

type julimiter struct {
	*ratelimit.Bucket
}

func newJulimiter(rate, cap int) *julimiter {
	return &julimiter{
		Bucket: ratelimit.NewBucketWithRate(float64(rate), int64(cap)),
	}
}

func (l *julimiter) Allow() bool {
	return l.TakeAvailable(1) != 0
}
