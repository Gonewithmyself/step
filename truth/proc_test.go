package truth

import (
	"sync"
	"testing"
)

func Test_class(t *testing.T) {
	classifyManyPlayers()
}

func BenchmarkClassify(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			classifyManyPlayers()
		}
	})
}

func classifyManyPlayers() {
	pcount := 1000
	var wg sync.WaitGroup
	for i := 0; i < pcount; i++ {
		wg.Add(1)
		go func() {
			classifyPlayerItem()
			wg.Done()
		}()
	}
	wg.Wait()
}

func classifyPlayerItem() {
	for i := 0; i < 1000; i++ {

	}
}
