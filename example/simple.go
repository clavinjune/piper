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

// Package main is an example on how to use piper
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/clavinjune/piper"
)

func add65(m *piper.M[int]) (int, error) {
	name, _ := m.Load("name")
	log.Printf("%s: add65(%d)", name, m.In)
	return m.In + 65, nil
}

func toRune(m *piper.M[int]) (rune, error) {
	name, _ := m.Load("name")
	log.Printf("%s: toRune(%d)", name, m.In)
	return rune(m.In), nil
}

func toString(m *piper.M[rune]) (string, error) {
	name, _ := m.Load("name")
	log.Printf("%s: toString(%#q)", name, m.In)

	if m.In > 'Z' {
		return "", fmt.Errorf("piper example: alphabet out of range %#q", m.In)
	}

	return string(m.In), nil
}

func main() {
	// generating 20 - 26
	nums := make([]int, 7)
	for i := range nums {
		nums[i] = i + 20
	}

	// converts nums to a channel
	ch, total := piper.From(nums...)

	// use 1/3 of the total data as a worker
	totalWorker := total / 3

	ctx := context.Background()
	m := map[string]any{
		"name": "piper simple example",
	}

	popt := &piper.Opt{
		Context:     ctx,
		Data:        m,
		TotalWorker: totalWorker,
		FilterError: true,
	}
	// create pipeline functions
	addOneFn := piper.New(add65, popt)
	toRunFn := piper.New(toRune, popt)
	toStringFn := piper.New(toString, popt)

	start := time.Now()

	// run pipeline functions
	addOneCh := addOneFn.Do(ch)

	// Fan-Out
	toRuneCh1 := toRunFn.Do(addOneCh)
	toRuneCh2 := toRunFn.Do(addOneCh)

	// Fan-In the toRuneCh
	// Fan-Out the toStringCh
	toStringCh1 := toStringFn.Do(toRuneCh1, toRuneCh2)
	toStringCh2 := toStringFn.Do(toRuneCh1, toRuneCh2)

	// Fan-In
	result := piper.To(toStringCh1, toStringCh2)
	for _, r := range result {
		log.Println(r.String(), r.Error())
	}

	// to get the non-error result
	// result := piper.ToData(toStringCh1, toStringCh2)
	// for _, r := range result {
	//	log.Println(r)
	// }

	log.Println(time.Since(start))
}
