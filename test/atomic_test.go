package test

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Rectangle struct {
	length int
	width  int
}

func TestAtomic(t *testing.T) {

	var rect atomic.Value

	update := func(width, length int) {
		rectLocal := new(Rectangle)
		rectLocal.width = width
		rectLocal.length = length
		rect.Store(rectLocal)
	}

	wg := sync.WaitGroup{}
	wg.Add(10)
	// 10 个协程并发更新
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			update(i, i+5)
		}()
	}
	wg.Wait()
	_r := rect.Load().(Rectangle)
	t.Logf("rect.width=%d\nrect.length=%d\n", _r.width, _r.length)
}
