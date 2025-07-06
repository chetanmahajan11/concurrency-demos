package main

// import (
// 	"fmt"
// 	"sync"
// )

/*
	Key summary:
	- For some unknown reason, 'pong' goroutine is the first picking up the signal from the goroutine.
	   So putting the if-condition in the 'ping' channel to limit over-sending in the channel.
*/

// func main() {
// 	var (
// 		wg    sync.WaitGroup
// 		turns = make(chan bool)
// 		num   = 5
// 	)

// 	wg.Add(2)

// 	// ping
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < num; i++ {
// 			<-turns
// 			fmt.Println("ping")
// 			if i+1 < num {
// 				turns <- true
// 			}
// 		}
// 	}()

// 	// pong
// 	go func() {
// 		defer wg.Done()
// 		for j := 0; j < num; j++ {
// 			<-turns
// 			fmt.Println("pong")
// 			turns <- true
// 		}
// 	}()

// 	// start counter
// 	turns <- true
// 	wg.Wait()
// 	close(turns)
// }
