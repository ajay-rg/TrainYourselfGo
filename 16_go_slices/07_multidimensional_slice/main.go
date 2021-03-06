package main

import "fmt"

func main() {
	// String
	var records [][]string
	// student 1
	student1 := make([]string, 4)
	student1[0] = "Kylo"
	student1[1] = "Ren"
	student1[2] = "100.00"
	student1[3] = "74.00"
	// store the record
	records = append(records, student1)
	// student 2
	student2 := make([]string, 4)
	student2[0] = "Rey"
	student2[1] = "Skywalker"
	student2[2] = "92.00"
	student2[3] = "96.00"
	// store the record
	records = append(records, student2)
	// print
	fmt.Println(records)

	// Int
	transactions := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		transaction := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, i+j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
}
