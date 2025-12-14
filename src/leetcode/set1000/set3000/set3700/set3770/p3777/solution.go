package p3777

func minDeletions(s string, queries [][]int) []int {
	n := len(s)
	cnt := make(BIT, n+3)

	for i := range n {
		if i == 0 || s[i] != s[i-1] {
			cnt.update(i, 1)
		}
	}

	buf := []byte(s)

	change := func(i int) {
		if i > 0 {
			if buf[i] != buf[i-1] {
				// BA => BB
				cnt.update(i, -1)
			} else {
				// BB => BA
				cnt.update(i, 1)
			}
		}
		if i < n-1 {
			if buf[i] != buf[i+1] {
				// AB => AA
				cnt.update(i+1, -1)
			} else {
				// AA => AB
				cnt.update(i+1, 1)
			}
		}
		if buf[i] == 'B' {
			buf[i] = 'A'
		} else {
			buf[i] = 'B'
		}
	}

	find := func(l int, r int) int {
		res := cnt.rangeQuery(l, r)
		if l > 0 && buf[l-1] == buf[l] {
			res++
		}
		// 这个区间内，ABABA的分段数
		return r - l + 1 - res
	}

	var ans []int

	for _, cur := range queries {
		if cur[0] == 1 {
			change(cur[1])
		} else {
			ans = append(ans, find(cur[1], cur[2]))
		}
	}

	return ans
}

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) query(i int) int {
	i++
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) rangeQuery(l int, r int) int {
	return bit.query(r) - bit.query(l-1)
}
