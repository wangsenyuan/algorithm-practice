package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans, _ := process(reader)
	if len(ans) == 0 {
		fmt.Println(-1)
		return
	}

	fmt.Println(len(ans))
	s := fmt.Sprintf("%v", ans)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) ([]int, [][]int) {
	n := readNum(reader)
	cards := make([][]int, n)
	for i := range n {
		cards[i] = readNNums(reader, 4)
	}
	return solve(cards), cards
}

type pair struct{ b, i int }

var g [][]pair

type seg []struct{ l, r, min int }

func (t seg) maintain(o int) {
	t[o].min = min(t[o<<1].min, t[o<<1|1].min)
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].min = g[l][0].b
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) delete(o, qr, maxY int, f func(int)) {
	l := t[o].l
	if l > qr || t[o].min > maxY {
		return
	}
	if l == t[o].r {
		f(l)
		t[o].min = g[l][0].b
		return
	}
	t.delete(o<<1|1, qr, maxY, f)
	t.delete(o<<1, qr, maxY, f)
	t.maintain(o)
}

func solve(cards [][]int) []int {
	n := len(cards)
	a := make([]struct{ a, b, c, d int }, n+1)
	xs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = struct{ a, b, c, d int }{
			a: cards[i-1][0],
			b: cards[i-1][1],
			c: cards[i-1][2],
			d: cards[i-1][3],
		}
		xs[i] = a[i].a
	}
	slices.Sort(xs)
	xs = slices.Compact(xs)
	m := len(xs)

	g = make([][]pair, m)
	for i := 1; i <= n; i++ {
		x := sort.SearchInts(xs, a[i].a)
		g[x] = append(g[x], pair{a[i].b, i})
	}
	for i, ps := range g {
		slices.SortFunc(ps, func(a, b pair) int { return a.b - b.b })
		g[i] = append(ps, pair{2e9, 0}) // 哨兵
	}
	t := make(seg, 2<<bits.Len(uint(m-1)))
	t.build(1, 0, m-1)

	q := []int{0}
	pre := make([]int, n+1)
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		if i == n {
			ans := []int{}
			for ; i > 0; i = pre[i] {
				ans = append(ans, i)
			}
			slices.Reverse(ans)
			return ans
		}

		maxX := sort.SearchInts(xs, a[i].c+1) - 1
		maxY := a[i].d
		t.delete(1, maxX, maxY, func(l int) {
			ps := g[l]
			for ps[0].b <= maxY {
				pre[ps[0].i] = i
				q = append(q, ps[0].i)
				ps = ps[1:]
			}
			g[l] = ps
		})
	}
	return nil
}
