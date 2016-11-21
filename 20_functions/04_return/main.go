package main

import "fmt"

func main() {
	fmt.Println(nombres("Juan", "Pablo"))
}

func nombres(a, b string) string {
	return fmt.Sprint(a, b)
}
