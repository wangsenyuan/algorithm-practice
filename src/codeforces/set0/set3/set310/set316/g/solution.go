package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const maxn = 1_000_005

// SuffixAutomaton represents the suffix automaton structure
type SuffixAutomaton struct {
	ch  [maxn][26]int32
	len [maxn]int32
	fa  [maxn]int32
	tot int32
}

// NewSuffixAutomaton creates a new suffix automaton
func NewSuffixAutomaton() *SuffixAutomaton {
	sa := &SuffixAutomaton{tot: 1}
	return sa
}

// insert adds a character to the suffix automaton
func (sa *SuffixAutomaton) insert(p, c int32) int32 {
	if sa.ch[p][c] != 0 && sa.len[p]+1 == sa.len[sa.ch[p][c]] {
		return sa.ch[p][c]
	}
	sa.tot++
	now := sa.tot
	chk := sa.ch[p][c]
	sa.len[now] = sa.len[p] + 1

	for p != 0 && sa.ch[p][c] == 0 {
		sa.ch[p][c] = now
		p = sa.fa[p]
	}

	if p == 0 {
		sa.fa[now] = 1
		return now
	}

	q := sa.ch[p][c]
	if sa.len[p]+1 == sa.len[q] {
		sa.fa[now] = q
		return now
	}
	sa.tot++
	nq := sa.tot
	sa.len[nq] = sa.len[p] + 1

	for i := 0; i < 26; i++ {
		sa.ch[nq][i] = sa.ch[q][i]
	}

	sa.fa[nq] = sa.fa[q]
	sa.fa[now] = nq
	sa.fa[q] = nq

	for sa.ch[p][c] == q {
		sa.ch[p][c] = nq
		p = sa.fa[p]
	}

	if chk != 0 {
		return nq
	}
	return now
}

type Rule struct {
	pattern string
	l, r    int
}

func solve(s string, rules []Rule) int {
	sa := NewSuffixAutomaton()
	lst := int32(1)
	sz := make([][]int, maxn)
	for i := range maxn {
		sz[i] = make([]int, 11)
	}
	for i := range s {
		x := int32(s[i] - 'a')
		lst = sa.insert(lst, x)
		sz[lst][0]++
	}

	for i := range rules {
		lst := int32(1)
		for j := range rules[i].pattern {
			x := int32(rules[i].pattern[j] - 'a')
			lst = sa.insert(lst, x)
			sz[lst][i+1]++
		}
	}
	buc := make([]int32, sa.tot+1)
	for i := int32(1); i <= sa.tot; i++ {
		buc[sa.len[i]]++
	}
	for i := int32(1); i <= sa.tot; i++ {
		buc[i] += buc[i-1]
	}
	id := make([]int32, sa.tot+1)
	for i := sa.tot; i > 0; i-- {
		id[buc[sa.len[i]]] = i
		buc[sa.len[i]]--
	}
	n := len(rules)
	for i := sa.tot; i >= 1; i-- {
		for j := 0; j <= n; j++ {
			sz[sa.fa[id[i]]][j] += sz[id[i]][j]
		}
	}
	var ans int
	for i := int32(1); i <= sa.tot; i++ {
		if sz[i][0] > 0 {
			ok := true
			for j := 1; j <= n; j++ {
				if sz[i][j] < rules[j-1].l || sz[i][j] > rules[j-1].r {
					ok = false
					break
				}
			}
			if ok {
				ans += int(sa.len[i] - sa.len[sa.fa[i]])
			}
		}
	}
	return ans
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	result := process(reader)
	fmt.Println(result)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	// Read input
	s := readString(reader)
	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	rules := make([]Rule, n)
	for i := 0; i < n; i++ {
		var pattern string
		var l, r int
		fmt.Fscanf(reader, "%s %d %d\n", &pattern, &l, &r)
		rules[i] = Rule{pattern: pattern, l: l, r: r}
	}

	// Solve
	return solve(s, rules)
}
