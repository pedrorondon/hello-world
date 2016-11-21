package main

import "fmt"

func main() {
	fmt.Println(nombres("Juan", "Pablo"))
}

func nombres(a, b string) (s string) {
	s = fmt.Sprint(a, b)
	return
}
