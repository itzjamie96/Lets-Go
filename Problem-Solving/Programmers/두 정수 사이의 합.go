package Programmers

func sumOfTwoIntegers(a int, b int) int64 {

	// a가 b보다 큰 경우에만 두 값들은 swap
	if a > b {
		a, b = b, a
	}

	result := 0
	for i := a; i <= b; i++ {
		result += i
	}

	return int64(result)
}
