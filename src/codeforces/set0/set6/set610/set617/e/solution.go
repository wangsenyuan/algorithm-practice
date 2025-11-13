package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(k, a, queries)
}

type query struct {
	id int
	l  int
	r  int
}

func solve(k int, a []int, queries [][]int) []int {
	var arr []int
	arr = append(arr, 0)
	var sum int
	for _, v := range a {
		sum ^= v
		arr = append(arr, sum)
	}

	slices.Sort(arr)
	arr = slices.Compact(arr)
	m := len(arr)
	n := len(a)
	pos := make([][2]int, n+1)
	history := make([][]int, m)
	history[0] = append(history[0], 0)
	sum = 0
	for i, v := range a {
		sum ^= v
		j := sort.SearchInts(arr, sum)
		// pos[i][0] = j
		i++
		history[j] = append(history[j], i)
		j = sort.SearchInts(arr, sum^k)
		if j < m && arr[j] == sum^k {
			pos[i][1] = j
		} else {
			pos[i][1] = -1
		}
	}

	pk := sort.SearchInts(arr, k)
	if pk < m && arr[pk] == k {
		pos[0][1] = pk
	} else {
		pos[0][1] = -1
	}

	sz := make([]int, m)
	var cur int
	open := make([]int, m)
	for i := range m {
		open[i] = cur
		for _, j := range history[i] {
			pos[j][0] = cur
			sz[i]++
			cur++
		}
	}

	qs := make([]query, len(queries))
	for i, cur := range queries {
		l, r := cur[0], cur[1]
		qs[i] = query{id: i, l: l - 1, r: r}
	}

	block_size := int(math.Sqrt(float64(n)))

	slices.SortFunc(qs, func(a, b query) int {
		if a.r/block_size != b.r/block_size {
			return a.r - b.r
		}
		s := a.r / block_size
		if s&1 == 0 {
			return a.l - b.l
		}
		return b.l - a.l
	})

	var tot int
	bit := make(BIT, n+3)

	add := func(i int) {
		// 如果 k = 0,
		j := pos[i][1]
		if j >= 0 {
			tmp := bit.query(open[j], open[j]+sz[j]-1)
			tot += tmp
		}
		bit.update(pos[i][0], 1)
	}

	rem := func(i int) {
		bit.update(pos[i][0], -1)
		j := pos[i][1]
		if j >= 0 {
			tmp := bit.query(open[j], open[j]+sz[j]-1)
			tot -= tmp
		}
	}

	ans := make([]int, len(queries))

	var l, r int
	for _, cur := range qs {
		for r <= cur.r {
			add(r)
			r++
		}
		for r-1 > cur.r {
			r--
			rem(r)
		}
		for l > cur.l {
			l--
			add(l)
		}
		for l < cur.l {
			rem(l)
			l++
		}
		ans[cur.id] = tot
	}

	return ans
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

func (bit BIT) query(l int, r int) int {
	return bit.get(r) - bit.get(l-1)
}
