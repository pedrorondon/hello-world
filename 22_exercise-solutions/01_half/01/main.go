package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0

}
func main() {
	a, b := half(5) // when calling function half(1), it must be assigned to two variables, a and b
	// since the function has two returns
	fmt.Println(a, b)
}
