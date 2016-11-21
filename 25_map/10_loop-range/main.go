package main

import "fmt"

func main() {
	myGreeting := map[int]string{ // using composite literal to then input literal values
		0: "Good morning",
		1: "Bonjour",
		2: "Buenos dias",
		3: "Bonjourno", // note the trailing comma
	}

	for x, y := range myGreeting { // use the range loop to loop over our map
		//x position gives the key and y position gives the value/element
		fmt.Println(x, " - ", y)
	}
}
