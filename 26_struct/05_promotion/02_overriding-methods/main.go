package main

import "fmt"

type person struct {
	nombre string
	edad   int
}

type DoubleZero struct { // outer type
	person        // inner type;  inner type promotion means that it is available to th outer type
	LicenseToKill bool
}

func (k person) Sayings() { // method of the inner type person
	fmt.Println("make one percent off of 100.")
}

func (q DoubleZero) Sayings() { // method of the outer type DoubleZero
	fmt.Println("rich people constantly learn and grow.")
}

func main() {
	p1 := person{
		nombre: "J. Paul Getty",
		edad:   40,
	}
	p2 := DoubleZero{
		person: person{
			nombre: "MLK",
			edad:   50,
		},
		LicenseToKill: true,
	}
	p1.Sayings()
	p2.Sayings() // will print the DoubleZero Sayings since the outer type overrides the inner type
	p2.person.Sayings()
}
