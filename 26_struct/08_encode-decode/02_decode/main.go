package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type persona struct {
	Primero     string
	Ultimo      string
	Edad        int
	notExported int
}

func main() {
	var p1 persona // create variable p1 of type person and initialize the variable to its zero value
	rdr := strings.NewReader(`{"Primero": "James", "Ultimo": "Bond", "Edad": 20}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1.Primero)
	fmt.Println(p1.Ultimo)
	fmt.Println(p1.Edad)
}
