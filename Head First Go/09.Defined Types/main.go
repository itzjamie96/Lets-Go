package main

import "fmt"

func main() {
	number := Number(4)
	fmt.Println("Original value of number:", number)

	number.Double()
	fmt.Println("Updated value of number:", number)
}

type Number int

func (n *Number) Double() {
	*n *= 2
}