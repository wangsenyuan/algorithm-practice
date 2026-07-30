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

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a, b []int) int {
	n := len(a)
	if n == 1 {
		return min(a[0], b[0])
	}

	check := func(w int) bool {
		var arr []int

		for i := range n {
			var cnt int
			if a[i] >= w {
				cnt++
			}
			if b[i] >= w {
				cnt++
			}
			if cnt != 1 {
				arr = append(arr, cnt)
			}
		}

		var cnt [2]int
		for i := 0; i < len(arr); {
			j := i
			for i < len(arr) && arr[i] == arr[j] {
				i++
			}
			if arr[j] == 0 {
				cnt[0]++
			} else {
				cnt[1] += i - j
			}
		}

		return cnt[0] < cnt[1]
	}

	nums := slices.Clone(a)
	nums = append(nums, b...)
	slices.Sort(nums)
	nums = slices.Compact(nums)

	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if check(nums[mid]) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l-1]
}
