package main

import "fmt"

func main() {
	records := make([][]string, 0)
	// student 1
	stud1 := make([]string, 4)
	stud1[0] = "Foster"
	stud1[1] = "Nathan"
	stud1[2] = "100.00"
	stud1[3] = "74.00"
	// store the record
	records = append(records, stud1)
	// student 2
	stud2 := make([]string, 4)
	stud2[0] = "Gomez"
	stud2[1] = "Lisa"
	stud2[2] = "92.00"
	stud2[3] = "96.00"
	// store the record
	records = append(records, stud2)
	//print
	fmt.Println(records)
}
