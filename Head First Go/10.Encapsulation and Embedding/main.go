package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 13 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
func main() {
	//date := Date{year: 2022, month: 4, day: 4}
	date := Date{}
	err := date.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.year)

	err = date.SetMonth(4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.month)

	err = date.SetDay(33)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.day)
}
