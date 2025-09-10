package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const H = 30

func solve(a []int) int {

	var dfs func(arr []int, d int) int

	dfs = func(arr []int, d int) int {
		if d < 0 || len(arr) <= 1 {
			return 0
		}
		var hi []int
		var lo []int

		for _, num := range arr {
			if (num>>d)&1 == 0 {
				lo = append(lo, num)
			} else {
				hi = append(hi, num)
			}
		}
		if len(hi) == 0 {
			return dfs(lo, d-1)
		}

		if len(lo) == 0 {
			return dfs(hi, d-1)
		}

		res := dfs(lo, d-1) + dfs(hi, d-1)
		// 要把lo和hi种的连接起来

		if len(lo) > len(hi) {
			lo, hi = hi, lo
		}

		tr := NewTrie()
		for _, num := range lo {
			tr.Add(num)
		}

		best := 1 << (d + 1)

		for _, num := range hi {
			best = min(best, tr.GetXorMin(num))
		}

		return res + best
	}

	return dfs(a, H-1)
}

type Trie struct {
	next [][2]int
}

func NewTrie() *Trie {
	return &Trie{next: make([][2]int, 1)}
}

func (t *Trie) Add(num int) {
	var cur int
	for i := H - 1; i >= 0; i-- {
		x := (num >> i) & 1
		if t.next[cur][x] == 0 {
			t.next = append(t.next, [2]int{0, 0})
			t.next[cur][x] = len(t.next) - 1
		}
		cur = t.next[cur][x]
	}
}

func (t *Trie) GetXorMin(num int) int {
	var cur int
	var res int
	for i := H - 1; i >= 0; i-- {
		x := (num >> i) & 1
		if t.next[cur][x] != 0 {
			cur = t.next[cur][x]
		} else {
			res |= 1 << i
			cur = t.next[cur][1^x]
		}
	}
	return res
}
