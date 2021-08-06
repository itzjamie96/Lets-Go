package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	myBill := createBill()
	options(myBill)

}

func getInput(ask string, reader *bufio.Reader) (string, error) {
	fmt.Print(ask)

	// the reader will return multiple values, but we'll only use
	// the first part = string
	input, err := reader.ReadString('\n')
	// trim any possible spaces
	input = strings.TrimSpace(input)

	return input, err

}

func options(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	// switch case
	switch opt {
	case "a":
		menu, _ := getInput("Name of menu: ", reader)
		price, _ := getInput("Price of menu: ", reader)

		p, err := strconv.ParseInt(price, 10, 32)

		// if there is an error
		if err != nil {
			fmt.Println("The price must be number")
			options(b)
		}
		intPrice := int(p)
		b.addMenu(menu, intPrice)

		fmt.Printf("Added menu %v - %v\n", menu, intPrice)
		options(b)
	case "s":
		b.saveBill()
		fmt.Println("Saved bill")
	case "t":
		tip, _ := getInput("Tip: ", reader)

		t, err := strconv.ParseInt(tip, 10, 32)

		// if there is an error
		if err != nil {
			fmt.Println("The price must be number")
			options(b)
		}
		intTip := int(t)
		b.updateTip(intTip)
		fmt.Printf("Updated tip to %v\n", intTip)
		options(b)
	default:
		fmt.Println("Not a valid option. Please try again")
		options(b)
	}

}
