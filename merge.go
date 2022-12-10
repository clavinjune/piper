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

import "sync"

// Merge merges channels to one
func Merge[T any](chs ...<-chan *W[T]) <-chan *W[T] {
	out := make(chan *W[T])
	wg := new(sync.WaitGroup)
	wg.Add(len(chs))

	work := func(wg *sync.WaitGroup, ch <-chan *W[T], o chan<- *W[T]) {
		defer wg.Done()
		for c := range ch {
			o <- c
		}
	}

	for i := range chs {
		go work(wg, chs[i], out)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// firstOrMerge helps to avoid merging operation if there's only 1 input
func firstOrMerge[T any](chs ...<-chan *W[T]) <-chan *W[T] {
	if len(chs) == 1 {
		return chs[0]
	}

	return Merge(chs...)
}
