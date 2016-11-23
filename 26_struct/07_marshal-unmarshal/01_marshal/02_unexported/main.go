package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	primero string // notice the lower case for all of the field names causing them to be unexported
	ultimo  string
	edad    int
}

func main() {
	p1 := persona{"James", "Bond", 20}
	fmt.Println(p1)
	byte, _ := json.Marshal(p1)
	fmt.Println(byte)         // only prints byte equivalent of curly braces{}
	fmt.Println(string(byte)) // only prints {} with nothing else because the fields for persona are lower
	// case and therefore, not exported

}
