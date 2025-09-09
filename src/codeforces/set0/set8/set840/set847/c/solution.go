package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	res := solve(n, k)
	fmt.Println(res)
}

func solve(n int, k int) string {
	sum := (n - 1) * n / 2
	if sum < k {
		return "Impossible"
	}

	if k == 0 {
		return strings.Repeat("()", n)
	}
	i := sort.Search(n+1, func(i int) bool {
		return i*(i-1)/2 > k
	})
	i--
	k -= i * (i - 1) / 2

	nums := []int{i - 1}

	for k > 0 {
		last := nums[len(nums)-1]
		w := min(last, k)
		k -= w
		nums = append(nums, w)
	}

	var buf bytes.Buffer

	first := nums[0]
	for range first {
		buf.WriteByte('(')
	}

	buf.WriteByte('(')
	buf.WriteByte(')')

	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		diff := first - cur
		for range diff {
			buf.WriteByte(')')
		}
		buf.WriteByte('(')
		buf.WriteByte(')')
		first = cur
	}
	for range first {
		buf.WriteByte(')')
	}

	for buf.Len() < 2*n {
		buf.WriteByte('(')
		buf.WriteByte(')')
	}

	return buf.String()
}
