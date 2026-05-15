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
	var n int
	var l, r int64
	fmt.Fscan(reader, &n, &l, &r)
	a := make([]int64, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(l, r, a)
}

func solve(l int64, r int64, a []int64) int64 {
	var arr []int64
	var res int64

	for _, v := range a {
		if v <= l {
			arr = append(arr, l)
			res += l - v
		} else if v >= r {
			arr = append(arr, r)
			res += v - r
		} else {
			arr = append(arr, v)
		}
	}

	slices.Sort(arr)
	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		res += arr[r] - arr[l]
	}
	return res
}
