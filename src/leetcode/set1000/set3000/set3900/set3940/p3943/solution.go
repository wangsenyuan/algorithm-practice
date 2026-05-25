package p3943

import "math"

func numberOfPairs(nums1 []int, nums2 []int, queries [][]int) []int {
	n := len(nums2)
	blockSize := int(math.Sqrt(float64(n)))
	m := (n + blockSize - 1) / blockSize

	blocks := make([]map[int]int, m)
	add := make([]int, m)
	for i := range m {
		blocks[i] = make(map[int]int)
	}

	for i, v := range nums2 {
		blocks[i/blockSize][v]++
	}

	update := func(x int, y int, val int) {
		l := x / blockSize
		r := y / blockSize
		if l == r {
			for i := x; i <= y; i++ {
				blocks[i/blockSize][nums2[i]]--
				nums2[i] += val
				blocks[i/blockSize][nums2[i]]++
			}
			return
		}
		for i := l + 1; i < r; i++ {
			add[i] += val
		}
		for i := x; i < (l+1)*blockSize; i++ {
			blocks[i/blockSize][nums2[i]]--
			nums2[i] += val
			blocks[i/blockSize][nums2[i]]++
		}

		for i := r * blockSize; i <= y; i++ {
			blocks[i/blockSize][nums2[i]]--
			nums2[i] += val
			blocks[i/blockSize][nums2[i]]++
		}
	}

	get := func(y int) int {
		var ans int
		for i := range m {
			x := y - add[i]
			ans += blocks[i][x]
		}
		return ans
	}

	query := func(tot int) int {
		var ans int
		for _, x := range nums1 {
			ans += get(tot - x)
		}
		return ans
	}

	var res []int
	for _, cur := range queries {
		if cur[0] == 1 {
			update(cur[1], cur[2], cur[3])
		} else {
			res = append(res, query(cur[1]))
		}
	}

	return res
}
