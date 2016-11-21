package main

import "fmt"

func main() {
	for i := 1; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i))) // alternate method of wriiting this code is
		// found in next folder down 02
	}
	foo := 'a'                     // single quotation marks, ' ', are used to indicate a rune
	fmt.Println(foo)               // code will print out decimal equivalent of rune a which is 97
	fmt.Printf("%T \n", foo)       // %T will show the type which is rune or int32 in this case
	fmt.Printf("%T \n", rune(foo)) // code will print the conversion of foo to rune which is rune to rune

}
