package p3640

func maxSumTrionic(nums []int) int64 {
	// l...r 似乎是无关的
	n := len(nums)
	sum := make([]int, n+1)
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}
	L := make([]int, n)
	Z := make([]int, n)
	for i := 0; i < n; i++ {
		L[i] = i
		if i > 0 && nums[i-1] < nums[i] {
			L[i] = L[i-1]
		}
		Z[i] = i
		if i > 0 && nums[i-1] > nums[i] {
			Z[i] = Z[i-1]
		}
	}

	R := make([]int, n+1)
	R[n] = n
	for i := n - 1; i >= 0; i-- {
		R[i] = i
		if i < n-1 && nums[i] < nums[i+1] {
			R[i] = R[i+1]
		}
	}

	var best = -inf
	for q := 0; q < n; q++ {
		p := Z[q]
		r := R[q]
		l := L[p]
		if r == q || p == q || p == l {
			continue
		}

		if sum[r+1] < sum[q+2] {
			r = q + 1
		}

		if nums[p] > 0 {
			l = search(l, p, func(i int) bool {
				return nums[i] >= 0
			})
			if l == p {
				l--
			}
		} else {
			l = p - 1
		}
		best = max(best, sum[r+1]-sum[l])
	}

	return int64(best)
}

const inf = 1 << 60

func search(l int, r int, f func(int) bool) int {
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
