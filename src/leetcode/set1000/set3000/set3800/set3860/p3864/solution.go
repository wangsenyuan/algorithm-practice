package p3864

func minCost(s string, encCost int, flatCost int) int64 {
	n := len(s)

	pref := make([]int, n+1)
	for i := range n {
		pref[i+1] = pref[i]
		if s[i] == '1' {
			pref[i+1]++
		}
	}

	var f func(l int, r int) int

	f = func(l int, r int) int {
		var cur int
		if pref[r]-pref[l] == 0 {
			cur = flatCost
		} else {
			cur = (r - l) * (pref[r] - pref[l]) * encCost
		}
		if (r-l)%2 == 1 {
			// 奇数长度
			return cur
		}

		mid := (l + r) / 2
		return min(cur, f(l, mid)+f(mid, r))
	}

	res := f(0, n)

	return int64(res)
}
