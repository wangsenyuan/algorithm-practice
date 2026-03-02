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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	first := readString(reader)
	ss := strings.Split(first, " ")
	n, _ := strconv.Atoi(ss[0])
	m, _ := strconv.Atoi(ss[1])
	requirements := make([]string, m)
	for i := range m {
		requirements[i] = readString(reader)
	}
	return solve(n, requirements)
}

type data struct {
	l, r int
	sign string // = < > <= >=
}

func parse(s string) data {
	ss := strings.Split(s, " ")
	l, _ := strconv.Atoi(ss[0])
	r, _ := strconv.Atoi(ss[2])
	sign := ss[1]
	return data{l - 1, r - 1, sign}
}

func flip(sign string) string {
	if sign == "<" {
		return ">"
	}
	if sign == ">" {
		return "<"
	}
	if sign == "<=" {
		return ">="
	}
	if sign == ">=" {
		return "<="
	}
	return sign
}

func (d data) flip() data {
	return data{d.r, d.l, flip(d.sign)}
}

func solve(n int, requirements []string) int {
	pos := make([][]data, 2*n)

	for i := range requirements {
		tmp := parse(requirements[i])
		if tmp.sign[0] == '>' {
			tmp = tmp.flip()
		}

		if strings.HasSuffix(tmp.sign, "=") && tmp.l == tmp.r {
			// 这条规则是多余的
			continue
		}

		if tmp.l == tmp.r {
			return 0
		}
		// tmp.sign[0] == '=' or tmp.sign[0] == '<
		pos[tmp.l] = append(pos[tmp.l], tmp)
	}

	cmp := func(r int, cur data, L int, R int) bool {
		if cur.sign == "=" {
			return cur.r == r
		}
		if cur.sign == "<=" {
			return cur.r == r || L <= cur.r && cur.r <= R
		}
		// cur.sign == "<"
		return L <= cur.r && cur.r <= R
	}

	check := func(l int, r int, L int, R int) bool {
		for _, cur := range pos[l] {
			if !cmp(r, cur, L, R) {
				return false
			}
		}

		for _, cur := range pos[r] {
			if !cmp(l, cur, L, R) {
				return false
			}
		}

		return true
	}

	dp := make([][]int, 2*n)
	ndp := make([][]int, 2*n)
	for i := range 2 * n {
		dp[i] = make([]int, 2*n)
		ndp[i] = make([]int, 2*n)
		if i+1 < 2*n {
			// 看能否在 [i, i+1] 放置n
			if check(i, i+1, 2*n, -1) {
				dp[i][i+1] = 1
			}
		}
	}

	for range n - 1 {
		for r := range 2 * n {
			for l := range r {
				if dp[l][r] == 0 {
					continue
				}
				if l >= 2 && check(l-2, l-1, l, r) {
					ndp[l-2][r] += dp[l][r]
				}
				if l >= 1 && r+1 < 2*n && check(l-1, r+1, l, r) {
					ndp[l-1][r+1] += dp[l][r]
				}
				if r+2 < 2*n && check(r+1, r+2, l, r) {
					ndp[l][r+2] += dp[l][r]
				}
			}
		}
		for r := range 2 * n {
			for l := range r {
				dp[l][r] = ndp[l][r]
				ndp[l][r] = 0
			}
		}
	}

	return dp[0][2*n-1]
}
