package Programmers

func coverPhoneNumberWithForLoop(phone_number string) string {

	var result string

	star := len(phone_number) - 4
	for i := 0; i < star; i++ {
		result += "*"
	}
	result += phone_number[star:]

	return result
}
