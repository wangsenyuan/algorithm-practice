package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, k := readTwoNums(reader)
	s := readString(reader)
	fmt.Println(solve(k, s))
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

func solve(k int, s string) string {
	n := len(s)
	buf := make([]byte, n+1)
	copy(buf[1:], s)
	next := make([]int, n+2)
	next[n] = n + 1
	for i := n - 1; i > 0; i-- {
		if buf[i] == '4' && buf[i+1] == '7' {
			next[i] = i
		} else {
			next[i] = next[i+1]
		}
	}

	buf[0] = '0'

	if next[1] == n+1 {
		return s
	}
	pos := 1
	for k > 0 && pos < n {
		pos = next[pos]
		k--
		if pos%2 == 1 {
			// change to 44
			buf[pos] = '4'
			buf[pos+1] = '4'
			if pos+2 <= n && buf[pos+2] == '7' {
				next[pos+1] = pos + 1
			}
		} else {
			buf[pos] = '7'
			buf[pos+1] = '7'
			if pos > 1 && buf[pos-1] == '4' {
				pos = pos - 1
				break
			}
		}
		pos = next[pos+1]
	}
	k %= 2
	if k == 1 && pos < n {
		if pos%2 == 1 {
			buf[pos] = '4'
			buf[pos+1] = '4'
		} else {
			buf[pos] = '7'
			buf[pos-1] = '7'
		}
	}

	return string(buf[1:])
}
