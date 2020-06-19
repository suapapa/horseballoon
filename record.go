package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math"
	"sync"
	"time"

	"github.com/gordonklaus/portaudio"
)

const (
	sampleRate = 16000
)

func record(ctx context.Context, wg *sync.WaitGroup) {
	defer func() {
		log.Println("end recordContext")
		wg.Done()
	}()

	portaudio.Initialize()
	defer portaudio.Terminate()
	time.Sleep(1 * time.Second)
	log.Print("recording...")

	in := make([]int16, recordBuffSize)                                       // signed 16bit
	stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(in), in) // 16000Hz
	chk(err)
	defer stream.Close()
	chk(stream.Start())

	sumSlice := func(vs []int16) float64 {
		var s float64
		for _, v := range vs {
			s += (float64(v) * float64(v))
		}
		return s
	}

	for {
		// wait for voice input
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			// calc dB for 1 sec
			var sum float64
			for i := 0; i < sampleRate/recordBuffSize; i++ {
				stream.Read() // chk(stream.Read())
				sum += sumSlice(in)
			}
			rms := math.Sqrt(sum / sampleRate)
			dB := 20 * math.Log10(rms/0x7fff) // 0x7fff: max value for 16 bit
			fmt.Printf("\rwaiting... (db: %.02f)", dB)
			if dB > float64(quiteThreshhold) {
				break
			}
		}

		log.Println("!! on air !!")
		pR, pW := io.Pipe()
		go func() {
			err := kakaoSTT(ctx, pR)
			if err != nil {
				log.Printf("fail to stt: %v", err)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			err := binary.Write(pW, binary.LittleEndian, in) // little endian
			if err != nil {
				// log.Printf("err in writing to pipe: %v", err)
				log.Println("!! off air !!")
				break
			}
			stream.Read() // chk(stream.Read())
		}
	}
}
