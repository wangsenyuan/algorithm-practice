package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n, &m)
	res := solve(n, m)
	fmt.Println(res)
}

func solve(n int, m int) int {
	x := max(2*n, 3*m)
	// (x + 1) / 2 + (x + 2) / 3 - (x + 5) / 6 >= n + m
	// 3 * (x + 1) + 2 * (x + 2) - (x + 5) >= 6 * (n + m))
	// 4 * x + 2 >= 6 * (n + m)
	x = max(x, (6*(n+m)-2+3)/4)

	for i := max(x-6, 1); i <= x+6; i++ {
		if i/2 >= n && i/3 >= m && (i/2+i/3-i/6) >= n+m {
			return i
		}
	}

	return x
}

func solve1(n int, m int) int {
	x := max(2*n, 3*m)

	for {
		n1 := x / 2
		m1 := x / 3
		k := x / 6
		// n1 - k 只能整除2, m1 - k 只能整除3
		// 假设k中a个被用来作为2的倍数, b个被用来作为3的倍数
		// n1 - k + a >= n
		// m1 - k + b >= m
		// n1 + m1 - k >= n + m
		if n1 >= n && m1 >= m && n1+m1-k >= n+m {
			break
		}
		x++
	}
	return x
}
