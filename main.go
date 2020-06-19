package main

import (
	"context"
	"flag"
	"log"
	"sync"
)

type ctxKey struct{}

var (
	recordBuffSize  int
	quiteThreshhold int

	// guiCh chan string
)

func main() {
	flag.IntVar(&recordBuffSize, "b", sampleRate /* 64 */, "record buffer size")
	flag.IntVar(&quiteThreshhold, "t", -30, "theshold for turn to onair")
	flag.Parse()

	log.Println("press ctrl-c to stop.")

	var wg sync.WaitGroup
	ctx := context.Background()
	ctx, cancle := context.WithCancel(ctx)
	// guiCh = make(chan string, 10)
	wg.Add(1)
	go record(ctx, &wg)
	gui()

	cancle()
	wg.Wait()
	log.Println("bye")
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
