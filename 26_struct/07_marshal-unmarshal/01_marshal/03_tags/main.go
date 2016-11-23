package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Primero string
	Ultimo  string `json:"-"`            // this tag removes this field
	Edad    int    `json:"wisdom score"` // this take changes the field name
}

func main() {
	p1 := persona{"James", "Bond", 20}
	byte, _ := json.Marshal(p1)
	fmt.Println(string(byte))
}
