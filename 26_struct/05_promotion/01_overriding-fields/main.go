package main

import "fmt"

type person struct {
	x string
	y string
	z int
}

type DoubleZero struct { // DoubleZero is the outer type embedding the inner type
	person        // person is the inner type embedded in the outer type
	x             string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{ // give value to the three fields: person, x, and LicenseToKill
		person: person{
			x: "Professor",
			y: "Xavier",
			z: 65,
		},
		x:             "is a great leader.",
		LicenseToKill: true,
	}
	p2 := DoubleZero{
		person: person{
			x: "Jean",
			y: "Grey",
			z: 21,
		},
		x:             "moves things with her mind.",
		LicenseToKill: false,
	}
	fmt.Println(p1.x, p1.person.x) // outer type x overrides inner type x found within person; you can
	// access the inner type x through .dot notation
	fmt.Println(p2.x, p2.person.x)
}
