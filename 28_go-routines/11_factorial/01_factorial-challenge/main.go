package main

import "fmt"

func main() {
	f := factorial(4)
	fmt.Println("Total:", f)
}

func factorial(n int) int {
	total := 1
	for i := n; n > 0; i-- {
		total *= i
	}
	return total
}

// 4! notation indicated that you want the factorial of 4
// 4*3*2*1 = 24
