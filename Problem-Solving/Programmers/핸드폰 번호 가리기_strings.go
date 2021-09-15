package Programmers

import "strings"

func coverPhoneNumberWithStrings(phone_number string) string {

	star := len(phone_number) - 4

	result := strings.Repeat("*", star)
	result += phone_number[star:]

	return result
}
