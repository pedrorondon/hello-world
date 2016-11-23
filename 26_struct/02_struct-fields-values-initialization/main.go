package main

import "fmt"

type person struct { // in Go, you can create your own type that has an underlying type like struct in this case;
	// a struct is an aggregate type which is a collection of other things; this is encapsulation;
	// many fields can be encapsulated into one data structure; a struct is a sequence of named elements called fields
	x string // x, y, and z are fields that have been encapsulated; in the video, they are written as
	// first, second, and age; each field has a name and a type
	y string
	z int
}

func main() {
	p1 := person{"James", "Bond", 20}      // create variables of the type you created
	p2 := person{"Miss", "Moneypenny", 18} // value of type person is being assigned to the variable p2
	// {} contains a struct literal

	fmt.Println(p1.x, p1.y, p1.z)
	fmt.Println(p2.x, p2.y, p2.z)
}

/*
   a struct is a user defined type

   We used shorhand notation/shorthand variable declarations, :=, to create a variable named p1 of type person
   and variable p2 of type person
   We initialized those variable with specific values
   We used the short varialbe declaration operator with a struct literal to initialize

   line 5 declares the type and line 8, 10, and 11 are fields
   the type we created, person, has an underlying type of struct
   We delcare variables, p1 and p2, of the type and initialize those variables with a specific value or to
   to the zero value

   zero value:
   for numeric types, the zero value would be 0; for strings, it would be emptye; for booleans, it would be false

   a struct us a composite/aggregate type

   when initializing to zero, it is idiomatic to use var; otherwise use the short variable declaration with
   struct literal: := type{}
*/
