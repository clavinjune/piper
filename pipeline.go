package piper

import (
	"context"
	"sync"
)

type P[IN, OUT any] struct {
	ctx         context.Context
	totalWorker int
	data        map[string]any
	fn          F[IN, OUT]
}

func (p *P[IN, OUT]) Do(in <-chan IN) <-chan OUT {
	out := make(chan OUT)

	wg := new(sync.WaitGroup)
	wg.Add(p.totalWorker)

	go func() {
		for i := 0; i < p.totalWorker; i++ {
			go func() {
				defer wg.Done()
				for n := range in {
					select {
					case <-p.ctx.Done():
						break
					default:
						out <- p.fn(&M[IN]{
							Ctx:  p.ctx,
							In:   n,
							Data: p.data,
						})
					}
				}
			}()
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func New[IN, OUT any](ctx context.Context, totalWorker int, data map[string]any, fn F[IN, OUT]) *P[IN, OUT] {
	return &P[IN, OUT]{
		ctx:         ctx,
		totalWorker: totalWorker,
		data:        data,
		fn:          fn,
	}
}
