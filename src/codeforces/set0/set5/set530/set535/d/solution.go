package main

import (
	"bufio"
	"fmt"
	"os"
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
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
	n, m := readTwoNums(reader)
	p := readString(reader)
	y := readNNums(reader, m)
	return solve(n, p, y)
}

const mod = 1e9 + 7

func mul(a, b int) int {
	return a * b % mod
}

func solve(n int, p string, y []int) int {
	next := kmp(p)
	m := len(p)

	good := make([]bool, m+1)

	for k := m; k > 0; k = next[k-1] {
		good[next[k-1]] = true
	}

	buf := make([]byte, n)
	for i := range n {
		buf[i] = '.'
	}

	var where int
	for i, j := 0, 0; i < n; i++ {
		if j < len(y) && i == y[j]-1 {
			// 需要放置一个新的字符
			if where <= i {
				copy(buf[i:], p)
			} else {
				// 如果重叠的 where - i的部分，不是前缀, 那么是非法的
				if !good[where-i] {
					return 0
				}
				// abcabc
				copy(buf[where:], p[where-i:])
			}
			where = i + m
			j++
		}
	}

	res := 1
	for i := range n {
		if buf[i] == '.' {
			res = mul(res, 26)
		}
	}

	return res
}

func kmp(s string) []int {
	n := len(s)
	p := make([]int, n)
	for i := 1; i < n; i++ {
		j := p[i-1]
		for j > 0 && s[i] != s[j] {
			j = p[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		p[i] = j
	}
	return p
}
