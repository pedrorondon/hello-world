package main

import "fmt"

const (
	A = iota // 0 - code will return numbers 0 through 3 when executed
	B = iota //still valid is the word iota is removed from B through D
	C = iota
	D = iota
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
}
