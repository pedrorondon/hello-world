package main

import "fmt"

func main() {
	alphabet := make([]string, 3, 5)

	fmt.Println("Length before:", len(alphabet))
	fmt.Println("Capacity before: ", cap(alphabet))

	alphabet[0] = "A"
	alphabet[1] = "B"
	alphabet[2] = "C"
	alphabet = append(alphabet, "D")
	alphabet = append(alphabet, "E")
	alphabet = append(alphabet, "F")
	alphabet = append(alphabet, "G")

	fmt.Println(alphabet[6])
	fmt.Println("Length after:", len(alphabet))
	fmt.Println("Capacity after: ", cap(alphabet))
}
