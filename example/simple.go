package main

import (
	"context"
	"log"
	"time"

	"github.com/clavinjune/piper"
)

func addOne(m *piper.M[int]) int {
	time.Sleep(time.Second)
	return m.In + 1
}

func main() {
	ch, total := piper.From(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	totalWorker := total / 2

	start := time.Now()
	out := piper.New(context.Background(), totalWorker, nil, addOne).Do(ch)

	for o := range out {
		log.Println(o)
	}

	log.Println(time.Since(start))
}
