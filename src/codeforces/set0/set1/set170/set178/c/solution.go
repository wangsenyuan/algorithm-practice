package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	s := readString(reader)
	firstLine := readNums(s)
	h, m, n := firstLine[0], firstLine[1], firstLine[2]
	ops := make([]string, n)
	for i := range n {
		ops[i] = readString(reader)
	}

	return solve(h, m, ops)
}

func readNums(s string) []int {
	ss := strings.Split(s, " ")
	res := make([]int, len(ss))
	for i, x := range ss {
		res[i], _ = strconv.Atoi(x)
	}
	return res
}

func solve(h int, m int, ops []string) int {
	set := NewSegTree(h, inf, func(a, b int) int {
		return min(a, b)
	})
	// 感觉有可能出现圈
	marked := make([]int, h)

	for i := range h {
		marked[i] = -1
	}

	var cycles [][]int
	pos := make([]int, h)
	posInCycle := make([]int, h)

	for s, i := 0, 0; s < h; s++ {
		if marked[s] < 0 {
			u := s
			var cur []int
			for marked[u] < 0 {
				pos[u] = i
				set.Update(i, i)
				i++
				posInCycle[u] = len(cur)
				marked[u] = len(cycles)
				cur = append(cur, u)
				u = (u + m) % h
			}
			cycles = append(cycles, cur)
		}

	}

	var res int

	history := make(map[int]int)

	add := func(id int, hash int) {
		cycle := cycles[marked[hash]]
		free := set.Query(pos[hash], pos[hash]+len(cycle)-posInCycle[hash])
		// 它真实的位置
		if free != inf {
			res += free - pos[hash]
		} else {
			res += len(cycle) - posInCycle[hash]
			free = set.Query(pos[cycle[0]], pos[hash])
			res += free - pos[cycle[0]]
		}
		history[id] = free
		set.Update(free, inf)
	}

	rem := func(id int) {
		i := history[id]
		set.Update(i, i)
		delete(history, id)
	}

	for _, cmd := range ops {
		args := readNums(cmd[2:])
		if cmd[0] == '+' {
			add(args[0], args[1])
		} else {
			rem(args[0])
		}
	}

	return res
}

const inf = 1 << 60

type SegTree struct {
	val       []int
	initValue int
	f         func(int, int) int
}

func NewSegTree(n int, iv int, f func(int, int) int) *SegTree {

	arr := make([]int, 2*n)
	for i := range 2 * n {
		arr[i] = iv
	}

	return &SegTree{arr, iv, f}
}

func (tr *SegTree) Update(pos int, v int) {
	n := len(tr.val) / 2
	pos += n
	tr.val[pos] = v
	for pos > 1 {
		tr.val[pos>>1] = tr.f(tr.val[pos], tr.val[pos^1])
		pos >>= 1
	}
}

func (tr *SegTree) Query(l int, r int) int {
	n := len(tr.val) / 2
	l += n
	r += n
	res := tr.initValue
	for l < r {
		if l&1 == 1 {
			res = tr.f(res, tr.val[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = tr.f(res, tr.val[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
