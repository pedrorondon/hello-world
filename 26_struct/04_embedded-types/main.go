package main

import "fmt"

type person struct {
	x string
	y string
	z int
}
type DoubleZero struct {
	person
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{ // struct literal in between {} creates values for the two fields within type DoubleZero;
		// using a colon :, you are specifically naming which fields get which values; this us good practice when
		// writing out embedded types since it is more clear and readable
		person: person{
			x: "James",
			y: "Bond",
			z: 20,
		},
		LicenseToKill: true, //  note the trailing comma
	}
	p2 := DoubleZero{
		person: person{"Miss", "Moneypenny", 18}, // all of these inner types of person are promoted to
		// the outer type of DoubleZero; you have access to them with .dot notation as seen in print lines 31 and 32
		LicenseToKill: false,
	}
	fmt.Println(p1.x, p1.y, p1.z, p1.LicenseToKill)
	fmt.Println(p2.x, p2.y, p2.z, p2.LicenseToKill)
}
