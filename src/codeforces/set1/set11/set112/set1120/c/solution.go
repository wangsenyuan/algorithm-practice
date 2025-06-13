package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	fmt.Println(ans)
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

func process(reader *bufio.Reader) int {
	_, a, b := readThreeNums(reader)
	s := readString(reader)
	return solve(s, a, b)
}

const inf = 1 << 60

func solve(s string, a int, b int) int {
	n := len(s)
	dp := make([]int, n)
	for i := range n {
		dp[i] = inf
	}
	dp[0] = a
	p := make([]int, n)

	for i := 1; i < n; i++ {
		dp[i] = min(dp[i], dp[i-1]+a)
		ln := min(i, n-i)
		kmp(s[i:i+ln], p)
		var k int
		for j := 0; j < i; j++ {
			for k > 0 && s[j] != s[i+k] {
				k = p[k-1]
			}
			if s[j] == s[i+k] {
				k++
			}
			if k > 0 {
				dp[i+k-1] = min(dp[i+k-1], dp[i-1]+b)
			}
			if k == ln {
				k = p[k-1]
			}
		}
	}
	return dp[n-1]
}

func kmp(s string, p []int) {
	n := len(s)
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
}
