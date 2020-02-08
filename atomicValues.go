package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var ort1 time.Duration
	var ort2 time.Duration
	var ort3 time.Duration
	var syncer sync.WaitGroup

	divider := 100
	syncer.Add(1)
	go func() {

		for i := 0; i < divider; i++ {
			c := mutexFunc(50, 100)
			ort1 += c
		}
		syncer.Done()
	}()
	syncer.Add(1)
	go func() {

		for i := 0; i < divider; i++ {
			t := atomicFunc(50, 100)
			ort2 += t
		}
		syncer.Done()

	}()
	syncer.Add(1)
	go func() {
		for i := 0; i < divider; i++ {
			e := channelFunc(50, 100)
			ort3 += e
		}
		syncer.Done()
	}()
	syncer.Wait()

	fmt.Println("\n\n")
	fmt.Printf("Mutex:%v\n", ort1/time.Duration(divider))
	fmt.Printf("Atomic:%v\n", ort2/time.Duration(divider))
	fmt.Printf("Channel:%v\n", ort3/time.Duration(divider))

}

func mutexFunc(first, second int) time.Duration {
	var ops uint64

	var wg sync.WaitGroup

	var mtx sync.Mutex

	t := time.Now()

	for i := 0; i < first; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < second; c++ {
				mtx.Lock()
				ops = ops + 1
				mtx.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	elp := time.Since(t)
	fmt.Printf("Mutex Func %v \n", elp)

	fmt.Println("ops:", ops)
	return elp
}

func atomicFunc(first, second int) time.Duration {
	var ops uint64

	var wg sync.WaitGroup
	t := time.Now()

	for i := 0; i < first; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < second; c++ {

				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	elp := time.Since(t)
	fmt.Printf("atomic Func %v\n", elp)
	fmt.Println("ops:", ops)
	return elp
}

func channelFunc(first, second int) time.Duration {
	var ops uint64
	ch := make(chan uint64, first*second)

	go func() {
		for {
			select {
			case value := <-ch:
				ops = ops + value
			}
		}
	}()
	var wg sync.WaitGroup
	t := time.Now()

	for i := 0; i < first; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < second; c++ {

				ch <- 1
			}
			wg.Done()
		}()
	}

	wg.Wait()
	elp := time.Since(t)
	fmt.Printf("Channel Func %v\n", elp)
	fmt.Println("ops:", ops)
	return elp
}
