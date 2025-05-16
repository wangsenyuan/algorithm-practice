package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const inf = 1 << 60

func process(reader *bufio.Reader) int {
	readString(reader)
	s := readString(reader)
	t := readString(reader)
	return solve(s, t)
}

func zf(s []byte, z []int) {
	n := len(s)
	clear(z[:n])
	var l, r int

	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && s[i+z[i]] == s[z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func solve(s string, t string) int {
	n := len(s)
	m := len(t)

	lpos := make([]int, m)
	rpos := make([]int, m)
	for i := range m {
		if i > 0 {
			lpos[i] = lpos[i-1] + 1
		}
		// lpos[i]是t[i]在s中最靠左边的位置，且在t[i-1]的后面
		for lpos[i] < n && s[lpos[i]] != t[i] {
			lpos[i]++
		}
		if lpos[i] == n {
			return -1
		}
	}
	for i := m - 1; i >= 0; i-- {
		if i < m-1 {
			rpos[i] = rpos[i+1] - 1
		} else {
			rpos[i] = n - 1
		}
		for s[rpos[i]] != t[i] {
			rpos[i]--
		}
		// 因为已经是非bad了，不会到负值
	}
	ans := inf
	buf := make([]byte, n+m+1)
	z := make([]int, n+m+1)
	for pos := range n + 1 {
		copy(buf, s[:pos])
		reverse(buf[:pos])
		buf[pos] = '#'
		copy(buf[pos+1:], t)
		k := pos + 1 + len(t)
		reverse(buf[pos+1 : k])
		zf(buf[:k], z)
		for suf := range m + 1 {
			if pos-suf < 0 {
				continue
			}
			if suf < m && rpos[suf] < pos {
				continue
			}
			if suf-1 >= 0 && lpos[suf-1] > pos {
				continue
			}
			var rg int
			if suf != 0 {
				sum := pos - z[k-suf] + pos - suf
				if sum != 0 {
					sum++
				}
				rg = sum
			} else {
				rg = pos
			}
			ans = min(ans, rg+n-pos)
		}
	}
	return ans
}

func solve2(s string, t string) int {
	n := len(s)
	m := len(t)
	f := make([][3]int, n+1)
	nf := make([][3]int, n+1)
	f[0][0] = 1
	for i, x := range s {
		if i < m {
			f[i+1] = [3]int{1e9, 1e9, 1e9}
		}
		nf[0][0] = f[0][0] + 2
		nf[0][1] = nf[0][0]
		nf[0][2] = f[0][2] + 1
		for j, y := range t[:min(i+1, m)] {
			nf[j+1][0] = f[j+1][0] + 2
			nf[j+1][1] = nf[j+1][0]
			nf[j+1][2] = min(f[j+1][2]+1, nf[j+1][0])
			if x == y {
				nf[j+1][0] = min(nf[j+1][0], f[j][0]+1)
				nf[j+1][1] = min(nf[j+1][1], f[j][1])
				nf[j+1][2] = min(nf[j+1][2], f[j][2]+1, f[j][1])
			}
		}
		f, nf = nf, f
	}
	ans := f[m][2]
	if ans > n {
		ans = -1
	}
	return ans
}

func solve1(s string, t string) int {
	// 从s中获取t
	n := len(s)
	m := len(t)
	if n < m {
		return -1
	}

	// 假设到达位置(i, j)，t[:i] = 处理后的s[:j]
	// 这个时候，t[i:] 应该包含在s[j:]中（否则就无效）
	// 从后往前移动并删除的距离，应该是t[i:]和s[j:]最长的匹配串的位置k, n - k
	fp := make([][]int, m+1)
	gp := make([][]int, m+1)
	for i := range m + 1 {
		fp[i] = make([]int, n+1)
		gp[i] = make([]int, n+1)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if t[i] == s[j] {
				fp[i][j] = fp[i+1][j+1] + 1
				gp[i][j] = gp[i+1][j+1] + 1
			} else {
				fp[i][j] = 0
				// 如果i/j不匹配，那么 i/(j+1)就必须匹配
				gp[i][j] = gp[i][j+1]
			}
		}
	}

	if gp[0][0] != m {
		return -1
	}

	// 从后面往前移动的最少的位置
	ans := n - fp[0][0]

	// dp可以被压缩，fp，表示从(i, j)开始的最远距离, 好像没法压缩呐
	dp := make([][]int, 2)
	for i := range 2 {
		dp[i] = make([]int, n+1)
		for j := range n + 1 {
			dp[i][j] = inf
		}
	}

	// dp[i][j]表示从头部开始移动到(i, j)且从t[:j]中得到s[:i]的最优解
	for j := 0; j <= n; j++ {
		// 要删掉s[:j]去和 t[:0]匹配
		dp[0][j] = j * 2
		if gp[0][j] == m {
			ans = min(ans, dp[0][j]+1+n-(j+fp[0][j]))
		}
	}

	for i := 1; i <= m; i++ {
		for j := range n + 1 {
			// reset it
			dp[i&1][j] = inf
		}
		for j := 1; j <= n; j++ {
			// 需要删除掉s[j-1]
			dp[i&1][j] = min(dp[i&1][j], dp[i&1][j-1]+2)
			if t[i-1] == s[j-1] {
				// 只需要移动不需要删除
				dp[i&1][j] = min(dp[i&1][j], dp[(i-1)&1][j-1]+1)
			}
			// 先跳到Home, 然后匹配到位置(i, j)
			// 移动到Home + 1
			tmp := dp[i&1][j] + 1 + n - (j + fp[i][j])
			if gp[i][j] == m-i {
				// s[j+1:]必须包含t[i+1:]
				ans = min(ans, tmp)
			}
		}
	}

	return ans
}
