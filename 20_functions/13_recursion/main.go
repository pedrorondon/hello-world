package main

import "fmt"

func factorial(x int) int { // factorial is the identifier that names the function; parameters in parens()
	// says that the function takes an int; the return shows that an int is returned
	if x == 0 {
		return 1 // when you hit a return statement in a function, that function returns and is done and will NOT
		// execute anymore code
	}
	return x * factorial(x-1)
}
func main() {
	fmt.Println(factorial(4)) // a factorial is taking a number and multiplying it by all the numbers below it;
	// ie 4 * 3 * 2 * 1 = 24
	// most of the time, recursion is not the most performant and running a loop would be better
}
