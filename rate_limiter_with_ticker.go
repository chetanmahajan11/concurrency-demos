package main

// func main() {
// 	requests := make(chan int, 5)
// 	for i := 1; i <= 5; i++ {
// 		requests <- i
// 	}
// 	close(requests)

// 	limiter := time.NewTicker(500 * time.Millisecond)
// 	defer limiter.Stop()

// 	for req := range requests {
// 		<-limiter.C
// 		fmt.Println("Processing request", req)
// 	}
// }
