package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/luckrats/pubchannle"
)

var (
	RECEIVER_COUNT = int(1e4)
	MSG_COUNT      = int(1e4)
)

func main() {
	rc := flag.Int("rc", 1e4, "receive count")
	mc := flag.Int("mc", 1e4, "message count")
	RECEIVER_COUNT = *rc
	MSG_COUNT = *mc
	fmt.Println("receive count:", RECEIVER_COUNT)
	fmt.Println("msg count:", MSG_COUNT)
	p := pubchannle.NewPublishChannle()

	wg := &sync.WaitGroup{}
	for i := 0; i < RECEIVER_COUNT; i++ {
		Receive_Process(wg, p.NewSubscribChannle())
	}

	fmt.Println("Send start")
	st := time.Now()
	for i := 0; i < MSG_COUNT; i++ {
		p.Write(i)
	}
	fmt.Println("Send ok, use", time.Since(st))

	p.Close()
	wg.Wait()
	fmt.Println("Receive ok, use", time.Since(st))
}

func Receive_Process(wg *sync.WaitGroup, r *pubchannle.SubscribChannle) {
	wg.Add(1)
	go func() {
		count := 0
		defer wg.Done()
	exit:
		for {
			select {
			case <-r.WaitNotify():
				d, ok := r.Read()
				if !ok {
					break exit
				}
				if d != nil {
					_ = d
				} else {
					fmt.Println(count, d)
				}
				count++
			}
		}
	}()
}
