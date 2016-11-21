package main

import "fmt"

func main() {
	age := 32
	fmt.Println(age, "a")
	fmt.Println(&age, "b") // this is an address that is storing an int, 32

	fountainOfYouth(&age)

	fmt.Println(&age, "c")
	fmt.Println(age, "d")
}

func fountainOfYouth(y *int) { // *int is of type pointer to an address that is storing an int, 32;
	// the parameter() then assigns that value to variable y; asterix* is part of the type;
	// the address is being "passed by value" into this function
	fmt.Println(y, "e")
	fmt.Println(*y, "f") // asterix* here is an operator to dereference an address causing it to show/reveal
	// the int that is stored there
	*y = 18 // asterix* is saying give me the value stored at this address, y, and assign a new value, 18
	fmt.Println(y, "g")
	fmt.Println(*y, "h")
}

// age and y are two different variable that point to the same address
