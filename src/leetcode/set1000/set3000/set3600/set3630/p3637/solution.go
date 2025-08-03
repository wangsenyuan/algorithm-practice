package p3637

func isTrionic(nums []int) bool {
	var p int
	n := len(nums)
	for p+1 < n && nums[p] < nums[p+1] {
		p++
	}
	if p == 0 || p == n-1 {
		return false
	}
	q := p
	for q+1 < n && nums[q] > nums[q+1] {
		q++
	}

	if p == q || q == n-1 {
		return false
	}
	r := q
	for r+1 < n && nums[r] < nums[r+1] {
		r++
	}

	return r == n-1
}
