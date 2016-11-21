package main

import "fmt"

func main() {
	greet := make([]string, 3, 5)

	greet[0] = "Hello"
	greet[1] = "Hola"
	greet[2] = "Bonjour"
	greet[3] = "Sup" // this is outside of the range

	fmt.Println(greet[2])
}
