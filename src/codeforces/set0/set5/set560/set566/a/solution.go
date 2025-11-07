package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, score, res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, score)
	for _, cur := range res {
		fmt.Fprintln(writer, cur[0], cur[1])
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

func drive(reader *bufio.Reader) (names []string, pseudonym []string, score int32, res [][]int32) {
	n := readNum(reader)
	names = make([]string, n)
	for i := 0; i < n; i++ {
		names[i] = readString(reader)
	}
	pseudonym = make([]string, n)
	for i := 0; i < n; i++ {
		pseudonym[i] = readString(reader)
	}
	score, res = solve(names, pseudonym)
	return
}

func solve(names []string, pseudonym []string) (int32, [][]int32) {

	t1 := NewTrie()

	for i, cur := range names {
		t1.Add(cur, int32(i))
	}

	t2 := NewTrie()
	for i, cur := range pseudonym {
		t2.Add(cur, int32(i))
	}

	var sum int32
	var res [][]int32

	var dfs func(a *Trie, b *Trie, d int) ([]int32, []int32)

	// 返回的是还没有被匹配的
	dfs = func(a *Trie, b *Trie, d int) (x []int32, y []int32) {
		if a == nil && b == nil {
			return nil, nil
		}

		if a != nil {
			x = a.val
		}
		if b != nil {
			y = b.val
		}

		for i := range 26 {
			var u, v *Trie
			if a != nil {
				u = a.children[i]
			}
			if b != nil {
				v = b.children[i]
			}
			l, r := dfs(u, v, d+1)
			x = append(x, l...)
			y = append(y, r...)
		}
		w := min(len(x), len(y))
		sum += int32(w) * int32(d)
		for i := range w {
			res = append(res, []int32{x[i] + 1, y[i] + 1})
		}
		x = x[w:]
		y = y[w:]

		return x, y
	}

	dfs(t1, t2, 0)

	return sum, res
}

type Trie struct {
	children []*Trie
	val      []int32
}

func NewTrie() *Trie {
	return &Trie{
		children: make([]*Trie, 26),
	}
}

func (tr *Trie) Add(s string, v int32) {
	if len(s) == 0 {
		tr.val = append(tr.val, v)
		return
	}
	x := int(s[0] - 'a')
	if tr.children[x] == nil {
		tr.children[x] = NewTrie()
	}
	tr.children[x].Add(s[1:], v)
}
