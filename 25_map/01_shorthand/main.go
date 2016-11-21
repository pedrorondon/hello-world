package main

import "fmt"

func main() {

	myMutant := make(map[string]string) // to make shorthand, take away the var and add a : to the =
	myMutant["Wolverine"] = "is a beast."
	myMutant["Sabertooth"] = "will be defeated."

	fmt.Println(myMutant)
}
