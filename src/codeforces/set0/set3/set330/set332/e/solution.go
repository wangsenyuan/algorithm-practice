package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

// readString reads a line; content is ASCII 32–126, line terminator is \n or \r\n.
func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimRight(s, "\r\n")
}

func drive(reader *bufio.Reader) string {
	p := readString(reader)
	s := readString(reader)
	q := readString(reader)
	k, _ := strconv.Atoi(q)
	return solve(p, s, k)
}

func solve(p string, s string, k int) string {
	// len(s) <= 200, k <= 1000, len(p) <= 1e6
	n := len(s)
	m := len(p)
	dp := make([][]int, k)
	fp := make([][]int, k)
	for i := range k {
		dp[i] = make([]int, n)
		fp[i] = make([]int, n)
	}
	reset := func() {
		for i := range k {
			clear(dp[i])
			clear(fp[i])
		}
	}
	play := func(x int) string {
		// 当q中有x个1时
		reset()
		for i := range k {
			for w := range min(i+1, x) {
				dp[i][w] = 1
				i1 := i
				for j := w; j < n; j += x {
					if i1 >= m || s[j] != p[i1] {
						dp[i][w] = 0
						break
					}
					i1 += k
				}
				if i1 < m {
					// 因为key[i] = 1, 所以i1应该要被选中
					dp[i][w] = 0
				}
			}
		}
		// fp[i][w]
		for i := range k {
			for w := range min(i+1, x) {
				if i > 0 {
					fp[i][w] = fp[i-1][w]
					if w > 0 && dp[i][w] > 0 && fp[i-1][w-1] > 0 {
						fp[i][w] = 1
					}
				}
				if w == 0 && dp[i][w] > 0 {
					fp[i][w] = 1
				}
			}
		}
		if fp[k-1][x-1] == 0 {
			return ""
		}

		ans := make([]byte, k)
		for i := range k {
			ans[i] = '0'
		}
		for i := k - 1; i >= 0 && x > 0; i-- {
			if dp[i][x-1] == 1 {
				ans[i] = '1'
				x--
			}
		}

		return string(ans)
	}
	var ans string
	w := max(1, m/k)
	// w > 0
	x := n / w
	for x > 0 && x*(w+1) >= n {
		if x*w <= n {
			res := play(x)
			if len(res) > 0 {
				if len(ans) == 0 || res < ans {
					ans = res
				}
			}
		}
		x--
	}
	if len(ans) == 0 {
		ans = "0"
	}

	return ans
}
