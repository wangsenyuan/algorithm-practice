package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, x := range drive(reader) {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, q)
	for i := range q {
		var t int
		fmt.Fscan(reader, &t)
		if t == 3 {
			queries[i] = []int{t}
			continue
		}
		var x int
		fmt.Fscan(reader, &x)
		queries[i] = []int{t, x}
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	nums := slices.Clone(a)
	for _, cur := range queries {
		if len(cur) == 2 {
			nums = append(nums, cur[1])
		}
	}
	slices.Sort(nums)
	nums = slices.Compact(nums)

	cnt := make(fenwick, len(nums)+2)
	sum := make(fenwick, len(nums)+2)

	slices.Sort(a)
	var tot int
	for j, x := range a {
		i, _ := slices.BinarySearch(nums, x)
		cnt.add(i, 1)
		sum.add(i, x)
		tot = add(tot, mul(j+1, x))
	}

	var res []int

	for _, cur := range queries {
		if cur[0] == 3 {
			res = append(res, tot)
			continue
		}
		x := cur[1]
		j, _ := slices.BinarySearch(nums, x)
		i := cnt.query(0, j)
		// i就是左边有多少数  <= x
		tmp := sum.query(j+1, len(nums)-1)
		switch cur[0] {
		case 1:
			tot = add(tot, mul(i+1, x))
			tot = add(tot, tmp)

			cnt.add(j, 1)
			sum.add(j, x)
		case 2:
			tot = sub(tot, mul(i, x))
			tot = sub(tot, tmp)

			cnt.add(j, -1)
			sum.add(j, sub(mod, x))
		}
	}

	return res
}

const mod = 998244353

func add(a int, b int) int {
	return (a + b) % mod
}

func sub(a int, b int) int {
	res := (a - b + mod) % mod
	if res < 0 {
		res += mod
	}
	return res
}

func mul(a int, b int) int {
	return a * b % mod
}

type fenwick []int

func (f fenwick) add(i int, x int) {
	i++
	for i < len(f) {
		f[i] = add(f[i], x)
		i += i & -i
	}
}

func (f fenwick) sum(i int) int {
	i++
	res := 0
	for i > 0 {
		res = add(res, f[i])
		i -= i & -i
	}
	return res
}

func (f fenwick) query(l int, r int) int {
	return sub(f.sum(r), f.sum(l-1))
}
