package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	numbers, err := GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(numbers)
}

// GetFloats reads a float64 from each line of a file.
// GetFloats will return an array of numbers and any error encountered.
func GetFloats(filename string) ([3]float64, error) {

	// returning array
	var numbers [3]float64

	// open the data file for reading
	file, err := os.Open(filename)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
