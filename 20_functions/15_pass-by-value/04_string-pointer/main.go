package main

import "fmt"

func main() {
	nombre := "Drizzy"
	fmt.Println(&nombre, "a")

	nameChange(&nombre)

	fmt.Println(&nombre, "b")
	fmt.Println(nombre, "c")
}

func nameChange(x *string) {
	fmt.Println(x, "d")
	fmt.Println(*x, "e")
	*x = "Rocky"
	fmt.Println(x, "f")
	fmt.Println(*x, "g")
}
