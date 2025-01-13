package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	// FAN OUT
	// Multiple functions reading from the same channel until it is closed.
	// Distribute work across multiple function (ten goroutines) that all read from in.
	xc :=fanOut(in, 10)

	// FAN IN
	// Multiplex multiple channels onto a single channel.
	// Merge the channels fro c0 through c9 into a single channel.

	for n := range merge(xc...) {
		fmt.Println(n)
	}
	fmt.Println("---------- Finish!----------")
}

func gen() <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			for j := 1; j < 10; j++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func fanOut(in <-chan int, n int) []<-chan int {
	xc := make([]<-chan int, n)
	for i := 0; i < n; i++ {
		xc = append(xc, factorial(in))
	}
	return xc
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are done.
	// This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

/*
CHALLENGE #1
-- This code throws an error: fatal error: all goroutines are asleep  - deadlock!
-- fix this code
*/