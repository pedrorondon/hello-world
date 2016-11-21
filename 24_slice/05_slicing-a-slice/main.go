package main

import "fmt"

func main() {
	lenguas := []string{
		"Good morning",
		"Buenos Dias",
		"Bonjour",
		"Bongiorno",
		"Ohayo",
		"Gutten morgen",
	}
	fmt.Print("[1:2]")
	fmt.Println(lenguas[1:2])
	fmt.Print("[:2]")
	fmt.Println(lenguas[:2])
	fmt.Print("[4:]")
	fmt.Println(lenguas[4:])
	fmt.Print("[:]")
	fmt.Println(lenguas[:])

}
