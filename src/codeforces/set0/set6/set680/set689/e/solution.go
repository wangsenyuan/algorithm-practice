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
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	segments := make([][]int, n)
	for i := 0; i < n; i++ {
		segments[i] = make([]int, 2)
		fmt.Fscan(reader, &segments[i][0], &segments[i][1])
	}
	return solve(k, segments)
}

const mod = 1e9 + 7

func add(a, b int) int {
	return (a + b) % mod
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func solve(k int, segments [][]int) int {

	m := len(segments)

	F := make([]int, m+1)
	F[0] = 1
	for i := 1; i <= m; i++ {
		F[i] = mul(F[i-1], i)
	}

	I := make([]int, m+1)
	I[m] = pow(F[m], mod-2)
	for i := m - 1; i >= 0; i-- {
		I[i] = mul(I[i+1], i+1)
	}

	C := func(a int, b int) int {
		if a < b {
			return 0
		}
		return mul(F[a], mul(I[b], I[a-b]))
	}

	// n := len(segments)
	var arr []int
	for _, cur := range segments {
		arr = append(arr, cur[0], cur[1], cur[1]+1)
	}

	slices.Sort(arr)
	arr = slices.Compact(arr)

	diff := make([]int, len(arr)+1)

	for _, cur := range segments {
		l, r := cur[0], cur[1]
		i := sort.SearchInts(arr, l)
		j := sort.SearchInts(arr, r)
		diff[i]++
		diff[j+1]--
	}

	var res int

	n := len(arr)

	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
		dist := arr[i] - arr[i-1]
		res = add(res, mul(dist, C(diff[i-1], k)))
	}

	return res
}
