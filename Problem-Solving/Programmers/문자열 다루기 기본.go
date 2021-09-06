package Programmers

func dealingWithStringsBasic(s string) bool {

	// 길이가 4 혹은 6이 아닌 경우는 false
	if len(s) != 4 && len(s) != 6 {
		return false
	}

	// 문자열 s를 돌면서 ascii 값 0-9 사이에 없으면 false
	for _, v := range s {
		if v < 48 || v > 57 {
			return false
		}
	}

	return true
}
