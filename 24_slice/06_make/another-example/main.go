package main

import "fmt"

func main() {
	nums := make([]int, 3)
	// when you only specify one number, it gives both the length and the capacity;
	// 3 is both length and capacity
	// // length - number os elements referred to by the slice
	// // capacity - number of elements in the underlying array
	nums[0] = 9 // setting/assigning elements
	nums[1] = 11
	nums[2] = 42

	fmt.Println(nums[0]) // printing elements
	fmt.Println(nums[1])
	fmt.Println(nums[2])

	alt := make([]string, 3, 5) // creating a number slice and giving a higher array capacity so the
	// slice has room to grow
	// 3 is length of slice
	// 5 is capacity of underlying array
	alt[0] = "Hello"
	alt[1] = "Hola"
	alt[2] = "Bonjour"

	fmt.Println(alt[0])
	fmt.Println(alt[1])
	fmt.Println(alt[2])
}
