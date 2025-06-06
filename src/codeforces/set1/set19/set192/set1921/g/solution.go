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
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) int {
	n, _, k := readThreeNums(reader)
	grid := make([]string, n)
	for i := range n {
		grid[i] = readString(reader)
	}
	return solve(grid, k)
}

func solve(grid []string, k int) int {
	var ans int
	for range 4 {
		ans = max(ans, play(grid, k))
		grid = rotate(grid)
	}
	return ans
}

func rotate(buf []string) []string {
	n := len(buf)
	m := len(buf[0])
	res := make([][]byte, m)
	for i := range m {
		res[i] = make([]byte, n)
	}

	for i := range n {
		for j := range m {
			res[j][i] = buf[n-1-i][j]
		}
	}

	a := make([]string, m)
	for i := range m {
		a[i] = string(res[i])
	}

	return a
}

func play(grid []string, k int) int {
	// 始终从左上到右下
	n := len(grid)
	m := len(grid[0])
	// 最大是n
	dias := make(BIT, n*m+10)
	cnt := make([]int, n+m-1)
	for i := range n {
		for j := range m {
			cnt[i+j]++
		}
	}
	pos := make([]int, n+m)
	for i := 1; i < n+m; i++ {
		pos[i] = pos[i-1] + cnt[i-1]
	}

	where := func(i int, j int) int {
		p := pos[i+j]
		if i+j < m {
			return p + i
		}
		return p + m - j - 1
	}

	for i := range n {
		for j := range m {
			if grid[i][j] == '#' {
				u := where(i, j)
				dias.Set(u, 1)
			}
		}
	}

	calc := func(i int, j int) int {
		// 计算以(i, j)为中心的第k层的数量
		id := i + j - k
		// 这个好难算呐
		u := pos[id]
		if i-k > 0 {
			u = where(i-k, j)
		}
		v := pos[id] + cnt[id] - 1
		if j-k > 0 {
			v = where(i, j-k)
		}
		return dias.Query(u, v)
	}

	var ans int
	dp := make([]int, m)
	ndp := make([]int, m)

	for i := range n {
		clear(ndp)
		var sum int
		for j := range m {
			if grid[i][j] == '#' {
				sum++
			}
			if j-k > 0 && grid[i][j-k-1] == '#' {
				sum--
			}
			ndp[j] = sum + dp[j]
			if i > 0 && i+j > k {
				ndp[j] -= calc(i-1, j)
			}
			ans = max(ans, ndp[j])
		}
		copy(dp, ndp)
	}

	return ans
}

type BIT []int

func (bit BIT) Set(p int, v int) {
	p++
	for p < len(bit) {
		bit[p] += v
		p += p & -p
	}
}

func (bit BIT) Get(p int) int {
	p++
	var res int
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}

func (bit BIT) Query(l int, r int) int {
	return bit.Get(r) - bit.Get(l-1)
}
