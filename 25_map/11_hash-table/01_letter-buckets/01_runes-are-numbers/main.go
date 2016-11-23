package main

import "fmt"

func main() {
	lett := 'A' //'A' is a rune; letters are wrapped in single quotes'; strings are wrapped double quotes"
	// or back ticks`
	fmt.Println(lett)         // 65; prints the letter's decimal value
	fmt.Printf("%T \n", lett) // of type int32 because it is UTF8 which is a 1 to 4 byte coding scheme
}
