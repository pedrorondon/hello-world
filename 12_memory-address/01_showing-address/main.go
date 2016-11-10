package main

import "fmt"

func main() {
	a := 43 // the variable a is being set equal to 43

	fmt.Println("a -", a)                                   // this function prints the value of a
	fmt.Println("a's memory address in hexadecimal - ", &a) // the memory address is displayed using the ampersand sign;
	// & is known as an Operator
	// after running the code and printing the memory address, 0x is seen which indicates hexadecimal
	fmt.Printf("a's memory address in decimal - %d\n", &a)
}
