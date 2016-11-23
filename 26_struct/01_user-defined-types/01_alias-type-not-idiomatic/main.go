package main

import "fmt"

type foo int

func main() {
	var edad foo
	edad = 32
	fmt.Printf("%T %v \n", edad, edad)
}
