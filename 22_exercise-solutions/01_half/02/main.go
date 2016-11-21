package main

import "fmt"

func half(x int) (float64, bool) {
	return float64(x) / 2, x%2 == 0 //  convert int, x to float64 by surrounding it with parens()
}
func main() {
	a, b := half(1)
	fmt.Println(a, b)
}
