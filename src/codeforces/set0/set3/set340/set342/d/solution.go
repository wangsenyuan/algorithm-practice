package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	readNum(reader)
	s := make([]string, 3)
	for i := range 3 {
		s[i] = readString(reader)
	}
	return solve(s)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(s []string) int {
	n := len(s[0])

	dp := make([]int, 8*2)
	dp[7*2] = 1
	ndp := make([]int, 16)

	get := func(j int) int {
		var res int
		for i := range 3 {
			if s[i][j] != '.' {
				res |= 1 << i
			}
		}
		return res
	}

	play := func(prev int, val int, f int, mask int, j int) {
		// 这里只考虑放置横向的部分, mask表示已经被放置的位置
		tmp := mask
		nf := f
		for i := range 3 {
			if (prev>>i)&1 == 0 {
				if (mask>>i)&1 == 1 {
					return
				}
				// 这里必须放置一个横向的
				tmp |= 1 << i
				if j-2 >= 0 && s[i][j-2] == 'O' {
					nf = 1
				}
				if j+1 < n && s[i][j+1] == 'O' {
					nf = 1
				}
			}
		}
		ndp[tmp*2+nf] = add(ndp[tmp*2+nf], val)
	}

	for i := range n {
		for prev := range 8 {
			for f := range 2 {
				if dp[prev*2+f] == 0 {
					continue
				}
				flag := get(i)
				if flag&3 == 0 {
					play(prev, dp[prev*2+f], f|check(s[2][i] == 'O'), flag|3, i)
				}
				if flag&6 == 0 {
					play(prev, dp[prev*2+f], f|check(s[0][i] == 'O'), flag|6, i)
				}
				play(prev, dp[prev*2+f], f, flag, i)
			}
		}
		copy(dp, ndp)
		clear(ndp)
	}

	return dp[15]
}

func solve1(s []string) int {
	n := len(s[0])

	dp := make([][][]int, n)
	for i := range n {
		dp[i] = make([][]int, 8)
		for j := range 8 {
			dp[i][j] = make([]int, 2)
			for k := range 2 {
				dp[i][j][k] = -1
			}
		}
	}

	use := make([]int, n)
	for i := range n {
		use[i] = 7
		for j := range 3 {
			if s[j][i] == '.' {
				use[i] ^= 1 << j
			}
		}
	}

	var dfs func(pos int, mask int, flag int) int
	dfs = func(pos int, mask int, flag int) (ans int) {
		if pos == n {
			return check(mask == 0 && flag == 1)
		}

		ret := &dp[pos][mask][flag]
		if *ret != -1 {
			return *ret
		}

		defer func() {
			*ret = ans
		}()

		mask |= use[pos]

		if mask == 7 {
			ans = dfs(pos+1, 0, flag)
			return
		}

		if mask&3 == 0 {
			ans = add(ans, dfs(pos, mask|3, flag|check(s[2][pos] == 'O')))
		}
		if mask&6 == 0 {
			ans = add(ans, dfs(pos, mask|6, flag|check(s[0][pos] == 'O')))
		}
		var i, nf, nm int
		for i < 3 {
			if (mask>>i)&1 == 0 && (pos == n-1 || (use[pos+1]>>i)&1 == 1) {
				break
			}

			if pos > 0 && s[i][pos-1] == 'O' || pos+2 < n && s[i][pos+2] == 'O' {
				nf = 1
			}
			if (mask>>i)&1 == 0 {
				nm |= (1 << i)
			}

			i++
		}
		if i == 3 {
			ans = add(ans, dfs(pos+1, nm, nf|flag))
		}

		return
	}

	res := dfs(0, 0, 0)
	return res
}

func check(f bool) int {
	if f {
		return 1
	}
	return 0
}
