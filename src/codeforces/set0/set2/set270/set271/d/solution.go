package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	s := readString(reader)
	flag := readString(reader)
	var k int
	fmt.Fscan(reader, &k)
	return solve(s, flag, k)
}

func solve(s string, flag string, k int) int {
	root := &Trie{}

	for l := 0; l < len(s); l++ {
		var cnt int
		tmp := root
		for r := l; r < len(s) && cnt <= k; r++ {
			if flag[s[r]-'a'] == '0' {
				cnt++
			}
			if cnt > k {
				break
			}
			tmp = tmp.Add(s[r])
		}
	}

	var ans int

	var dfs func(t *Trie)

	dfs = func(t *Trie) {
		if t == nil {
			return
		}
		ans++
		for i := range 26 {
			dfs(t.children[i])
		}
	}

	dfs(root)

	return ans - 1
}

type Trie struct {
	children [26]*Trie
}

func (t *Trie) Add(c byte) *Trie {
	x := int(c - 'a')
	if t.children[x] == nil {
		t.children[x] = &Trie{}
	}
	return t.children[x]
}
