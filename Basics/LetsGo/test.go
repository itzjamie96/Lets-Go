package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func switchCase() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter score: ")
	score, _ := reader.ReadString('\n')
	score = strings.TrimSpace(score)

	switch score {
	case "A":
		fmt.Println(100)
	case "B":
		fmt.Println(90)
	case "C":
		fmt.Println(80)
	default:
		fmt.Println("Please enter a valid score")
		switchCase()
	}
}

func parseString(s string) {
	intN, _ := strconv.ParseInt(s, 10, 32)
	floatN, _ := strconv.ParseFloat(s, 64)
	uintN, _ := strconv.ParseUint(s, 10, 64)

	fmt.Printf("%d is type %v\n", intN, reflect.TypeOf(intN))
	fmt.Printf("%f is type %v\n", floatN, reflect.TypeOf(floatN))
	fmt.Printf("%v is type %v\n", uintN, reflect.TypeOf(uintN))
}
