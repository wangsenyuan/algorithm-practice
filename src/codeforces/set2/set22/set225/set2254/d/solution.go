package main

import (
	"bufio"
	"cmp"
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
		for _, v := range drive(reader) {
			fmt.Fprint(writer, v, " ")
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) []int64 {
	var n int
	fmt.Fscan(reader, &n)
	b := make([]int64, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(b)
}

func solve(b []int64) []int64 {
	n := len(b)
	if n == 1 {
		// b[0] = sum of v < a[0], 只能是0
		if b[0] == 0 {
			return []int64{1}
		}
		return []int64{-1}
	}

	type data struct {
		id  int
		val int64
	}

	arr := make([]data, len(b))
	for i, v := range b {
		arr[i] = data{i, v}
	}

	slices.SortFunc(arr, func(a data, b data) int {
		return cmp.Or(int(a.val-b.val), a.id-b.id)
	})

	slices.Sort(b)
	b = slices.Compact(b)
	if b[0] != 0 {
		return []int64{-1}
	}

	ans := make([]int64, n)
	var sum int64
	var r int
	for i, v := range b {
		l := r
		for r < n && arr[r].val == v {
			r++
		}

		// 这 r - l 个数的和 = diff
		var w int64
		if l > 0 {
			w = ans[arr[l-1].id]
		}
		w++

		if i+1 < len(b) {
			// 一共 r - l 个数
			diff := b[i+1] - sum
			if diff <= 0 || diff%int64(r-l) > 0 {
				return []int64{-1}
			}

			// 必须 >= w
			x := diff / int64(r-l)
			if x < w {
				return []int64{-1}
			}
			sum += x * int64(r-l)
			for l < r {
				ans[arr[l].id] = x
				l++
			}
		} else {
			for l < r {
				ans[arr[l].id] = w
				l++
			}
		}
	}

	return ans
}
