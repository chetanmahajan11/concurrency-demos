package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Producing", i, ", time:", time.Now())
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for val := range ch {
		fmt.Println("Consumed", val, ", time:", time.Now())
	}
}

// func main() {
// 	ch := make(chan string)

// 	go func() {
// 		time.Sleep(time.Second * 2)
// 		ch <- "completed execution !!"
// 	}()

// 	select {
// 	case val := <-ch:
// 		fmt.Println(val)

// 	case <-time.After(time.Second * 1):
// 		fmt.Println("Timeout occurred !!")

// 	default:
// 		fmt.Println("Nothing received !!")

// 	}

// }
