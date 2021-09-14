package Programmers

import (
	"fmt"
	"sort"
)

func kthNumber(array []int, commands [][]int) []int {

	// 결과는 몇개가 될지 모르니 우선 empty slice를 만들어둔다
	result := []int{}

	// commands가 몇개든 상관없이 우선 for문을 돌린다
	for _, c := range commands {
		start, end, k := c[0]-1, c[1], c[2]-1

		// 루프 한번 돌 때마다 새로운 arr를 만들어준 후 거기에 array의 주어진 범위만큼 넣어준다
		// slice를 쓰지 않는 이유는 slice는 가져온 원본 array를 "참조"하기 때문에
		// sort를 썼다가는 원본 배열이 sort되어버리기 때문이다...
		size := end - start
		arr := make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = array[i+start]
		}
		sort.Ints(arr)
		result = append(result, arr[k])

	}
	fmt.Println(result)

	return result
}

// 아무리봐도 효율이 똥망인 것 같다....slice를 다시 보고 풀어봐야겠다...
