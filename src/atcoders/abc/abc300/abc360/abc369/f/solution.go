package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, best, path := drive(reader)
	fmt.Println(best)
	fmt.Println(path)
}

func drive(reader *bufio.Reader) (h int, w int, coins [][]int, best int, path string) {
	var n int
	fmt.Fscan(reader, &h, &w, &n)
	coins = make([][]int, n)
	for i := range n {
		var r, c int
		fmt.Fscan(reader, &r, &c)
		coins[i] = []int{r, c}
	}
	best, path = solve(h, w, coins)
	return
}

type coin struct {
	r  int
	c  int
	id int
}

func solve(h int, w int, coins [][]int) (int, string) {
	n := len(coins)
	arr := make([]coin, n+2)
	arr[0] = coin{1, 1, 0}
	for i := range n {
		arr[i+1] = coin{coins[i][0], coins[i][1], i + 1}
	}
	arr[n+1] = coin{h, w, n + 1}

	slices.SortFunc(arr, func(a, b coin) int {
		return cmp.Or(a.r-b.r, a.c-b.c)
	})

	for i := range n + 2 {
		arr[i].id = i
	}
	stack := make([]coin, n+2)
	var top int
	stack[top] = arr[0]
	top++
	from := make([]int, n+2)
	for i := 1; i < n+2; i++ {
		cur := arr[i]
		j := sort.Search(top, func(j int) bool {
			return stack[j].c > cur.c
		})
		from[cur.id] = stack[j-1].id
		stack[j] = cur
		if j == top {
			top++
		}
	}

	// 头尾不包括
	best := top - 2

	var buf []byte

	cur := stack[top-1].id

	for cur > 0 {
		prev := from[cur]
		r1, c1 := arr[cur].r, arr[cur].c
		r2, c2 := arr[prev].r, arr[prev].c
		for r1 > r2 {
			buf = append(buf, 'D')
			r1--
		}
		for c1 > c2 {
			buf = append(buf, 'R')
			c1--
		}
		cur = prev
	}

	slices.Reverse(buf)
	return best, string(buf)
}

func solve1(h int, w int, coins [][]int) (int, string) {
	arr := make([][]coin, w)

	for i, cur := range coins {
		r, c := cur[0]-1, cur[1]-1
		arr[c] = append(arr[c], coin{r, c, i})
	}

	for c := range w {
		slices.SortFunc(arr[c], func(a, b coin) int {
			return a.r - b.r
		})
	}

	dp := make(SegTree, 2*h)

	for i := range 2 * h {
		dp[i] = pair{0, -1}
	}

	n := len(coins)

	from := make([]int, n)

	var best int
	var bestCol int
	var bestId int
	var bestRow int
	for c := range w {
		var todo [][]int
		fromTop := pair{0, -1}
		for _, cur := range arr[c] {
			r, i := cur.r, cur.id
			// 如果从左边过来
			v := dp.query(0, r+1)

			if v.first > fromTop.first {
				from[i] = v.second
			} else {
				from[i] = fromTop.second
			}

			todo = append(todo, []int{i, r, max(v.first+1, fromTop.first+1)})

			fromTop = max_pair(fromTop, v)
			fromTop.first++
			fromTop.second = i
		}

		for _, cur := range todo {
			i, r, v := cur[0], cur[1], cur[2]
			dp.update(r, v, i)
		}
		tmp := dp.query(0, h)
		if tmp.first > best {
			best = tmp.first
			bestCol = c
			bestId = tmp.second
			bestRow = coins[bestId][0] - 1
		}
	}

	var buf []byte

	row, col := h-1, w-1
	for row > bestRow {
		buf = append(buf, 'D')
		row--
	}
	for col > bestCol {
		buf = append(buf, 'R')
		col--
	}

	for bestId >= 0 {
		prevId := from[bestId]
		var prevRow, prevCol int
		if prevId >= 0 {
			prevRow = coins[prevId][0] - 1
			prevCol = coins[prevId][1] - 1
		}
		for bestRow > prevRow {
			buf = append(buf, 'D')
			bestRow--
		}
		for bestCol > prevCol {
			buf = append(buf, 'R')
			bestCol--
		}
		bestId = prevId
	}

	slices.Reverse(buf)
	return best, string(buf)
}

type pair struct {
	first  int
	second int
}

func max_pair(a, b pair) pair {
	if a.first > b.first {
		return a
	}
	if a.first == b.first && a.second > b.second {
		return a
	}
	return b
}

type SegTree []pair

func (st SegTree) update(i int, v int, w int) {
	n := len(st) / 2
	i += n
	if st[i].first > v {
		return
	}
	st[i] = pair{v, w}

	for i > 1 {
		st[i>>1] = max_pair(st[i], st[i^1])
		i >>= 1
	}
}

func (st SegTree) query(l int, r int) pair {
	n := len(st) / 2
	l += n
	r += n
	res := pair{0, -1}
	for l < r {
		if l&1 == 1 {
			res = max_pair(res, st[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max_pair(res, st[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
