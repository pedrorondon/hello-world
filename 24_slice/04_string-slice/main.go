package main

import "fmt"

func main() {
	lenguas := []string{
		"Good morning",
		"Buenos Dias",
		"Bonjour", // note the ending comma
	}
	for i, entry := range lenguas {
		fmt.Println(i, entry)

	}
	{
		fmt.Println("Length: ", len(lenguas))
	}
	{
		fmt.Println("--------------------")
	}

	for q := 0; q < len(lenguas); q++ { // loop while q is less than the length of the slice, lenguas[]
		fmt.Println(lenguas[q])

	}
	fmt.Println("Best language: ", lenguas[1])
}
