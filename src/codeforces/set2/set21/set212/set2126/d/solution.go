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
	var tc int
	fmt.Fscan(reader, &tc)
	var buf bytes.Buffer
	for range tc {
		buf.WriteString(fmt.Sprintf("%d\n", process(reader)))
	}
	fmt.Print(buf.String())
}

func process(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	casinos := make([][]int, n)
	for i := range n {
		casinos[i] = make([]int, 3)
		fmt.Fscan(reader, &casinos[i][0], &casinos[i][1], &casinos[i][2])
	}
	return solve(k, casinos)
}

type pair struct {
	first  int
	second int
}

const inf = 1 << 60

func solve(k int, casinos [][]int) int {
	slices.SortFunc(casinos, func(a, b []int) int {
		return a[2] - b[2]
	})

	res := k
	for _, casino := range casinos {
		if casino[0] <= res && res <= casino[2] {
			res = casino[2]
		}
	}
	return res
}

func solve1(k int, casinos [][]int) int {
	var L []int
	for _, casino := range casinos {
		L = append(L, casino[0])
	}
	sort.Ints(L)
	L = slices.Compact(L)

	arr := make([][]pair, len(L))

	for _, cur := range casinos {
		l, r, real := cur[0], cur[1], cur[2]
		i := sort.SearchInts(L, l)
		arr[i] = append(arr[i], pair{r, real})
	}
	n := len(L)
	for i := range n {
		arr[i] = append(arr[i], pair{-1, 0})
		slices.SortFunc(arr[i], func(a, b pair) int {
			return b.first - a.first
		})
	}

	tr := make([]int, 4*n)

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		if l == r {
			tr[i] = arr[l][0].first
			return
		}
		mid := (l + r) / 2
		build(i*2+1, l, mid)
		build(i*2+2, mid+1, r)
		tr[i] = max(tr[i*2+1], tr[i*2+2])
	}

	build(0, 0, n-1)

	var update func(i int, l int, r int, x int, f func(int))

	update = func(i int, l int, r int, x int, f func(int)) {
		if tr[i] < x || x < L[l] {
			return
		}
		// l <= x && x <= tr[i]
		if l == r {
			f(l)
			tr[i] = arr[l][0].first
			return
		}
		mid := (l + r) / 2
		update(i*2+1, l, mid, x, f)
		update(i*2+2, mid+1, r, x, f)
		tr[i] = max(tr[i*2+1], tr[i*2+2])
	}

	var que []int
	que = append(que, k)

	var res int

	for len(que) > 0 {
		x := que[0]
		res = max(res, x)
		que = que[1:]

		update(0, 0, n-1, x, func(l int) {
			for len(arr[l]) > 0 && x <= arr[l][0].first {
				que = append(que, arr[l][0].second)
				arr[l] = arr[l][1:]
			}
		})
	}

	return res
}
