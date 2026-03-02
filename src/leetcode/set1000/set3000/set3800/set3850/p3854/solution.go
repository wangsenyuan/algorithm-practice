package p3854

import "slices"

func makeParityAlternating(nums []int) []int {
	if len(nums) == 1 {
		return []int{0, 0}
	}

	gMin := slices.Min(nums)
	gMax := slices.Max(nums)

	play := func(d int) []int {
		mn := inf
		mx := -inf
		var sum int
		for i, v := range nums {
			if (v-i)&1 != d {
				sum++
				switch v {
				case gMin:
					v++
				case gMax:
					v--
				}
			}
			mn = min(mn, v)
			mx = max(mx, v)
		}
		return []int{sum, max(mx-mn, 1)}
	}

	ans1 := play(0)
	ans2 := play(1)
	if ans1[0] < ans2[0] || ans1[0] == ans2[0] && ans1[1] < ans2[1] {
		return ans1
	}
	return ans2
}

const inf = 1 << 60
