package main

import "fmt"

func main() {

	myMutant := map[string]string{ // this method is known as composite literal since it is composed of literal values
		"Wolverine":  "is a beast.", // note the positioning of the : and ,
		"Sabertooth": "will be defeated.",
	}

	fmt.Println(myMutant)
}
