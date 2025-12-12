package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, v := range res {
		if v {
			writer.WriteString("YES\n")
		} else {
			writer.WriteString("NO\n")
		}
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
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

func drive(reader *bufio.Reader) []bool {
	good := readString(reader)
	pat := readString(reader)
	m := readNum(reader)
	words := make([]string, m)
	for i := 0; i < m; i++ {
		words[i] = readString(reader)
	}
	return solve(good, pat, words)
}

func solve(good string, pat string, words []string) []bool {
	var flag int
	for _, v := range good {
		flag |= 1 << int(v-'a')
	}

	// * 只会出现一次，？必须被替换，所以可以尽量匹配前后缀
	starPos := -1
	for i := range pat {
		if pat[i] == '*' {
			starPos = i
			break
		}
	}

	n := len(pat)

	check := func(s string) bool {
		m := len(s)
		var l int
		for l < m && l < n && (pat[l] == s[l] || pat[l] == '?' && (flag>>int(s[l]-'a'))&1 == 1) {
			l++
		}

		var r int
		for r < m && r < n && (pat[n-1-r] == s[m-1-r] || pat[n-1-r] == '?' && (flag>>int(s[m-1-r]-'a'))&1 == 1) {
			r++
		}
		if starPos == -1 {
			if l+r < n {
				return false
			}
			// l + r >= n
			return m == n
		}
		// starPos == -1
		if l+r+1 < n || m+1 < n {
			return false
		}
		// l + r + 1 == n, 正好starPos
		for i := l; i+r < m; i++ {
			if (flag>>int(s[i]-'a'))&1 == 1 {
				return false
			}
		}

		return true
	}

	ans := make([]bool, len(words))
	for i, w := range words {
		ans[i] = check(w)
	}
	return ans
}
