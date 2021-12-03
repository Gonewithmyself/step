package truth

import (
	"runtime"
	"time"
)

func tm() {
	time.NewTimer(time.Hour)

	time.AfterFunc(time.Second, func() {})
	runtime.Gosched()
}
