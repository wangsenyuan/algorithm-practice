package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	k := readNum(reader)
	s := readString(reader)
	return solve(k, s)
}

func solve(k int, s string) int {
	n := len(s)
	// 下一个可以分割的位置
	next := make([]int, n)
	special := n
	for i := n - 1; i >= 0; i-- {
		next[i] = special
		if s[i] == '-' || s[i] == ' ' {
			special = i
		}
	}

	getWidth := func(pos int, end int) int {
		w := end - pos
		if end < n {
			// s[end] = ' ' or '-'
			w++
		}
		return w
	}

	check := func(w int) bool {
		var row int
		// 一行有w个字符
		var pos int
		for pos < n {
			j := pos
			for j < n && getWidth(pos, next[j]) <= w {
				j = next[j]
			}
			if j == pos {
				// 这个单词太长了
				return false
			}
			row++
			pos = j + 1
		}

		return row <= k
	}

	return sort.Search(n, check)
}
