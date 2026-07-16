package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	data, _ := io.ReadAll(reader)
	// \r\n需要处理
	var pos int
	for i := 0; i < len(data); i++ {
		j := getLetterId(data[i])
		if j >= 0 {
			data[pos] = data[i]
			pos++
		}
	}
	return solve(string(data[:pos]))
}

func getLetterId(x byte) int {
	if x >= 'a' && x <= 'z' {
		return int(x - 'a')
	}
	if x == '-' {
		return 26
	}

	if x == '\'' {
		return 27
	}

	if x == '.' {
		return 28
	}
	if x == ',' {
		return 29
	}
	if x == '!' {
		return 30
	}
	if x == '?' {
		return 31
	}
	if x == ' ' {
		return 32
	}
	if x == '\n' {
		return 33
	}

	return -1
}

type trie struct {
	nodes [][34]int
	cnt   []int
	leaf  []bool
}

func (tr *trie) next() int {
	tr.nodes = append(tr.nodes, [34]int{})
	tr.cnt = append(tr.cnt, 0)
	tr.leaf = append(tr.leaf, false)
	return len(tr.nodes) - 1
}

func solve(text string) int {
	n := len(text)
	var save int

	var tr trie
	tr.next()

	vis := make(map[string]bool)

	var autoComplete int
	for i := 0; i < n; i++ {
		// read to next word separator
		j := i
		for i < n {
			id := getLetterId(text[i])
			if id >= 26 {
				break
			}
			i++
		}
		if i == j {
			continue
		}
		first := !vis[text[j:i]]
		vis[text[j:i]] = true
		var root int
		var tmp int
		for i1 := j; i1 < i; i1++ {
			id := getLetterId(text[i1])
			if tr.nodes[root][id] == 0 {
				tr.nodes[root][id] = tr.next()
			} else {
				nxt := tr.nodes[root][id]
				if tr.cnt[nxt] == 1 {
					tmp++
				}
				if tr.leaf[nxt] && tmp > 0 {
					if tmp > 1 {
						save += tmp - 1
						autoComplete++
					}
					tmp = 0
				}
			}

			nxt := tr.nodes[root][id]
			if first {
				tr.cnt[nxt]++
			}
			root = nxt
		}
		tr.leaf[root] = true
	}

	return n - save + autoComplete
}
