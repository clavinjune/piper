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
	"log"
	"time"

	"github.com/clavinjune/piper"
)

func addOne(m *piper.M[int]) (int, error) {
	time.Sleep(time.Second)
	return m.In + 1, nil
}

func timesTwo(m *piper.M[int]) (int, error) {
	time.Sleep(time.Second)
	return m.In * 2, nil
}

func main() {
	ch, total := piper.From(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	totalWorker := total / 2

	start := time.Now()
	addOnCh := piper.New(context.Background(), totalWorker, nil, addOne).Do(ch)
	timesTwoCh := piper.New(context.Background(), totalWorker, nil, timesTwo).Do(addOnCh)
	result := piper.ToData(timesTwoCh)

	for _, r := range result {
		log.Println(r)
	}

	log.Println(time.Since(start))
}
