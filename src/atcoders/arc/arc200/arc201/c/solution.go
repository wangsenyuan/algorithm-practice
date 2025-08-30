package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func drive(reader *bufio.Reader) []int {
	n := readNum(reader)
	S := make([]string, n)
	for i := 0; i < n; i++ {
		S[i] = readString(reader)
	}
	return solve(S)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func solve(S []string) []int {
	tr := &Trie{}
	tr.expand()
	tr.expand()

	ans := make([]int, len(S))
	for i, w := range S {
		tr.Add(w)
		ans[i] = tr.val[1]
	}
	return ans
}

type Trie struct {
	children [][2]int
	sz       []int
	cnt      []int
	val      []int
}

func (t *Trie) expand() int {
	t.children = append(t.children, [2]int{})
	t.sz = append(t.sz, 0)
	t.val = append(t.val, 0)
	t.cnt = append(t.cnt, 0)
	return len(t.children) - 1
}

func (t *Trie) Add(s string) {

	var dfs func(node int, i int)
	dfs = func(node int, i int) {
		if i == len(s) {
			t.val[node] = add(t.val[node], pow(2, t.sz[node]))
			t.sz[node]++
			t.cnt[node]++
			return
		}
		x := int(s[i] - 'A')
		if t.children[node][x] == 0 {
			j := t.expand()
			t.children[node][x] = j
		}
		dfs(t.children[node][x], i+1)
		l := t.children[node][x]
		r := t.children[node][x^1]
		t.sz[node] = t.sz[l] + t.sz[r]
		t.val[node] = mul(t.val[l], t.val[r])
		if t.cnt[node] > 0 {
			t.val[node] = add(t.val[node], pow(2, t.sz[node]))
		}
		t.sz[node] += t.cnt[node]
	}

	dfs(1, 0)
}
