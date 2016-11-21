package main

import "fmt"

func main() {

	myMutant := map[string]string{
		"Wolverine":  "is a beast.",
		"Sabertooth": "will be defeated.", // note that this ending comma is required
	}
	myMutant["Magneto"] = "dislikes humans." // at this key, "Magneto", give it this value, "dislikes humans."

	fmt.Println(myMutant)
}
