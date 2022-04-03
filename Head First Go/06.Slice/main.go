package main

import "fmt"

func main() {
	severalInts(1, 2, 3)
	severalStrings("abc", "hello", "txt")
}

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func severalStrings(strings ...string) {
	fmt.Println(strings)
}
