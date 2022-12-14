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
	"fmt"
)

// W wraps the data passed on each pipeline so user can handle the Error
type W[T any] struct {
	_    struct{}
	Data T
	Err  error
}

// Unwrap is implemented to satisfy errors.Unwrap
func (w *W[T]) Unwrap() error {
	return w.Err
}

// Error is implemented to satisfy error interface
func (w *W[T]) Error() string {
	if w.Err != nil {
		return w.Err.Error()
	}
	return ""
}

// String is implemented to satisfy fmt.Stringer interface
func (w *W[T]) String() string {
	if w.Err != nil {
		return ""
	}
	return fmt.Sprintf("%#v", w.Data)
}
