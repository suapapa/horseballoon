// Copyright 2020 Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
)

func main() {
	flag.IntVar(&recordBuffSize, "b", sampleRate /* 64 */, "record buffer size")
	flag.IntVar(&quiteThreshhold, "t", -30, "theshold for turn to onair")
	flag.Parse()

	log.Println("press ctrl-c to stop.")

	var wg sync.WaitGroup
	ctx := context.Background()
	ctx, cancle := context.WithCancel(ctx)

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
