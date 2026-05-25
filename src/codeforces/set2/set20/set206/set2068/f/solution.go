package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	s := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}
	var t string
	fmt.Fscan(reader, &t)

	return solve(s, t)
}

func solve(s []string, t string) string {
	ques := make([][]int, 26)
	for i, x := range s {
		if containsSubsequence(x, t) {
			return ""
		}
		c := int(x[0] - 'a')
		ques[c] = append(ques[c], i)
	}

	tot := len(s)

	play := func(c int) {
		var todo []int
		for _, id := range ques[c] {
			s[id] = s[id][1:]
			if len(s[id]) == 0 {
				tot--
			} else {
				todo = append(todo, id)
			}
		}
		ques[c] = ques[c][:0]
		for _, id := range todo {
			d := int(s[id][0] - 'a')
			ques[d] = append(ques[d], id)
		}
	}

	var buf bytes.Buffer
	for tot > 0 {
		ok := false
		// 这个过程太慢了

		for c := range 26 {
			if c != int(t[0]-'a') && len(ques[c]) > 0 {
				buf.WriteByte(byte(c + 'a'))
				play(c)
				ok = true
				break
			}
		}
		if !ok {
			// have no choice, use t[0]
			buf.WriteByte(t[0])
			c := int(t[0] - 'a')
			t = t[1:]
			// 只有一个len(ques[c]) > 0
			play(c)
		}
	}

	return buf.String()
}

func containsSubsequence(s, t string) bool {
	for i, j := 0, 0; i < len(t); i++ {
		for j < len(s) && t[i] != s[j] {
			j++
		}
		if j == len(s) {
			return false
		}
		j++
	}

	return true
}
