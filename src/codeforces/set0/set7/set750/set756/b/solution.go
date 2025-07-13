package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	t := make([]int, n)
	for i := range n {
		t[i] = readNum(reader)
	}
	return solve(t)
}

func solve(t []int) []int {
	n := len(t)

	d := make([]int, n+1)
	copy(d[1:], t)

	ans := make([]int, n)

	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 20
		j := search(0, i, func(j int) bool {
			return d[i]-d[j] < 90
		})
		if j > 0 {
			j--
		}
		dp[i] = min(dp[i], dp[j]+50)
		j = search(0, i, func(j int) bool {
			return d[i]-d[j] < 1440
		})
		if j > 0 {
			j--
		}
		dp[i] = min(dp[i], dp[j]+120)
	}
	for i := range n {
		ans[i] = dp[i+1] - dp[i]
	}
	return ans
}

func search(l int, r int, f func(int) bool) int {
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
