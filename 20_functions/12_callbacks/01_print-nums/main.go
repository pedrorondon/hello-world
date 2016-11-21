package main

import "fmt"

func visit(numbers []int, callback func(int)) { // first parameter is a slice of int and the second parameter
	// is a function; numbers and callback are the variables followed by their types, []int and func(int) respectively
	for _, n := range numbers {
		callback(n)

	}

}
func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

// callback is passing a function as an argument
