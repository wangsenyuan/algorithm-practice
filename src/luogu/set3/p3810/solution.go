package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"os"
	"slices"
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
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([][3]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i][0], &a[i][1], &a[i][2])
	}
	return solve(k, a)
}

func solve(k int, a [][3]int) []int {
	freq := make(map[[3]int]int)
	for _, cur := range a {
		freq[cur]++
	}
	type item struct {
		a int
		b int
		c int
		f int
		r int
	}
	// n := len(freq)
	var arr []item
	for k, v := range freq {
		arr = append(arr, item{k[0], k[1], k[2], v, v - 1})
	}

	slices.SortFunc(arr, func(first item, second item) int {
		return cmp.Or(first.a-second.a, first.b-second.b, first.c-second.c)
	})

	t := make(BIT, k+3)

	var play func(l int, r int)

	play = func(l int, r int) {
		if l+1 == r {
			return
		}
		mid := (l + r) >> 1
		play(l, mid)
		play(mid, r)
		i := l
		for j := mid; j < r; j++ {
			for i < mid && arr[i].b <= arr[j].b {
				t.update(arr[i].c, arr[i].f)
				i++
			}
			arr[j].r += t.get(arr[j].c)
		}
		for i--; i >= l; i-- {
			t.update(arr[i].c, -arr[i].f)
		}
		type pair struct {
			it  item
			ans int
		}

		slices.SortFunc(arr[l:r], func(first item, second item) int {
			return cmp.Or(first.b-second.b, first.c-second.c)
		})
	}

	play(0, len(arr))

	g := make([]int, len(a))

	for _, v := range arr {
		g[v.r] += v.f
	}

	return g
}

type BIT []int

func (bit BIT) update(p int, v int) {
	p++
	for p < len(bit) {
		bit[p] += v
		p += p & -p
	}
}

func (bit BIT) get(p int) int {
	var res int
	p++
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}
