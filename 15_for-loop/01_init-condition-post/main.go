package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ { // code ++ means to add 1; the code will jump out of the loop once the statement
		// becomes false; in this case, the statement becomes false once i is greater than 100
		fmt.Println(i)
	}
}
