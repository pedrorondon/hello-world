package main

import "fmt"

func main() {
	transactions := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		trans := make([]int, 0)
		for j := 0; j < 4; j++ {
			trans = append(trans, j)
		}
		transactions = append(transactions, trans)
	}
	fmt.Println(transactions)
}
