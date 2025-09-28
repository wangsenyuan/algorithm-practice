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

	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	segements := make([][]int, n)
	for i := 0; i < n; i++ {
		segements[i] = make([]int, 2)
		fmt.Fscan(reader, &segements[i][0], &segements[i][1])
	}
	return solve(segements)
}

func solve(segements [][]int) []int {
	var nums []int
	for _, cur := range segements {
		nums = append(nums, cur[0], cur[1])
	}
	sort.Ints(nums)
	nums = slices.Compact(nums)

	m := len(nums)

	active := make([]int, m)
	deactive := make([]int, m)

	for _, cur := range segements {
		l, r := cur[0], cur[1]
		i := sort.SearchInts(nums, l)
		active[i]++
		j := sort.SearchInts(nums, r)
		deactive[j]++
	}
	n := len(segements)
	ans := make([]int, n+1)

	var cnt int
	for i, x := range nums {
		if i > 0 {
			// x的要单独计算
			ans[cnt] += x - nums[i-1] - 1
		}
		cnt += active[i]
		ans[cnt]++
		cnt -= deactive[i]
	}

	return ans[1:]
}
