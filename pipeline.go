// Copyright 2022 clavinjune/piper
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package piper

import (
	"context"
	"sync"
)

const (
	defaultTotalWorker = 1
)

var (
	defaultContext = context.Background()
)

// P is the pipeline
type P[IN, OUT any] struct {
	_           struct{}
	filterError bool
	ctx         context.Context
	totalWorker int
	data        *sync.Map
	fn          F[IN, OUT]
}

// Do executes the pipeline
func (p *P[IN, OUT]) Do(ins ...<-chan *W[IN]) <-chan *W[OUT] {
	in := firstOrMerge(ins...)
	out := make(chan *W[OUT])

	wg := new(sync.WaitGroup)
	wg.Add(p.totalWorker)

	go func() {
		for i := 0; i < p.totalWorker; i++ {
			go p.work(wg, in, out)
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func (p *P[IN, OUT]) work(wg *sync.WaitGroup, in <-chan *W[IN], out chan<- *W[OUT]) {
	defer wg.Done()

ChannelReader:
	for n := range in {
		select {
		case <-p.ctx.Done():
			break ChannelReader
		default:
			p.workDefaultAction(n, out)
		}
	}
}

func (p *P[IN, OUT]) workDefaultAction(n *W[IN], out chan<- *W[OUT]) {
	if n.Err != nil {
		if p.filterError {
			return
		}
		out <- &W[OUT]{
			Err: n.Err,
		}
	}

	o, err := p.fn(&M[IN]{
		Map: p.data,
		Ctx: p.ctx,
		In:  n.Data,
	})

	if err != nil && p.filterError {
		return
	}

	out <- &W[OUT]{
		Data: o,
		Err:  err,
	}
}

// New creates new pipeline
func New[IN, OUT any](fn F[IN, OUT], opts ...*Opt) *P[IN, OUT] {
	p := &P[IN, OUT]{
		fn: fn,
	}

	if len(opts) == 0 {
		return p
	}

	p.fill(opts[0])
	if p.totalWorker == 0 {
		p.totalWorker = defaultTotalWorker
	}

	if p.ctx == nil {
		p.ctx = defaultContext
	}

	return p
}

func (p *P[IN, OUT]) fill(o *Opt) {
	if o == nil {
		return
	}

	p.ctx = o.Context
	p.totalWorker = o.TotalWorker
	p.filterError = o.FilterError

	m := new(sync.Map)
	for k, v := range o.Data {
		m.Store(k, v)
	}

	p.data = m
}
