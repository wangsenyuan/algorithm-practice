package p3655

import "math"

func xorAfterQueries(nums []int, queries [][]int) int {
	// i % k = l % k
	n := len(nums)
	block_size := int(math.Sqrt(float64(n)))

	small := make([][]int, block_size+1)
	for i, cur := range queries {
		l, r, k, v := cur[0], cur[1], cur[2], cur[3]
		if k <= block_size {
			small[k] = append(small[k], i)
		} else {
			for i := l; i <= r; i += k {
				nums[i] = mul(nums[i], v)
			}
		}
	}

	diff := make([]int, n+block_size+1)
	for i := range diff {
		diff[i] = 1
	}
	for k := 1; k <= block_size; k++ {
		ids := small[k]
		for j := range diff {
			diff[j] = 1
		}
		for _, j := range ids {
			l, r, v := queries[j][0], queries[j][1], queries[j][3]
			diff[l] = mul(diff[l], v)
			end := l + (r-l)/k*k + k
			if end < len(diff) {
				diff[end] = mul(diff[end], inverse(v))
			}
		}
		for i := range n {
			if i >= k {
				diff[i] = mul(diff[i], diff[i-k])
			}
			nums[i] = mul(nums[i], diff[i])
		}
	}
	var ans int
	for _, v := range nums {
		ans ^= v
	}
	return ans
}

const mod = 1000000007

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 > 0 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func inverse(num int) int {
	return pow(num, mod-2)
}
