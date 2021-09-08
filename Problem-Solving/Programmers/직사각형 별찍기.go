package Programmers

import "fmt"

func rectangularStars() {
	var n, m int
	fmt.Scan(&n, &m)
	// fmt.Println(a + b)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(string('*'))
		}
		fmt.Println()
	}

}

// "" < 이걸 출력하는 것보다 string('') 이걸 출력하는게 훨씬 빠름!
