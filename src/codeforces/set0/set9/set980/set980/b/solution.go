package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	res := solve(n, k)
	if res == nil {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	for _, x := range res {
		fmt.Println(x)
	}
}

func solve(n int, k int) []string {
	buf := make([][]byte, 4)
	for i := range 4 {
		buf[i] = make([]byte, n)
		for j := range n {
			buf[i][j] = '.'
		}
	}
	if k%2 == 0 || k >= 5 {
		simple(n, k, buf)
	} else if k == 1 {
		buf[1][n/2] = '#'
	} else {
		// k == 3
		buf[1][n/2] = '#'
		buf[1][n/2-1] = '#'
		buf[1][n/2+1] = '#'
	}

	res := make([]string, 4)
	for i := range 4 {
		res[i] = string(buf[i])
	}
	return res
}

func simple(n int, k int, buf [][]byte) {

	s := 1
	if k%2 == 1 {
		buf[1][1] = '#'
		buf[2][1] = '#'
		buf[1][2] = '#'
		k -= 3
		s = 3
	}
	for i := s; k > 0; i++ {
		buf[1][i] = '#'
		buf[2][i] = '#'
		k -= 2
	}

}
