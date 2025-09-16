package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
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

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	queries := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &queries[i])
	}
	return solve(a, queries)
}

type data struct {
	g   int
	cnt int
}

func solve(a []int, queries []int) []int {
	n := len(a)

	var arr []data
	// mem[r][x] = l 表示

	tr := NewTree(n)

	for i := range n {
		tr.Set(i, a[i])
		r := i
		g := a[i]
		for {
			l := tr.FindLeftMostPos(g)
			if gcd(g, a[l]) != g {
				l++
			}
			// gcd a[l....i] = g
			arr = append(arr, data{g, r - l + 1})
			l--
			if l < 0 {
				break
			}
			g = gcd(g, a[l])
			r = l
		}
	}

	slices.SortFunc(arr, func(x, y data) int {
		return x.g - y.g
	})

	var it int
	for i := 0; i < len(arr); {
		var cnt int
		j := i
		for i < len(arr) && arr[i].g == arr[j].g {
			cnt += arr[i].cnt
			i++
		}
		arr[it] = data{arr[j].g, cnt}
		it++
	}
	arr = arr[:it]

	ans := make([]int, len(queries))

	for i, x := range queries {
		j := sort.Search(it, func(d int) bool {
			return arr[d].g >= x
		})

		if j < it && arr[j].g == x {
			ans[i] = arr[j].cnt
		}
	}
	return ans
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type Tree []int

func NewTree(n int) Tree {
	tr := make(Tree, 4*n)
	return tr
}

func (tr Tree) Set(pos int, v int) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			tr[i] = v
			return
		}
		mid := (l + r) >> 1
		if pos <= mid {
			f(i*2+1, l, mid)
		} else {
			f(i*2+2, mid+1, r)
		}
		tr[i] = gcd(tr[2*i+1], tr[2*i+2])
	}
	f(0, 0, len(tr)/4-1)
}

func (tr Tree) FindLeftMostPos(g int) int {
	var f func(i int, l int, r int) int
	f = func(i int, l int, r int) int {
		if l == r {
			return l
		}
		mid := (l + r) >> 1
		if tr[2*i+2]%g == 0 {
			return f(i*2+1, l, mid)
		}
		return f(i*2+2, mid+1, r)
	}
	return f(0, 0, len(tr)/4-1)
}
