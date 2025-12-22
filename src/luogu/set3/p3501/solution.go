package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	s := readString(reader)[:n]
	res := solve(s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

var BASE uint64 = 31

func solve(s string) int {
	n := len(s)
	f := make([]uint64, n+1)
	g := make([]uint64, n+2)
	p := make([]uint64, n+1)
	p[0] = 1
	for i := 1; i <= n; i++ {
		p[i] = BASE * p[i-1]
	}
	t := []byte(s)
	for i := 1; i <= n; i++ {
		if s[i-1] == '1' {
			t[i-1] = '0'
		} else {
			t[i-1] = '1'
		}
	}

	for i := 1; i <= n; i++ {
		x := uint64(s[i-1]-'0') + 1
		f[i] = f[i-1]*BASE + x
	}

	for i := n; i >= 1; i-- {
		x := uint64(t[i-1]-'0') + 1
		g[i] = g[i+1]*BASE + x
	}

	search := func(x int) int {
		l, r := 0, min(x, n-x)

		for l < r {
			mid := (l + r + 1) / 2
			if f[x]-f[x-mid]*p[mid] == g[x+1]-g[x+1+mid]*p[mid] {
				l = mid
			} else {
				r = mid - 1
			}
		}
		return l
	}

	var ans int

	for i := 1; i <= n; i++ {
		ans += search(i)
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
