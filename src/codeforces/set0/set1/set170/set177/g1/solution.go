package main

import (
	"bufio"
	"strconv"
	"strings"
)

func main() {

}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	res := make([]int, len(ss))
	for i := 0; i < len(ss); i++ {
		res[i], _ = strconv.Atoi(ss[i])
	}
	return res
}

func drive(reader *bufio.Reader) []int {
	nums := readNums(reader)
	k := nums[0]
	m := nums[1]
	queries := make([]string, m)
	for i := 0; i < m; i++ {
		queries[i] = readString(reader)
	}
	return solve(k, queries)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(k int, queries []string) []int {
	// 先计算s[i] >= 1000的几个序列
	var mxLen int
	for _, s := range queries {
		mxLen = max(mxLen, len(s))
	}
	// 计算出s[i] < 2 * maxLen
	f := []string{"a", "b"}
	for len(f) < k && len(f[len(f)-1]) < 4*mxLen {
		c := f[len(f)-1] + f[len(f)-2]
		f = append(f, c)
	}

	find := func(p []int, s string, w string) int {
		var cnt int
		var j int
		for i := 0; i < len(w); i++ {
			for j > 0 && w[i] != s[j] {
				j = p[j-1]
			}
			if w[i] == s[j] {
				j++
			}
			if j == len(s) {
				cnt++
				j = p[j-1]
			}
		}

		return cnt
	}

	play := func(s string) int {
		p := kmp(s)
		// 貌似不需要二分
		for i, x := range f {
			if len(x) > 2*len(s) {
				// s肯定发现不了
				break
			}
			if len(x) >= len(s) {
				c0 := find(p, s, x)
				if c0 == 0 {
					continue
				}
				// 有没有肯定i = len(f) - 1?
				if i == len(f)-1 {
					return c0
				}
				c1 := find(p, s, f[i+1])
				// what next?
			}
		}

		return 0
	}

	ans := make([]int, len(queries))

	for i, w := range queries {
		ans[i] = play(w)
	}

	return ans
}

func fib(a int, b int, n int) int {
	for range n {
		a, b = b, add(a, b)
	}

	return b
}

func kmp(s string) []int {
	res := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		j := res[i-1]
		for j > 0 && s[i] != s[j] {
			j = res[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		res[i] = j
	}
	return res
}
