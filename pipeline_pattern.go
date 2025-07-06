package main

// import (
// 	"fmt"
// )


// func main() {
// 	gen := generator(5)
// 	sq := square(gen)
// 	doubled := double(sq)

// 	for num := range doubled {
// 		fmt.Printf("\nFinal output is: %d", num)
// 	}

// }

// // generator
// func generator(num int) <-chan int {
// 	ch := make(chan int)
// 	go func() {
// 		defer close(ch)
// 		for i := 0; i < num; i++ {
// 			ch <- i
// 			fmt.Printf("\n pushed value on channel is: %d", i)
// 		}
// 	}()
// 	return ch
// }

// // square
// func square(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for num := range in {
// 			out <- num * num
// 			// fmt.Printf("\n square of %d is : %d", num, num*num)
// 		}
// 	}()
// 	return out
// }

// // double
// func double(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for num := range in {
// 			out <- num * 2
// 			// fmt.Printf("\n Doubled value of %d is %d", num, num*2)
// 		}
// 	}()
// 	return out
// }
