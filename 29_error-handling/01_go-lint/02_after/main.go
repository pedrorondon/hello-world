package main

import "fmt"

func main() {
	x := 33
	str := eval(x)
	fmt.Println(str)
}

func eval(n int) string {
	if n > 10 {
		return fmt.Sprint("x is greater than 10")
	}
	return fmt.Sprint("x is not greater than 10")
}

// the else statement has been removed
