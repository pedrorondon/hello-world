package main

import "fmt"

func main() {
	palabra := "Hello"
	letra := rune(palabra[0]) // from the string referred to by variable palabra, give me the letter at
	// position 0 and convert it to a rune; then, call that rune variable letra
	fmt.Println(letra) // 72 which is the decimal value of H
}
