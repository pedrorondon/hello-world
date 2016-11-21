package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	s := []int{1, 2, 3, 4}
	foo(s...)
	foo()
}

func foo(nums ...int) { // foo is a variadic function which takes from 0 to an unlimited number of ints
	fmt.Println(nums)

}
