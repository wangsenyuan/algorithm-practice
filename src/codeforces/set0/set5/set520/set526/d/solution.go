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

func drive(reader *bufio.Reader) string {
	_, k := readTwoNums(reader)
	s := readString(reader)
	return solve(k, s)
}

func solve(k int, s string) string {
	n := len(s)
	p := make([]int, n)

	var buf []byte
	if k > 1 {
		buf = append(buf, '0')
	} else {
		buf = append(buf, '1')
	}
	for i := 1; i < n; i++ {
		j := p[i-1]
		for j > 0 && s[i] != s[j] {
			j = p[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		p[i] = j

		s := i - j + 1
		if s > 0 {
			t := i%(i-j+1) + 1
			r := (i + 1) / s

			if s == t && r/k-r%k >= 0 || s != t && r/k-r%k > 0 {
				buf = append(buf, '1')
				continue
			}
		}
		buf = append(buf, '0')
	}

	return string(buf)
}

func solve1(k int, s string) string {
	n := len(s)
	z := make([]int, n)
	var l, r int
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	diff := make([]int, n+2)
	for t := 1; t <= n/k; t++ {
		lo := k * t
		if lo > n {
			break
		}
		h1 := (k + 1) * t
		h2 := t
		if t < n {
			h2 += z[t]
		}
		h := min(h1, h2, n)

		if lo <= h {
			diff[lo]++
			if h+1 <= n {
				diff[h+1]--
			}
		}
	}
	buf := make([]byte, n)
	for i := 1; i <= n; i++ {
		diff[i] += diff[i-1]
		if diff[i] > 0 {
			buf[i-1] = '1'
		} else {
			buf[i-1] = '0'
		}
	}
	return string(buf)
}
