package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	cnt, res := process(reader)
	fmt.Println(cnt)
	fmt.Println(res[0], res[1])
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

func process(reader *bufio.Reader) (cnt int, res []int) {
	n, _ := readTwoNums(reader)
	puzzle := make([]string, n)
	for i := range n {
		puzzle[i] = readString(reader)
	}
	return solve(puzzle)
}

func solve(puzzle []string) (cnt int, res []int) {
	n := len(puzzle)
	m := len(puzzle[0])

	get := func(r0 int, r1 int, c0 int, c1 int) string {
		buf := make([]byte, 0, (r1-r0)*(c1-c0))
		for i := r0; i < r1; i++ {
			for j := c0; j < c1; j++ {
				buf = append(buf, puzzle[i][j])
			}
		}
		s := string(buf)
		s0 := string(reverse(buf))
		if s0 < s {
			s = s0
		}

		if r1-r0 == c1-c0 {
			buf = buf[:0]
			for j := c1 - 1; j >= c0; j-- {
				for i := r0; i < r1; i++ {
					buf = append(buf, puzzle[i][j])
				}
			}
			s1 := string(buf)
			if s1 < s {
				s = s1
			}
			s2 := string(reverse(buf))
			if s2 < s {
				s = s2
			}
		}

		return s
	}

	check := func(a, b int) bool {
		freq := make(map[string]int)

		for i := 0; i < n; i += a {
			for j := 0; j < m; j += b {
				s := get(i, i+a, j, j+b)
				freq[s]++
				if freq[s] > 1 {
					return false
				}
			}
		}
		return true
	}

	type pair struct {
		first  int
		second int
	}

	all := make(map[pair]int)

	sz := n * m
	res = []int{n, m}

	checkAndUpdate := func(a, b int) {
		if check(a, b) {
			all[pair{a, b}]++
			if a*b < sz || a*b == sz && a < res[0] {
				sz = a * b
				res[0] = a
				res[1] = b
			}
		}
	}

	for a := 1; a <= n; a++ {
		if n%a != 0 {
			continue
		}
		for b := 1; b <= m; b++ {
			if m%b != 0 {
				continue
			}
			checkAndUpdate(a, b)
		}
	}

	cnt = len(all)

	return
}

func reverse(arr []byte) []byte {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
