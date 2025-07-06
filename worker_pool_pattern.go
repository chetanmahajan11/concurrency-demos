package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // // worker func to work on assigned jobs
// // func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
// // 	defer wg.Done()
// // 	for job := range jobs {
// // 		fmt.Printf("\n worker %d working on job no: %d", id, job)
// // 		time.Sleep(time.Millisecond * 500)
// // 		fmt.Printf("\n worker %d working on job no: %d", id, job)
// // 	}
// // }

// // func main() {
// // 	var (
// // 		wg   sync.WaitGroup
// // 		jobs = make(chan int)
// // 	)

// // 	// start worker goroutines
// // 	for i := 0; i < 5; i++ {
// // 		wg.Add(1)
// // 		go worker(i, jobs, &wg)
// // 	}

// // 	// start assigning jobs
// // 	for j := 0; j < 5; j++ {
// // 		jobs <- j
// // 	}

// // 	// closing job channel
// // 	close(jobs)

// // 	// waiting for completion of wg counter
// // 	wg.Wait()
// // }
