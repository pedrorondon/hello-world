package main

import "fmt"

func main() {
	fmt.Println(string([]byte{'a', 'b'}))
}

/*
my code:

    func main() {
	var x = 'a' // both x and y are runes as indicated by the single quotations''
	var y = 'b'
	fmt.Println(x)
	fmt.Println(y)
	q := []byte{'a'} converts rune to slice of bytes; must use curly braces{}
	r := []byte{'b'}
	fmt.Println(string(q)) converts slice of bytes to string
	fmt.Println(string(r))
}
*/
