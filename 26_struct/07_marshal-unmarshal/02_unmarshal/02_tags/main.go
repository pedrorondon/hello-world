package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Primero string
	Ultimo  string
	Edad    int `json:"wisdom score"`
}

func main() {
	var p1 persona //  create a value of type persona and assign it to variable p1; so p1 is of type persona;
	// since var was used, it will call get initialized to the zero value giving empty strings and 0 for int
	fmt.Println(p1.Primero) // prints nothing
	fmt.Println(p1.Ultimo)  // prints nothing
	fmt.Println(p1.Edad)    // prints 0

	byte := []byte(`{"Primero":"James", "Ultimo":"Bond", "wisdom score":20}`) // create a slice of bytes from
	// your JSON; used back ticks` since double quotes" are used as well
	json.Unmarshal(byte, &p1) // unmarshal the slice of bytes and pass in the address, &, to the struct, p1

	fmt.Println("--------------------------")
	fmt.Println(p1.Primero)
	fmt.Println(p1.Ultimo)
	fmt.Println(p1.Edad)
	fmt.Printf("%T \n", p1)
}
