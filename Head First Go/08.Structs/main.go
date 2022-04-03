package main

import "fmt"

type car struct {
	name     string
	topSpeed float64
	detail   // anonymous field
}

type detail struct {
	color  string
	height float64
	width  float64
}

func main() {
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)

	amount := 6
	double(&amount)
	fmt.Println(amount)
}

func double(number *int) {
	*number *= 2
}
