package p3911

import "sort"

func kthRemainingInteger(nums []int, queries [][]int) []int {
	n := len(nums)
	pref := make([]int, n+1)
	val := make([]int, n+1)

	for i, v := range nums {
		val[i+1] = val[i]
		pref[i+1] = pref[i]
		if v%2 == 0 {
			val[i+1] = v
			pref[i+1]++
		}
	}
	find := func(l int, r int, k int) int {
		if pref[r+1] == pref[l] {
			// l...r没有偶数
			return 2 * k
		}

		d := sort.Search(r-l+1, func(d int) bool {
			i := l + d
			// 到i出现了cnt个偶数
			// 一共有 val[i+1] / 2 个偶数
			cnt := pref[i+1] - pref[l]
			return val[i+1]/2-cnt >= k
		})
		i := l + d - 1
		cnt := val[i+1]/2 - (pref[i+1] - pref[l])
		return val[i+1] + 2*(k-cnt)
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = find(cur[0], cur[1], cur[2])
	}

	return ans
}
