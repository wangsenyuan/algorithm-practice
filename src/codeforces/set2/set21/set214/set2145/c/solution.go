package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	tc := readNum(reader)
	for range tc {
		readString(reader)
		s := readString(reader)
		res := solve(s)
		fmt.Fprintln(writer, res)
	}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	return s
}

func solve(s string) int {
	var diff int
	n := len(s)
	for i := range n {
		if s[i] == 'a' {
			diff++
		} else {
			diff--
		}
	}
	if diff == 0 {
		return 0
	}
	pos := make([]int, 2*n+1)
	for i := range 2*n + 1 {
		pos[i] = -1
	}

	ans := n

	pos[n] = 0
	var cur int
	for i := 1; i <= n; i++ {
		if s[i-1] == 'a' {
			cur++
		} else {
			cur--
		}
		// cur - prev = diff
		// prev = cur - diff
		if pos[n+cur-diff] != -1 {
			ans = min(ans, i-pos[n+cur-diff])
		}
		pos[n+cur] = i
	}
	if ans == n {
		return -1
	}
	return ans
}
