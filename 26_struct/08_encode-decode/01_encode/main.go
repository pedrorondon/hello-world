package main

import (
	"encoding/json"
	"os"
)

type persona struct {
	Primero     string
	Ultimo      string
	Edad        int
	notExported int
}

func main() {
	p1 := persona{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1)
}
