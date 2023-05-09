package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var balance int64
var mutex = &sync.Mutex{}

func say(s string, times int) {
	for i := 0; i < times; i++ {
		// inject a 100 ms delay
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, s)
	}
}

func credit(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		mutex.Lock()
		atomic.AddInt64(&balance, 100)
		fmt.Println("After crediting, balance is", balance)
		mutex.Unlock()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	// notify the WaitGroup when we are done
	defer wg.Done()
}

func debit(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		atomic.AddInt64(&balance, -100)
		fmt.Println("After the debiting, balance is", balance)
		mutex.Unlock()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	// notify the WaitGroup when we are done
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	balance = 200

	fmt.Println("Initial balance is", balance)
	wg.Add(1)
	go credit(&wg)
	wg.Add(1)
	go debit(&wg)
	wg.Wait()
	fmt.Println("Final balance is", balance)
}
