package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	fmt.Println(res)
}

func solve(s string) string {
	n := len(s)

	// dp[i][j] = 最大的下标x，满足在i的前面（包括i）能够存在一个长度为2 * j 的回文序列
	pos := make([][]int, 26)
	for i, x := range []byte(s) {
		pos[x-'a'] = append(pos[x-'a'], i)
	}
	for x := range 26 {
		if len(pos[x]) >= 100 {
			u := fmt.Sprintf("%c", byte(x+'a'))
			return strings.Repeat(u, 100)
		}
	}
	// n <= 26 * 100
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for r := 0; r < n; r++ {
		for l := r - 1; l >= 0; l-- {
			if s[l] == s[r] {
				dp[l][r] = dp[l+1][r-1] + 2
			} else {
				dp[l][r] = max(dp[l+1][r], dp[l][r-1])
			}
		}
	}

	ans := min(dp[0][n-1], 100)
	tmp := ans
	var first string
	var second string
	for l, r := 0, n-1; l <= r; {
		if s[l] == s[r] && (l == r || dp[l][r] >= tmp) {
			first += string(s[l])
			tmp--
			if l < r {
				second += string(s[r])
				tmp--
			}
			l++
			r--
		} else if tmp <= dp[l+1][r] {
			l++
		} else {
			r--
		}

		if len(first)+len(second) == ans {
			break
		}
	}

	return first + reverse(second)
}

func solve1(s string) string {
	n := len(s)
	m := min(n/2, 50)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, m+1)
		for j := range m + 1 {
			dp[i][j] = -1
		}
		dp[i][0] = n
	}

	// dp[i][j] = 最大的下标x，满足在i的前面（包括i）能够存在一个长度为2 * j 的回文序列
	pos := make([][]int, 26)
	for i, x := range []byte(s) {
		pos[x-'a'] = append(pos[x-'a'], i)
	}

	for i := range n {
		if i > 0 {
			copy(dp[i], dp[i-1])
		}
		// 使用x
		x := int(s[i] - 'a')
		for j := min(i+1, m); j > 0; j-- {
			v := dp[i][j-1]
			// 在v的前面找到一个位置u,
			u := sort.SearchInts(pos[x], v)
			if u == len(pos[x]) || pos[x][u] >= v {
				u--
			}
			if u >= 0 && pos[x][u] > i && pos[x][u] < v {
				dp[i][j] = max(dp[i][j], pos[x][u])
			}
		}
	}

	pi, pl := -1, -1
	for i := range n {
		for j := m; j > 0; j-- {
			if i < dp[i][j] {
				l := 2 * j
				if i+1 < dp[i][j] {
					l++
				}
				l = min(l, 100)
				if l > pl {
					pl = l
					pi = i
				}
			}
		}
	}
	mid := ""
	if pl&1 == 1 {
		mid = s[pi+1 : pi+2]
		pl--
	}
	pl /= 2
	var half string
	for i := pi; i >= 0 && pl > 0; i-- {
		if s[i] == s[dp[i][pl]] {
			half += string(s[i])
			pl--
		}
	}

	return reverse(half) + mid + half
}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

func bruteForce(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for j := 0; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return min(dp[0][n-1], 100)
}
