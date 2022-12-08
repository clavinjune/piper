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
	return m.In + 1, nil
}

func main() {
	ch, total := piper.From(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	totalWorker := total / 2

	start := time.Now()
	out := piper.New(context.Background(), totalWorker, nil, addOne).Do(ch)
	out2 := piper.New(context.Background(), totalWorker, nil, timesTwo).Do(out)

	for o := range out2 {
		log.Println(o.Data, o.Err)
	}

	log.Println(time.Since(start))
}
