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

// To converts back the channel of *W[T] to a list
func To[T any](chs ...<-chan *W[T]) []*W[T] {
	ch := firstOrMerge(chs...)
	result := make([]*W[T], 0)

	for c := range ch {
		result = append(result, c)
	}

	return result
}

// ToData converts back the channel of *W[T] to a list only containing the non-error
func ToData[T any](chs ...<-chan *W[T]) []T {
	ch := firstOrMerge(chs...)
	result := make([]T, 0)

	for c := range ch {
		if c.Err != nil {
			continue
		}

		result = append(result, c.Data)
	}

	return result
}
