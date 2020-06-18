package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
)

type ctxKey struct{}

var (
	recordBuffSize  int
	quiteThreshhold int
)

func main() {
	flag.IntVar(&recordBuffSize, "b", sampleRate /* 64 */, "record buffer size")
	flag.IntVar(&quiteThreshhold, "t", -30, "theshold for turn to onair")
	flag.Parse()

	log.Println("press ctrl-c to stop.")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	var wg sync.WaitGroup
	ctx := context.Background()
	ctx, cancle := context.WithCancel(ctx)
	wg.Add(1)
	go record(ctx, &wg)

	<-sig
	cancle()
	wg.Wait()
	log.Println("bye")
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
