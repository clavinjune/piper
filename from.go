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

// From converts list of T to channel of *W[T] and returns the length
func From[T any](t ...T) (<-chan *W[T], int) {
	out := make(chan *W[T])
	go func() {
		defer close(out)

		for _, i := range t {
			out <- &W[T]{
				Data: i,
			}
		}
	}()

	return out, len(t)
}

// FromMap converts map of <K, T> to channel of *W[T] and returns the length
func FromMap[K comparable, T any](m map[K]T) (<-chan *W[T], int) {
	out := make(chan *W[T])
	go func() {
		defer close(out)

		for _, i := range m {
			out <- &W[T]{
				Data: i,
			}
		}
	}()

	return out, len(m)
}
