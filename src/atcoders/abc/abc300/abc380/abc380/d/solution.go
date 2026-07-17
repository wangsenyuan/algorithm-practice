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

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func drive(reader *bufio.Reader) string {
	s := readString(reader)
	q := readNum(reader)
	queries := readNNums(reader, q)
	return solve(s, queries)
}

func solve(s string, queries []int) string {
	n := len(s)

	var f func(k int, h int, flag int) byte

	f = func(k int, h int, flag int) byte {
		if k < n {
			// k < n
			if flag == 0 {
				return s[k]
			}
			return flip(s[k])
		}
		if k >= n*(1<<(h-1)) {
			return f(k-(1<<(h-1))*n, h-1, flag^1)
		}
		// h > 0
		return f(k, h-1, flag)
	}

	var res []byte

	for _, k := range queries {
		var h int
		for n*(1<<h) < k {
			h++
		}
		tmp := f(k-1, h, 0)
		if len(res) > 0 {
			res = append(res, ' ')
		}
		res = append(res, tmp)
	}

	return string(res)
}

func flip(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return byte(c - 'a' + 'A')
	}
	return byte(c - 'A' + 'a')
}
