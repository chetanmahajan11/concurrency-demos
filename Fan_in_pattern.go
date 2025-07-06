package main

// import (
// 	"fmt"
// 	"sync"
// )

// /*
// 	Fan-in Pattern:
// 	merges data from two separate channels into one channel
// 	This example merges the data from two separate channels into one ooutput channel
// */

// func main() {
// 	chOdd := make(chan int)
// 	chEven := make(chan int)

// 	// generate data
// 	go func() {
// 		defer close(chEven)
// 		for i := 0; i < 10; i += 2 {
// 			chEven <- i
// 		}
// 	}()

// 	go func() {
// 		defer close(chOdd)
// 		for i := 1; i < 10; i += 2 {
// 			chOdd <- i
// 		}
// 	}()

// 	// merge from two channels
// 	finalOutput := merge(chOdd, chEven)

// 	// print final output
// 	for val := range finalOutput {
// 		fmt.Println(val)
// 	}
// }

// // merges two channels and returns merged channel
// func merge(ch1, ch2 <-chan int) chan int {
// 	out := make(chan int)
// 	var wg sync.WaitGroup

// 	forward := func(ch <-chan int) {
// 		defer wg.Done()
// 		for val := range ch {
// 			out <- val
// 		}
// 	}

// 	wg.Add(2)
// 	go forward(ch1)
// 	go forward(ch2)

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()
// 	return out
// }
