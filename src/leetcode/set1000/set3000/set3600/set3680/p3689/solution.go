package p3689

import (
	"slices"
	"sort"
)

type pair struct {
	first  int
	second int
}

func encode(nums, sorted []int) (res int) {
	for i, x := range nums {
		res |= sort.SearchInts(sorted, x) << (i * 3)
	}
	return
}

func minSplitMerge(nums1 []int, nums2 []int) int {
	if slices.Equal(nums1, nums2) {
		return 0
	}

	n := len(nums1)
	sorted := slices.Clone(nums1) // 用于离散化
	slices.Sort(sorted)

	val1 := encode(nums1, sorted)
	val2 := encode(nums2, sorted)

	vis := map[int]bool{val1: true}
	q := []int{val1}
	for ans := 1; ; ans++ {
		tmp := q
		q = nil
		for _, a := range tmp {
			for r := 1; r <= n; r++ { // 为方便实现，先枚举 r，再枚举 l
				t := a & (1<<(r*3) - 1)
				for l := range r {
					sub := t >> (l * 3)
					b := a&(1<<(l*3)-1) | a>>(r*3)<<(l*3) // 从 a 中移除 sub
					for i := range n - r + l + 1 {
						c := b&(1<<(i*3)-1) | sub<<(i*3) | b>>(i*3)<<((i+r-l)*3)
						if c == val2 {
							return ans
						}
						if !vis[c] {
							vis[c] = true
							q = append(q, c)
						}
					}
				}
			}
		}
	}
}

const inf = 1 << 30
