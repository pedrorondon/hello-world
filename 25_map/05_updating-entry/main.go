package main

import "fmt"

func main() {

	myMutant := map[string]string{
		"Wolverine":  "is a beast.",
		"Sabertooth": "will be defeated.",
	}
	myMutant["Magneto"] = "dislikes humans."
	fmt.Println(myMutant)
	myMutant["Magneto"] = "loves Charles though."
	fmt.Println(myMutant)
}

// underlying a slice is an array
// underlying a map is a hash table
