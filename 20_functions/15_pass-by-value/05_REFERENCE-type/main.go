package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m["Dro"], "a")
	cambio(m)
	fmt.Println(m["Dro"], "b")
}
func cambio(x map[string]int) {
	x["Dro"] = 32
}

// like slices, maps are also reference types that inherently have and pass memory addresses
