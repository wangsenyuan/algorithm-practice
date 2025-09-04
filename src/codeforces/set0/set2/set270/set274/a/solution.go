package main

import (
	"bufio"
	"fmt"
	"os"
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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k)
}

func solve(a []int, k int) int {
	sort.Ints(a)
	var nums []int
	var cnt []int
	n := len(a)
	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		nums = append(nums, a[j])
		cnt = append(cnt, i-j)
	}
	if k == 1 {
		return len(nums)
	}
	m := len(nums)

	marked := make([]bool, m)

	var res int

	for i := range m {
		if marked[i] {
			continue
		}
		var sum [2]int
		j := i
		var d int
		for j < m {
			marked[j] = true
			sum[d] += cnt[j]
			d ^= 1
			j = sort.SearchInts(nums, k*nums[j])
			if j == m || nums[j] > k*nums[i] {
				break
			}
		}
		res += max(sum[0], sum[1])
	}

	return res
}
