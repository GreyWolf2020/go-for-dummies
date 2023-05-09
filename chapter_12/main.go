package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(n int, c chan int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		c <- a // block until value is received from channel
		a, b = b, a+b
		time.Sleep(1 * time.Second)
	}
	close(c) // close the channel
}

func counter(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
		time.Sleep(2 * time.Second)
	}
	close(c)
}

//---send data into a channel---
func sendData(ch chan string) {
	fmt.Println("Send a string into channel ...")
	// comment out the following line
	// time.Sleep(2 * time.Second)
	ch <- "Hello"
	fmt.Println("String has been retrieved from the channel...")
}

//---getting data from channel---
func getData(ch chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("String retrieved from a channel:", <-ch)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
	fmt.Println("Done and can proceed.")
}

func main() {
	ch := make(chan string)
	c := make(chan int, 5)
	d1 := make(chan int)
	d2 := make(chan int)
	s := []int{}
	sliceSize := 10
	partSize := 2
	parts := sliceSize / partSize
	rand.Seed(time.Now().UnixNano())

	go sendData(ch)
	go getData(ch)

	for i := 0; i < sliceSize; i++ {
		s = append(s, rand.Intn(100))
	}
	for i := 0; i < parts; i++ {
		go sum(s[i*partSize:(i+1)*partSize], c)
	}
	total := 0
	for i := 0; i < parts; i++ {
		partialSum := <-c // read the channel
		fmt.Println("Partial Sum: ", partialSum)
		total += partialSum
	}
	fmt.Println("Total: ", total)

	go fib(10, d1)
	go counter(10, d2)

	d1Closed := false
	d2Closed := false

	go func() {
		for {
			select {
			case n, ok := <-d1:
				if !ok {
					// channel closed and drained
					d1Closed = true
					if d1Closed && d2Closed {
						return
					}
				} else {
					fmt.Println("fin()", n)
				}
			case m, ok := <-d2:
				if !ok {
					// channel closed and drained
					d2Closed = true
					if d2Closed && d1Closed {
						return
					}
				} else {
					fmt.Println("counter()", m)
				}

			}
		}
	}()

	fmt.Println("Continue doing something that you wanted to do on your own.")

	fmt.Scanln()
}
