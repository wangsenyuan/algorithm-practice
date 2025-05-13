package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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

func process(reader *bufio.Reader) int {
	s := readString(reader)
	n := readNum(reader)
	words := make([]string, n)
	for i := range n {
		words[i] = readString(reader)
	}
	return solve(s, words)
}

func solve(s string, words []string) int {
	n := len(s)

	rs := reverse(s)

	get := func(w string, s string) []int {
		p := kmp(w)
		m := len(w)
		pos := make([]int, m+1)
		var j int
		for i := 0; i < n; i++ {
			for j > 0 && s[i] != w[j] {
				j = p[j-1]
			}
			if s[i] == w[j] {
				j++
			}
			if pos[j] == 0 {
				pos[j] = i + 1
			}
			if j == m {
				break
			}
		}

		return pos
	}

	check := func(w string) bool {
		if len(w) == 1 {
			// 至少要两段
			return false
		}
		p1 := get(w, s)
		m := len(w)
		if p1[m] > 0 {
			// 有一个完整的
			return true
		}
		p2 := get(reverse(w), rs)
		for i := 1; i < m; i++ {
			if p1[i] >= 0 {
				l := p1[i]
				r := p2[m-i]
				if l > 0 && r > 0 && l+r <= n {
					return true
				}
			}
		}
		return false
	}

	var ans int
	for _, w := range words {
		if check(w) {
			ans++
		}
	}
	return ans

}

func reverse(s string) string {
	buf := []byte(s)
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
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
