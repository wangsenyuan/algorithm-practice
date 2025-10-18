package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
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

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readString(reader *bufio.Reader) string {
	bs, _ := reader.ReadString('\n')
	return strings.TrimSpace(bs)
}

func drive(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	queries := make([]string, m)
	for i := 0; i < m; i++ {
		queries[i] = readString(reader)
	}
	return solve(n, queries)
}

func solve(n int, queries []string) []int {
	var rows []int
	var cols []int

	for _, cur := range queries {
		var r, c int
		pos := readInt([]byte(cur), 0, &c) + 1
		readInt([]byte(cur), pos, &r)
		rows = append(rows, r)
		cols = append(cols, c)
	}

	rows = append(rows, 0)
	cols = append(cols, 0)

	rows = sortAndCompact(rows)
	cols = sortAndCompact(cols)

	t1 := NewTree(len(rows))
	t2 := NewTree(len(cols))

	ans := make([]int, len(queries))

	for i, cur := range queries {
		var r, c int
		pos := readInt([]byte(cur), 0, &c) + 1
		pos = readInt([]byte(cur), pos, &r) + 1

		pr := sort.SearchInts(rows, r)
		pc := sort.SearchInts(cols, c)
		if cur[pos] == 'U' {
			w := t2.Get(pc)
			ans[i] = r - w
			j := sort.SearchInts(rows, w)
			t1.Update(j, pr, c)
			t2.Update(pc, pc, r)
		} else {
			w := t1.Get(pr)
			ans[i] = c - w
			j := sort.SearchInts(cols, w)
			t2.Update(j, pc, r)
			t1.Update(pr, pr, c)
		}
	}
	return ans
}

func sortAndCompact(arr []int) []int {
	slices.Sort(arr)
	return slices.Compact(arr)
}

type Tree struct {
	val  []int
	lazy []int
	sz   int
}

func NewTree(n int) *Tree {
	t := &Tree{
		val:  make([]int, 4*n),
		lazy: make([]int, 4*n),
		sz:   n,
	}
	return t
}

func (t *Tree) update(i int, v int) {
	t.val[i] = max(t.val[i], v)
	t.lazy[i] = max(t.lazy[i], v)
}

func (t *Tree) push(i int) {
	if t.lazy[i] != 0 {
		t.update(2*i+1, t.lazy[i])
		t.update(2*i+2, t.lazy[i])
		t.lazy[i] = 0
	}
}

func (t *Tree) Update(L int, R int, v int) {
	var dfs func(i int, l int, r int)
	dfs = func(i int, l int, r int) {
		if r < L || R < l {
			return
		}
		if L <= l && r <= R {
			t.update(i, v)
			return
		}
		t.push(i)
		mid := (l + r) >> 1
		dfs(2*i+1, l, mid)
		dfs(2*i+2, mid+1, r)
		t.val[i] = max(t.val[2*i+1], t.val[2*i+2])
	}
	dfs(0, 0, t.sz-1)
}

func (t *Tree) Get(p int) int {
	var dfs func(i int, l int, r int) int
	dfs = func(i int, l int, r int) int {
		if l == r {
			return t.val[i]
		}
		t.push(i)
		mid := (l + r) >> 1
		if p <= mid {
			return dfs(2*i+1, l, mid)
		}
		return dfs(2*i+2, mid+1, r)
	}
	return dfs(0, 0, t.sz-1)
}
