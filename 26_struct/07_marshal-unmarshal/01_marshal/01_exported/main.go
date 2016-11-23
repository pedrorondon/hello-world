package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Primero     string // the fields that are capitalized are exported
	Ultimo      string
	Edad        int
	notExported int // since lower case, not exported
}

func main() {
	p1 := person{"Scott", "Summers", 25, 101} // value of type person is assigned to variable p1
	byteSlice, _ := json.Marshal(p1)          // turns struct p1 into a slice of bytes and assigns it to
	// variable byteSlice
	fmt.Println(byteSlice)         // prints your JSON
	fmt.Printf("%T \n", byteSlice) // prints []uint8 means of type slice of bytes since one byte has 8 bits
	fmt.Println(string(byteSlice)) // prints your JSON/slice of bytes after converting it to a string
}
