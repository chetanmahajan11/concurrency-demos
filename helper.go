package main



/*
func main(){

	digits := []int{2, 3, 4}
	fmt.Println("before: ", digits)

	plusOne(digits)

	fmt.Println("after: ", digits)

	}
*/

func plusOne(digits []int) []int {
	n := len(digits)

	x := digits[n-1] + 1
	if x <= 9 {
		digits[n-1]++
	}

	return digits
}
