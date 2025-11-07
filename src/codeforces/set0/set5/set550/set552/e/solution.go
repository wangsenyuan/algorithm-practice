package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	fmt.Println(res)
}

type pair struct {
	first  int
	second int
}

func solve(s string) int {
	// a[i]的位置 = 2 * i, 2 * i - 1是它前面的符号， 2 * i + 1是它后面的符号
	n := (len(s) + 1) / 2
	dp := make([]int, n)
	dp2 := make([]pair, n)

	for i := 0; i < n; i++ {
		val := 1
		j := i
		// 这里不会超过15次
		for ; j >= 0; j-- {
			y := int(s[2*j] - '0')
			val *= y
			if j > 0 && s[2*j-1] == '+' {
				break
			}
		}
		dp[i] = val
		dp2[i] = pair{val, j - 1}
		if j > 0 {
			dp[i] += dp[j-1]
		}
	}
	fp := make([]int, n)
	fp2 := make([]pair, n)

	for i := n - 1; i >= 0; i-- {
		val := 1
		j := i
		for ; j < n; j++ {
			y := int(s[2*j] - '0')
			val *= y
			if j < n-1 && s[2*j+1] == '+' {
				break
			}
		}

		fp[i] = val
		fp2[i] = pair{val, j + 1}
		if j+1 < n {
			fp[i] += fp[j+1]
		}
	}

	res := dp[n-1]

	play := func(l int, x int, r int) int {
		v1 := int(s[2*l] - '0')
		v2 := int(s[2*r] - '0')
		x *= dp2[l].first / v1
		x *= fp2[r].first / v2
		if dp2[l].second >= 0 {
			x += dp[dp2[l].second]
		}
		if fp2[r].second < n {
			x += fp[fp2[r].second]
		}
		return x
	}

	for i := range n {
		// 在i的前面添加
		buf := []int{int(s[2*i] - '0')}
		for j := i + 1; j < n; j++ {
			// 在j的后面添加
			x := int(s[2*j] - '0')
			if s[2*j-1] == '*' {
				// 如果是乘法，就喝前面的直接乘起来
				buf[len(buf)-1] *= x
				// len(buf) <= 1
			} else {
				buf = append(buf, x)
			}
			for len(buf) > 2 {
				buf[1] += buf[0]
				buf = buf[1:]
			}
			// len(buf) <= 2
			x = buf[0]
			if len(buf) == 2 {
				x += buf[1]
			}
			res = max(res, play(i, x, j))
		}
	}

	return res
}

func precedent(a byte, b byte) bool {
	return a == '*' || a == b
}
