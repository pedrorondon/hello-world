package main

import "fmt"

func main() {
	greet := make([]string, 3, 5)
	fmt.Println("Length before: ", len(greet))
	greet[0] = "Hello"
	greet[1] = "Hola"
	greet[2] = "Bonjour"
	greet = append(greet, "Sup")

	fmt.Println(greet[3])
	fmt.Println("Length after: ", len(greet))
}
