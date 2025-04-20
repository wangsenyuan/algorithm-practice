package p3524

func resultArray(nums []int, k int, queries [][]int) []int {
	// k > 1 and k <= 5
	n := len(nums)

	val := make([]int, 4*n)
	arr := make([][]int, 4*n)

	merge := func(res []int, left []int, right []int, lp int) {
		copy(res, left)
		for v := range k {
			res[(v*lp)%k] += right[v]
		}
	}

	var build func(i int, l int, r int)
	build = func(i int, l int, r int) {
		arr[i] = make([]int, k)
		if l == r {
			val[i] = nums[l] % k
			arr[i][val[i]]++
			return
		}
		mid := (l + r) / 2
		build(2*i+1, l, mid)
		build(2*i+2, mid+1, r)
		merge(arr[i], arr[2*i+1], arr[2*i+2], val[2*i+1])
		val[i] = (val[2*i+1] * val[2*i+2]) % k
	}

	build(0, 0, n-1)

	var update func(i int, l int, r int, pos int, v int)
	update = func(i int, l int, r int, pos int, v int) {
		if l == r {
			val[i] = v % k
			clear(arr[i])
			arr[i][v%k]++
			return
		}
		mid := (l + r) / 2
		if pos <= mid {
			update(2*i+1, l, mid, pos, v)
		} else {
			update(2*i+2, mid+1, r, pos, v)
		}
		merge(arr[i], arr[2*i+1], arr[2*i+2], val[2*i+1])
		val[i] = (val[2*i+1] * val[2*i+2]) % k
	}

	var query func(i int, l int, r int, L int, R int) (cnt []int, v int)
	query = func(i int, l int, r int, L int, R int) (cnt []int, v int) {
		if L <= l && r <= R {
			return arr[i], val[i]
		}
		mid := (l + r) / 2
		if R <= mid {
			return query(2*i+1, l, mid, L, R)
		}
		if mid < L {
			return query(2*i+2, mid+1, r, L, R)
		}
		ac, av := query(2*i+1, l, mid, L, R)
		bc, bv := query(2*i+2, mid+1, r, L, R)
		cnt = make([]int, k)
		merge(cnt, ac, bc, av)
		v = (av * bv) % k
		return
	}

	ans := make([]int, len(queries))

	for j, cur := range queries {
		i, v, s, x := cur[0], cur[1], cur[2], cur[3]
		update(0, 0, n-1, i, v)
		tmp, _ := query(0, 0, n-1, s, n-1)
		ans[j] = tmp[x]
	}

	return ans
}
