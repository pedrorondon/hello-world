package main

import "fmt"

type foo int

func main() {
	var edad foo
	edad = 33
	fmt.Printf("%T %v \n", edad, edad)

	var tuedad int
	tuedad = 18
	fmt.Printf("%T %v\n", tuedad, tuedad)

}

// this does NOT work:
// fmt.Println(edad + tuedad) this is a mismatch of types; type foo cannot be added to type int

// this does work:
// fmt.Println(int(edad) + tuedad) edad of type foo is being converted to type int; this can be done since the
// underlying type for foo is type int
