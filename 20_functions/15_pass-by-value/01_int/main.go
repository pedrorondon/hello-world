package main

import "fmt"

func main() {
	age := 32
	fountain(age)
	fmt.Println(age, "a")
}
func fountain(n int) {
	fmt.Println(n, "b")
	n = 18
}
