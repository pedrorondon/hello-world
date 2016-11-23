package main

import "fmt"

type person struct {
	x string
	y string
	z int
}

func (k person) fullName() string { // function is composed of func, receiver, name of function,
	// parameters in parens return, and function body in curly braces
	// this a method; the receiver (p person) attaches any value of type person to the function fullName
	// in other words, any value created of type person will have access to method fullName
	return k.x + k.y
}

func main() {
	p1 := person{"James", "Bond", 20} // James, Bond, 20 are the values for the fields inside of type person
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.fullName()) // since p1 is a variable that stores a value of type person, you can call
	// method fullName
	fmt.Println(p2.fullName())
	fmt.Printf("%T \n", p2)            // prints of type main.person
	fmt.Printf("%T \n", p2.fullName()) // prints of type string
}
