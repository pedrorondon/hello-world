package main

import "fmt"

func main() {
	fmt.Println(nombres("Juan ", "loves ", "Pablo ", "but ", "does not love "))
}
func nombres(a, b, c, d, e string) (string, string) {
	return fmt.Sprint(a, b, c), fmt.Sprint(d, c, e, a)
}
