package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Println(drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	ss := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &ss[i])
	}
	return solve(ss)
}

func solve(ss []string) string {
	n := len(ss)
	zs := make([][]int, n)
	for i, s := range ss {
		zs[i] = zFunction(s)
	}

	var check func(i int, j int) int
	check = func(i int, j int) int {
		if len(ss[i]) < len(ss[j]) {
			return -check(j, i)
		}
		x, y := ss[i], ss[j]
		a, b := len(x), len(y)
		for k := range b {
			if x[k] < y[k] {
				return -1
			}
			if x[k] > y[k] {
				return 1
			}
		}
		if a > b && zs[i][b] < a-b {
			p := zs[i][b]
			if x[b+p] < x[p] {
				return -1
			}
			return 1
		}
		for k := range b {
			if y[k] < x[a-b+k] {
				return -1
			}
			if y[k] > x[a-b+k] {
				return 1
			}
		}
		return 0
	}

	arr := make([]int, n)
	for i := range n {
		arr[i] = i
	}

	slices.SortFunc(arr, check)

	if n == 2 {
		arr[0], arr[1] = arr[1], arr[0]
		return playOrderWithStrings(ss, arr)
	}

	for i := range n - 1 {
		if check(arr[i], arr[i+1]) == 0 {
			return playOrderWithStrings(ss, arr)
		}
	}

	a := slices.Clone(arr)
	a[n-2], a[n-1] = a[n-1], a[n-2]
	b := slices.Clone(arr)
	b[n-3], b[n-2] = b[n-2], b[n-3]
	x := playOrderWithStrings(ss, a)
	y := playOrderWithStrings(ss, b)
	return min(x, y)
}

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	var l, r int
	for i := 1; i < n; i++ {
		if i < r {
			z[i] = min(r-i, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i] > r {
			l, r = i, i+z[i]
		}
	}
	return z
}

func playOrderWithStrings(ss []string, arr []int) string {
	var buf bytes.Buffer
	for _, i := range arr {
		buf.WriteString(ss[i])
	}
	return buf.String()
}
