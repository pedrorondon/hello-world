package main

import "fmt"

func main() {
	myMutant := map[string]string{
		"Wolverine":  "is a beast.",
		"Sabertooth": "will be deafeated.",
	}

	myMutant["Magneto"] = "dislikes humans."

	fmt.Println(len(myMutant))
}

// a map can grow and shrink as you use it
