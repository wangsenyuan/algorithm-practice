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
	_, _, res := drive(reader)
	fmt.Println(res[0], res[1])
}

func drive(reader *bufio.Reader) (k int, a []int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &k)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(a, k)
	return
}

func solve(a []int, k int) []int {
	nums := slices.Clone(a)
	slices.Sort(nums)
	nums = slices.Compact(nums)
	n := len(a)
	m := len(nums)
	freq := make([]int, m)

	var cnt int
	add := func(x int) {
		i := sort.SearchInts(nums, x)
		freq[i]++
		if freq[i] == 1 {
			cnt++
		}
	}
	rem := func(x int) {
		i := sort.SearchInts(nums, x)
		freq[i]--
		if freq[i] == 0 {
			cnt--
		}
	}

	res := []int{0, 0}

	for l, r := 0, 0; r < n; r++ {
		add(a[r])
		for l < r && cnt > k {
			rem(a[l])
			l++
		}
		if cnt <= k && r-l+1 > res[1]-res[0]+1 {
			res[0] = l
			res[1] = r
		}
	}
	res[0]++
	res[1]++
	return res
}
