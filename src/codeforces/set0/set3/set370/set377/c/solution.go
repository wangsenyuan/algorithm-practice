package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	n := readNum(reader)
	a := readNNums(reader, n)
	m := readNum(reader)
	ops := make([][]int, m)
	for i := range m {
		s := readString(reader)
		ops[i] = make([]int, 2)
		if s[0] == 'p' {
			ops[i][0] = 1
		} else {
			ops[i][0] = 2
		}
		ops[i][1] = int(s[2] - '0')
	}
	return solve(a, ops)
}

const inf = 1 << 60

func solve(a []int, ops [][]int) int {
	slices.Sort(a)
	slices.Reverse(a)
	// n := len(a)
	m := len(ops)
	n := min(len(a), m)
	// n 超过m没有意义
	a = a[:n]

	dp := make([][]int, m)
	ok := make([][]bool, m)
	for i := range m {
		dp[i] = make([]int, 1<<n)
		ok[i] = make([]bool, 1<<n)
	}

	sign := func(b bool) int {
		if b {
			return 1
		}
		return -1
	}

	var f func(i int, mask int) int

	f = func(i int, mask int) (res int) {
		if i == m {
			return 0
		}
		if ok[i][mask] {
			return dp[i][mask]
		}
		defer func() {
			dp[i][mask] = res
			ok[i][mask] = true
		}()

		if ops[i][0] == 1 {
			// to pick
			for j := range n {
				if (mask>>j)&1 == 0 {
					res = sign(ops[i][1] == 1)*a[j] + f(i+1, mask|(1<<j))
					return
				}
			}
			// all baned?
			return 0
		}
		// to ban
		res = inf * sign(ops[i][1] == 2)
		for j := range n {
			if (mask>>j)&1 == 0 {
				if ops[i][1] == 1 {
					res = max(res, f(i+1, mask|(1<<j)))
				} else {
					res = min(res, f(i+1, mask|(1<<j)))
				}
			}
		}
		return
	}

	return f(0, 0)
}
