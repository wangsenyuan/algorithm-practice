package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) string {
	x := readString(reader)
	xx := strings.Split(x, " ")
	t, _ := strconv.Atoi(xx[1])
	s := readString(reader)
	return solve(s, t)
}

const inf = 1 << 60

func solve(s string, t int) string {
	n := len(s)
	var dot int
	for dot < n && s[dot] != '.' {
		dot++
	}
	if dot == n {
		return s
	}
	f := s[dot+1:]
	m := len(f)
	dp := make([]int, m+1)
	dp[m] = inf
	for i := m - 1; i >= 0; i-- {
		if f[i] >= '5' {
			dp[i] = 0
		} else if f[i] < '4' {
			dp[i] = inf
		} else {
			dp[i] = dp[i+1] + 1
		}
	}

	w := s[:dot]

	if dp[0] < t && f[0] >= '4' {
		// 可以增加正数部分
		buf := []byte(w)
		i := len(w) - 1
		for i >= 0 && buf[i] == '9' {
			buf[i] = '0'
			i--
		}
		if i < 0 {
			return "1" + string(buf)
		}
		buf[i]++
		return string(buf)
	}

	buf := []byte(f)
	for i := range m - 1 {
		if dp[i+1] < t {
			buf[i]++
			return w + "." + string(buf[:i+1])
		}
	}

	return s
}
