package p3681

const H = 30

func maxXorSubsequences(nums []int) int {
	bases := make([]int, H)

	add := func(x int) {
		for i := H - 1; i >= 0; i-- {
			if (x >> i) & 1 == 0 {
				continue
			}
			if bases[i] == 0 {
				bases[i] = x
				return
			}
			x ^= bases[i]
		}
	}

	for _, num := range nums {
		add(num)
	}

	var res int
	for i := H - 1; i >= 0; i-- {
		res = max(res, res^bases[i])
	}

	return res
}
