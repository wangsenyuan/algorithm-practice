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
	tc := readNum(reader)
	var buf bytes.Buffer
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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
	n := readNum(reader)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = readString(reader)
	}
	return solve(words)
}

func solve(words []string) int {
	slices.SortFunc(words, func(a, b string) int {
		if len(a) != len(b) {
			return len(b) - len(a)
		}
		return strings.Compare(a, b)
	})

	tr := NewTrie()

	pos := make([][][]int32, len(words))

	for i, cur := range words {
		var where int32
		for j := 0; j < len(cur); j++ {
			where = tr.Add(where, cur[j])
			pos[i] = append(pos[i], make([]int32, 27))
			pos[i][j][26] = where
		}
	}

	tmp := make([]int32, 26)

	for i, cur := range words {
		clear(tmp)
		for j := range 26 {
			tmp[j] = tr.Find(tmp[j], j)
		}
		for j := range len(cur) {
			for c := range 26 {
				if tmp[c] != 0 {
					tmp[c] = tr.Find(tmp[c], int(cur[j]-'a'))
				}
			}
			copy(pos[i][j], tmp)
		}
	}

	m := len(words[0])
	at := make([][]int, m+1)
	var k int
	for i := range words {
		at[len(words[i])] = append(at[len(words[i])], i)
		k = max(k, int(pos[i][len(words[i])-1][26]))
	}

	marked := make([]bool, k+1)

	var res int

	for l := m; l > 0; l-- {
		for _, i := range at[l] {
			ok := !marked[pos[i][l-1][26]]
			for j := 0; j < 26 && ok; j++ {
				if pos[i][l-1][j] > 0 && marked[pos[i][l-1][j]] {
					ok = false
				}
			}
			if ok {
				res++
				marked[pos[i][l-1][26]] = true
			}
			// 将它移动到下一个位置
			at[l-1] = append(at[l-1], i)
		}
	}

	return res
}

type Trie struct {
	children [][26]int32
}

func NewTrie() *Trie {
	return &Trie{
		children: make([][26]int32, 1),
	}
}

func (tr *Trie) Add(pos int32, c byte) int32 {
	x := int(c - 'a')
	if tr.children[pos][x] == 0 {
		tr.children[pos][x] = int32(len(tr.children))
		tr.children = append(tr.children, [26]int32{})
	}
	return tr.children[pos][x]
}

func (tr *Trie) Find(pos int32, x int) int32 {
	return tr.children[pos][x]
}
