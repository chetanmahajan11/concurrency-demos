package main

// func main() {

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

// 	defer cancel()

// 	select{
// 	case <-time.After(time.Second*5):
// 		fmt.Println("Done")

// 	case <-ctx.Done():
// 		fmt.Println("ctx cancelled \nerror details:", ctx.Err())
// 	}

// }
