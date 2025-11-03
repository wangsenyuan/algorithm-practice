package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	queries := make([][]int, n)
	for i := 0; i < n; i++ {
		queries[i] = make([]int, 3)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return solve(queries)
}

func solve(queries [][]int) []int {
	// 对于每个query的数，可以单独处理
	var xs []int
	for _, cur := range queries {
		xs = append(xs, cur[2])
	}
	slices.Sort(xs)
	xs = slices.Compact(xs)

	n := len(xs)

	pos := make([][]int, n)

	for _, cur := range queries {
		x := cur[2]
		j := sort.SearchInts(xs, x)
		pos[j] = append(pos[j], cur[1])
	}

	bits := make([]BIT, n)
	for i := range pos {
		sort.Ints(pos[i])
		bits[i] = make(BIT, len(pos[i])+1)
	}

	var res []int

	for _, cur := range queries {
		t, x := cur[1], cur[2]
		j := sort.SearchInts(xs, x)
		k := sort.SearchInts(pos[j], t)
		switch cur[0] {
		case 1:
			// add
			bits[j].update(k, 1)
		case 2:
			bits[j].update(k, -1)
		default:
			res = append(res, bits[j].get(k))
		}
	}

	return res
}

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) get(i int) int {
	i++
	var res int

	for i > 0 {
		res += bit[i]
		i -= i & -i
	}

	return res
}
