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

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int64 {
	var n, m, x, y int
	fmt.Fscan(reader, &n, &m, &x, &y)
	a := make([]int, x)
	for i := range x {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, y)
	for i := range y {
		fmt.Fscan(reader, &b[i])
	}
	return solve(n, m, a, b)
}

func solve(n, m int, a, b []int) int64 {
	slices.Reverse(a)
	slices.Reverse(b)

	var c, d, e []int
	for i, j := 0, 0; i < len(a) || j < len(b); {
		if i < len(a) && j < len(b) && a[i] == b[j] {
			c = append(c, a[i])
			i++
			j++
		} else if j == len(b) || i < len(a) && a[i] > b[j] {
			d = append(d, a[i])
			i++
		} else {
			e = append(e, b[j])
			j++
		}
	}
	// 现在要找n+m-1个最大的元素
	if len(d) > n {
		d = d[:n]
	}
	if len(e) > m {
		e = e[:m]
	}
	arr := append(d, e...)
	slices.Sort(arr)
	slices.Reverse(arr)
	// 然后剩下的从c中获取
	sum := make([]int, len(c)+1)
	for i, v := range c {
		sum[i+1] = sum[i] + v
	}
	var best int
	var pref int
	for i := 0; i <= len(arr) && i <= n+m-1; i++ {
		// 需要 n + m - 1 - i 个c
		if n+m-1-i <= len(c) {
			best = max(best, pref+sum[n+m-1-i])
		} else {
			best = max(best, pref+sum[len(c)])
		}
		if i < len(arr) {
			pref += arr[i]
		}
	}

	return int64(best)
}
