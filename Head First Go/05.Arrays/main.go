package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// open the data file for reading
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	// create scanner
	scanner := bufio.NewScanner(file)
	// loop until the end of the file is reached
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// close the file to free resources
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// report error if error was found while scanning
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
