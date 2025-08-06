package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	var n, t int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &t)
	res := solve(n, t)
	fmt.Println(res)
}

func solve(n int, t int) int {
	if t&(t-1) != 0 {
		return 0
	}
	if n == 1 {
		if t == 2 {
			return 0
		}
		return 1
	}
	h := bits.Len(uint(t))
	// 2 ** h = t
	// 2 ** (h - 1) = t
	n++
	// h个1

	var digits []int
	for i := n; i > 0; i >>= 1 {
		digits = append(digits, i&1)
	}
	slices.Reverse(digits)
	m := len(digits)

	if h > m {
		return 0
	}

	C := make([][]int, m+1)
	for i := range m + 1 {
		C[i] = make([]int, i+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j] + C[i-1][j-1]
		}
	}

	var res int
	var cnt int
	for i := 0; i < m && h > cnt; i++ {
		if digits[i] == 1 {
			// 如果这里放置0，那么后面就可以随便放置1
			if h-cnt <= m-i-1 {
				res += C[m-i-1][h-cnt]
			}
			cnt++
		}
	}
	if cnt == h && h > 1 {
		res++
	}
	return res
}
