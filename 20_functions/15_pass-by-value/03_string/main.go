package main

import "fmt"

func main() {
	nombre := "Drizzy"
	fmt.Println(nombre, "a")

	nameChange(nombre)

	fmt.Println(nombre, "b")
}

func nameChange(x string) {
	fmt.Println(x, "c")
	x = "Rocky" // ths only changes tha=e value of variable x in its limited scope but does not permanently
	// change the original string, nombre
	// strings are immutable but a new string can always be assigned to a variable
	fmt.Println(x, "d")
}
