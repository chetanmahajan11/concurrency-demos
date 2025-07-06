package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg    sync.WaitGroup
		turns = make(chan string)
		num   = 5
	)

	wg.Add(2)

	// ping goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < num; i++ {
			<-turns
			fmt.Println("ping")
			if i+1 < num {
				turns <- "pong"
			}
		}
	}()

	// pong goroutine
	go func() {
		defer wg.Done()
		for j := 0; j < num; j++ {
			<-turns
			fmt.Println("pong")
			turns <- "ping"

		}
	}()

	// start counter
	turns <- "ping"
	wg.Wait()
	close(turns)
}
