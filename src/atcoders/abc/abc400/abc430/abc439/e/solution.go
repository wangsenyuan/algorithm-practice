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
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	segs := make([][]int, n)
	for i := range n {
		segs[i] = make([]int, 2)
		fmt.Fscan(reader, &segs[i][0], &segs[i][1])
	}
	return solve(segs)
}

func solve(segs [][]int) int {
	type data struct {
		a int
		b int
	}

	var pos []int
	n := len(segs)
	arr := make([]data, n)
	for i, cur := range segs {
		arr[i] = data{cur[0], cur[1]}
		pos = append(pos, cur[1])
	}

	slices.SortFunc(arr, func(x data, y data) int {
		return cmp.Or(x.a-y.a, x.b-y.b)
	})

	slices.Sort(pos)

	pos = slices.Compact(pos)
	bit := make(BIT, len(pos)+3)
	dp := make([]int, len(pos))
	for i := 0; i < n; {
		i1 := i
		for i < n && arr[i].a == arr[i1].a {
			j := sort.SearchInts(pos, arr[i].b)
			dp[j] = max(dp[j], bit.get(j-1)+1)
			i++
		}

		for i1 < i {
			j := sort.SearchInts(pos, arr[i1].b)
			bit.update(j, dp[j])
			i1++
		}
	}

	return slices.Max(dp)
}

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] = max(bit[i], v)
		i += i & -i
	}
}

func (bit BIT) get(i int) int {
	i++
	var res int
	for i > 0 {
		res = max(res, bit[i])
		i -= i & -i
	}

	return res
}
