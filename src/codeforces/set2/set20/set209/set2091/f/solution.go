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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	return s
}

func process(reader *bufio.Reader) int {
	n, m, d := readThreeNums(reader)
	wall := make([]string, n)
	for i := 0; i < n; i++ {
		wall[i] = readString(reader)[:m]
	}
	return solve(wall, d)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(wall []string, d int) int {
	// 从上到下也是一样的
	n := len(wall)
	m := len(wall[0])

	dp := make([]int, m)
	fp := make([]int, m)

	for j := 0; j < m; j++ {
		if wall[0][j] == 'X' {
			dp[j] = 1
		}
	}
	var sum int
	for j, l, r := 0, 0, 0; j < m; j++ {
		if wall[0][j] == '#' {
			continue
		}
		for l < j && j-l > d {
			sum = sub(sum, dp[l])
			l++
		}
		for r < m && r-j <= d {
			sum = add(sum, dp[r])
			r++
		}
		fp[j] = sum
	}

	checkDist := func(x int, y int) bool {
		return (x-y)*(x-y)+1 <= d*d
	}

	for i := 1; i < n; i++ {
		clear(dp)
		var sum int

		for j, l, r := 0, 0, 0; j < m; j++ {
			if wall[i][j] == '#' {
				continue
			}
			// wall[i][j] = 'X'
			for r < m && (r <= j || checkDist(r, j)) {
				sum = add(sum, fp[r])
				r++
			}
			for l < j && !checkDist(j, l) {
				sum = sub(sum, fp[l])
				l++
			}
			dp[j] = sum
		}
		clear(fp)
		sum = 0
		for j, l, r := 0, 0, 0; j < m; j++ {
			if wall[i][j] == '#' {
				continue
			}
			for l < j && j-l > d {
				sum = sub(sum, dp[l])
				l++
			}
			for r < m && r-j <= d {
				sum = add(sum, dp[r])
				r++
			}
			fp[j] = sum
		}
	}

	var res int
	for j := 0; j < m; j++ {
		res = add(res, fp[j])
	}
	return res
}
