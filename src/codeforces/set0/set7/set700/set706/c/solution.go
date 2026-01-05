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
	n := readNum(reader)
	c := readNNums(reader, n)
	words := make([]string, n)
	for i := range n {
		words[i] = readString(reader)
	}
	return solve(words, c)
}

const inf = 1 << 60

func solve(words []string, c []int) int {
	dp := make([]int, 2)
	dp[1] = c[0]

	reverse := func(s string, flag int) string {
		if flag == 0 {
			return s
		}
		buf := []byte(s)
		slices.Reverse(buf)
		return string(buf)
	}

	check := func(a string, b string) bool {
		return strings.Compare(a, b) <= 0
	}

	n := len(words)
	for i := 1; i < n; i++ {
		ndp := []int{inf, inf}
		for u := range 2 {
			for v := range 2 {
				if check(reverse(words[i-1], u), reverse(words[i], v)) {
					tmp := dp[u]
					if v == 1 {
						tmp += c[i]
					}
					ndp[v] = min(ndp[v], tmp)
				}
			}
		}
		copy(dp, ndp)
	}

	if min(dp[0], dp[1]) < inf {
		return min(dp[0], dp[1])
	}
	return -1
}
