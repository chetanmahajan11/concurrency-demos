package main

// import (
// 	"fmt"
// 	"sync"
// )

/*
	Key summary:
	- In the below example, first signal is given to chEven. So the printing is started from chEven and then the signal is pushed to chOdd.
	- Since last digit has to be printed from the chEven(which is extra one as compared to the numbers which has been printed so far in both categories),
	   the if-condition has been put up in the even goroutine. Otherwise it should be placed in odd goroutine.
*/

// func main() {
// 	chOdd := make(chan bool)
// 	chEven := make(chan bool)
// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	// goroutine for even number
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i <= 20; i += 2 {
// 			<-chEven
// 			fmt.Println(i)
// 			if i+1 < 21 {
// 				chOdd <- true
// 			}
// 		}
// 	}()

// 	// goroutine for odd number
// 	go func() {
// 		defer wg.Done()
// 		for i := 1; i < 20; i += 2 {
// 			<-chOdd
// 			fmt.Println(i)
// 			chEven <- true
// 		}
// 	}()

// 	// start the groutine from even numbers
// 	chEven <- true
// 	wg.Wait()
// }
