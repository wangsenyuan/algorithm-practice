package p3878

const H = 30

func countGoodSubarrays(nums []int) int64 {
	n := len(nums)

	last := make([]int, H)
	for i := range last {
		last[i] = -1
	}

	L := make([]int, n)
	for i := range n {
		pos := -1
		for d := range H {
			if (nums[i]>>d)&1 == 0 {
				pos = max(pos, last[d])
			}
		}
		L[i] = pos + 1
		for d := range H {
			if (nums[i]>>d)&1 == 1 {
				last[d] = i
			}
		}
	}

	next := make([]int, H)
	for i := range H {
		next[i] = n
	}

	book := make(map[int]int)

	var res int64
	for i := n - 1; i >= 0; i-- {
		pos := n
		for d := range H {
			if (nums[i]>>d)&1 == 0 {
				pos = min(pos, next[d])
			}
		}
		if j, ok := book[nums[i]]; ok {
			pos = min(pos, j)
		}
		r := pos - 1

		res += int64(r-i+1) * int64(i-L[i]+1)

		for d := range H {
			if (nums[i]>>d)&1 == 1 {
				next[d] = i
			}
		}
		book[nums[i]] = i
	}

	return res
}
