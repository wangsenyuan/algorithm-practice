package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
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
	s = strings.TrimSpace(s)
	return s
}

func drive(reader *bufio.Reader) (s string, words []string, res string) {
	readNum(reader)
	s = readString(reader)
	m := readNum(reader)
	words = make([]string, m)
	for i := range m {
		words[i] = readString(reader)
	}
	res = solve(s, words)
	return
}

type pair struct {
	first  int32
	second int32
}

func solve(s string, words []string) string {
	tr := &Trie{}
	// add root
	tr.next()

	m := len(words)

	pos := make([]int32, m)

	var marked []pair

	for i, cur := range words {
		cur = change(cur)
		// words[i] = cur
		// 如果在后面读到了pos[i]
		pos[i] = int32(tr.Add(cur))
		marked = append(marked, pair{pos[i], int32(i)})
	}

	k := len(tr.children)

	stop := make([]int32, k+1)
	for i := range k + 1 {
		stop[i] = -1
	}
	for _, p := range marked {
		stop[p.first] = p.second
	}

	n := len(s)
	dp := make([]int32, n+1)
	dp[n] = 0
	for i := n - 1; i >= 0; i-- {
		dp[i] = -1
		var cur int32
		for j := i; j < n; j++ {
			x := int(s[j] - 'a')
			cur = tr.children[cur][x]
			if cur == 0 {
				break
			}

			if dp[j+1] >= 0 && stop[cur] >= 0 {
				// 检查j是不是一个stop word
				dp[i] = stop[cur]
				break
			}
		}
	}

	var buf strings.Builder

	for i := 0; i < n; {
		j := dp[i]
		buf.WriteString(words[j])
		if i+len(words[j]) < n {
			buf.WriteByte(' ')
		}
		i += len(words[j])
	}

	return buf.String()

}

func change(s string) string {
	buf := []byte(bytes.ToLower([]byte(s)))
	slices.Reverse(buf)
	return string(buf)
}

type Trie struct {
	children [][26]int32
}

func (tr *Trie) next() int32 {
	id := len(tr.children)
	tr.children = append(tr.children, [26]int32{})
	return int32(id)
}

func (tr *Trie) Add(word string) int32 {
	var cur int32
	for i := range len(word) {
		x := int(word[i] - 'a')
		if tr.children[cur][x] == 0 {
			tr.children[cur][x] = tr.next()
		}
		cur = tr.children[cur][x]
	}
	return cur
}
