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
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) int {
	arr := slices.Clone(a)
	slices.Sort(arr)
	arr = slices.Compact(arr)

	nums := slices.Clone(a)
	for i, v := range nums {
		nums[i] = sort.SearchInts(arr, v)
	}

	n := len(arr)
	bit := make(BIT, n+2)

	m := len(nums)
	dp := make([]int, m)
	var tot int
	for i, v := range nums {
		dp[i] = bit.queryRange(v+1, n)
		tot += dp[i]
		bit.update(v, 1)
	}

	if tot <= k {
		return len(nums) * (len(nums) - 1) / 2
	}

	bit2 := make(BIT, n+2)

	var res int

	var l int
	// 要删除的区间包r
	sum := tot
	for r := range len(nums) - 1 {
		// 将r移进删除区间
		bit.update(nums[r], -1)
		if nums[r] > 0 {
			sum -= bit.queryRange(0, nums[r]-1)
		}

		sum -= bit2.queryRange(nums[r]+1, n)

		for l < r && sum <= k {
			// 将l再移出删除区间
			sum += dp[l]
			if nums[l] > 0 {
				sum += bit.queryRange(0, nums[l]-1)
			}
			if sum > k {
				// 删多了
				sum -= dp[l]
				if nums[l] > 0 {
					sum -= bit.queryRange(0, nums[l]-1)
				}
				break
			}
			bit2.update(nums[l], 1)
			l++
		}
		if sum <= k {
			res += l
		}
	}

	return res
}

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) query(i int) int {
	i++
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) queryRange(l, r int) int {
	return bit.query(r) - bit.query(l-1)
}
