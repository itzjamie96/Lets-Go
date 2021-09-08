package Programmers

func getLettersAtCenter(s string) string {

	// golang에서는 python처럼 string을 index로 가져오기위해 byte arr로 변환시킨 후 인덱스로 접근한다
	// => 사실 string을 인덱스 개념으로 접근은 가능한데 byte로 바꿔서 하는 것보다 느려서..
	given := []byte(s)
	mid := len(given) / 2 // 미리 변수에 중앙 인덱스값을 keep해둔다 (계속 계산하지 않도록)

	// 길이가 홀수인 경우에는 바로 중앙 인덱스에 있는 걸 리턴
	// 그게 아니라면 중앙-1:중앙+1 을 리턴
	if len(given)%2 == 1 {
		return string(given[mid])
	} else {
		return string(given[mid-1 : mid+1])
	}

}
