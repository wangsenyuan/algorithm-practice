package p3793

import "slices"

func minLength(nums []int, k int) int {
	x := slices.Max(nums)
	freq := make([]int, x+1)

	var sum int

	add := func(v int) {
		freq[v]++
		if freq[v] == 1 {
			sum += v
		}
	}

	rem := func(v int) {
		freq[v]--
		if freq[v] == 0 {
			sum -= v
		}
	}

	ans := -1

	update := func(l int, r int) {
		if ans == -1 || (r-l+1) < ans {
			ans = r - l + 1
		}
	}

	var l int
	for r, v := range nums {
		add(v)
		for sum >= k {
			update(l, r)
			rem(nums[l])
			if sum < k {
				add(nums[l])
				break
			}
			l++
		}
	}
	return ans
}
