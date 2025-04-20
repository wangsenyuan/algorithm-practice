package p3523

import "slices"

func resultArray(nums []int, k int) []int64 {
	ans := make([]int64, k)
	dp := make([]int, k)
	fp := make([]int, k)
	for _, v := range nums {
		clear(fp)
		fp[v%k] = 1
		for y, c := range dp {
			fp[y*v%k] += c // 刷表法
		}
		for x, c := range fp {
			ans[x] += int64(c)
		}
		copy(dp, fp)
	}
	return ans
}

func resultArray1(nums []int, k int) []int64 {
	n := len(nums)
	val := make([]int, 4*n)
	cnt := make([][]int64, 4*n)

	merge := func(a []int64, b []int64, u int) []int64 {
		c := slices.Clone(a)
		for i, v := range b {
			c[(i*u)%k] += v
		}
		return c
	}

	var build func(i int, l int, r int)

	build = func(i int, l int, r int) {
		if l == r {
			val[i] = nums[l] % k
			cnt[i] = make([]int64, k)
			cnt[i][val[i]]++
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		val[i] = (val[2*i+1] * val[2*i+2]) % k
		cnt[i] = merge(cnt[2*i+1], cnt[2*i+2], val[2*i+1])
	}

	var query func(i int, l int, r int, L int, R int) ([]int64, int)

	query = func(i int, l int, r int, L int, R int) ([]int64, int) {
		if L <= l && r <= R {
			return cnt[i], val[i]
		}
		mid := (l + r) / 2
		if R <= mid {
			return query(2*i+1, l, mid, L, R)
		}
		if L > mid {
			return query(2*i+2, mid+1, r, L, R)
		}
		a, v1 := query(2*i+1, l, mid, L, R)
		b, v2 := query(2*i+2, mid+1, r, L, R)
		c := merge(a, b, v1)
		v := (v1 * v2) % k
		return c, v
	}

	build(0, 0, n-1)

	res := make([]int64, k)

	for i := range n {
		tmp, _ := query(0, 0, n-1, i, n-1)
		for j := range k {
			res[j] += tmp[j]
		}
	}

	return res
}
