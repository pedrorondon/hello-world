package main

import "fmt"

const p string = "freedom baby" //still valid if type string is removed

func main() {
	const q = 33 //still valid code if this line 8 is moved under line 5 to the package level

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
}

//a CONSTANT is a simple unchaning value
