package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var n, k int
		fmt.Fscan(reader, &n, &k)
		res := solve(n, k)
		if len(res) == 0 {
			fmt.Fprintln(writer, 0)
		} else {
			s := fmt.Sprintf("%v", res)
			fmt.Fprintln(writer, s[1:len(s)-1])
		}
	}
}

const N = 31
const X = N * (N - 1) / 2

var dp [N][X + 1]bool

func init() {
	dp[0][0] = true

	for i := 1; i < N; i++ {
		for i1 := range i {
			for x, v := range dp[i1] {
				if v && x+(i1+1)*(i-i1) <= X {
					dp[i][x+(i1+1)*(i-i1)] = true
				}
			}
		}
		dp[i][0] = true
	}

}

func solve(n int, k int) []int {
	res := make([]int, n)
	for i := range n {
		res[i] = i + 1
	}
	if k == 0 {
		return res
	}
	// 考虑n的位置，pos[n], 那么它的贡献 = (pos[n] + 1) * (n - pos[n] - 1) + left(pos[n]) + right(n - pos[n] - 1)
	// 假设存在位置i1, i2, i3... P[i1] > P[i1+1], .. P[i3] > P[i3+1]
	// 那么贡献  = (i1 + 1) * (i2 - i1) + (i2 + 1) * (i3 - i2), ...
	// 比如 5， 4， 3， 2， 1
	// 1 * 1 + 2 * 1 + 3 * 1 + 4 * 1 = 10
	// 感觉在正确的方向上
	// 那么就是选择i1, i2, ... iw 满足 (i1 + 1) * (i2 - i1) + (i2 + 1) * (i3 - i2) + .. (iw + 1) * (n - iw - 1) = k

	var special []int
	m := n - 1
	for r := n - 2; r >= 0; r-- {
		for x := k; x >= 0; x-- {
			if dp[r][x] && (m-r)*(r+1)+x == k {
				// res[r] > buf[r+1], ...
				// res[r], res[r+1] = res[r+1], res[r]
				special = append(special, r)
				k = x
				m = r
				break
			}
		}
	}

	if k != 0 {
		return nil
	}

	slices.Reverse(special)
	hi := n
	lo := 1
	for i, j := 0, 0; i < n; i++ {
		if j < len(special) && special[j] == i {
			res[i] = hi
			hi--
			j++
		} else {
			res[i] = lo
			lo++
		}
	}

	return res
}
