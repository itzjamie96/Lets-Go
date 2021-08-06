package main

import (
	"bufio"
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]int
	tip   int
}

// create a new object as a bill struct
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill named: ", reader)

	resultBill := newBill(name)
	return resultBill
}

// construct an initial bill type
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]int{},
		tip:   0,
	}
	return b

}

// add new elements to the map in bill
func (b bill) addMenu(menu string, price int) {
	menuList := b.items
	menuList[menu] = price
}

// return all elements of the bill in formatted statements
func (b bill) format() string {
	list := "<<Your Bill>>\n"
	var total int = 0

	// list items
	for k, v := range b.items {
		list += fmt.Sprintf("%-25v ... W %v \n", k+":", v)
		total += v
	}

	// total
	list += fmt.Sprintf("%-25v ... W %v", "total:", total)

	return list
}

// update element of the bill
func (b *bill) updateTip(tip int) {
	b.tip = tip
}

// save the bill into a file
func (b *bill) saveBill() {

	// format the bill and make it into a slice of bytes
	data := []byte(b.format())

	// write into a file
	// save location, data to write, permission
	// this will return an error
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	// if there is an error,
	if err != nil {
		// stop the flow of the program and print the error
		panic(err)
	}

	fmt.Println("Bill was saved into .txt")
}
