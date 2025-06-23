package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) []int {
	n, q := readTwoNums(reader)
	p := readNNums(reader, n)
	qs := make([][]int, q)
	for i := range q {
		qs[i] = readNNums(reader, 4)
	}
	return solve(n, p, qs)
}
func solve(n int, p []int, qs [][]int) []int {
	// n * (n - 1) / 2 - 上下左右 + 四个角落

	trs := make([]*Tree, n+1)

	for i := 1; i <= n; i++ {
		trs[i] = trs[i-1].Add(p[i-1], 1, n)
	}

	ans := make([]int, len(qs))

	count := func(m int) int {
		return m * (m - 1) / 2
	}

	find := func(l int, d int, r int, u int) int {
		ans := count(n)

		ans -= count(l - 1)
		ans -= count(n - r)
		if d > 1 {
			w := trs[n].Query(d-1, 1, n)
			ans -= count(w)
		}
		if u < n {
			w := n - trs[n].Query(u, 1, n)
			ans -= count(w)
		}

		// 左下角
		if l > 1 {
			if d > 1 {
				w := trs[l-1].Query(d-1, 1, n)
				ans += count(w)
			}
			if u < n {
				w := trs[l-1].Query(n, 1, n) - trs[l-1].Query(u, 1, n)
				ans += count(w)
			}
		}
		if r < n {
			if d > 1 {
				w := trs[n].Query(d-1, 1, n) - trs[r].Query(d-1, 1, n)
				ans += count(w)
			}
			if u < n {
				w := trs[n].Query(n, 1, n) - trs[n].Query(u, 1, n)
				w -= trs[r].Query(n, 1, n) - trs[r].Query(u, 1, n)
				ans += count(w)
			}
		}

		return ans
	}

	for i, cur := range qs {
		ans[i] = find(cur[0], cur[1], cur[2], cur[3])
	}

	return ans
}

type Tree struct {
	children []*Tree
	cnt      int
}

func (tr *Tree) copy() *Tree {
	res := new(Tree)
	res.children = make([]*Tree, 2)

	if tr != nil {
		copy(res.children, tr.children)
		res.cnt = tr.cnt
	}
	return res
}

func (tr *Tree) Count() int {
	if tr == nil {
		return 0
	}
	return tr.cnt
}

func (tr *Tree) Add(x int, l int, r int) *Tree {
	res := tr.copy()
	if l == r {
		res.cnt++
		return res
	}
	mid := (l + r) >> 1
	if x <= mid {
		res.children[0] = res.children[0].Add(x, l, mid)
	} else {
		res.children[1] = res.children[1].Add(x, mid+1, r)
	}
	res.cnt = res.children[0].Count() + res.children[1].Count()
	return res
}

func (tr *Tree) Query(x int, l int, r int) int {
	if tr == nil {
		return 0
	}
	if l == r {
		return tr.cnt
	}
	mid := (l + r) >> 1
	if x <= mid {
		return tr.children[0].Query(x, l, mid)
	}
	res := tr.children[0].Count()
	return res + tr.children[1].Query(x, mid+1, r)
}
