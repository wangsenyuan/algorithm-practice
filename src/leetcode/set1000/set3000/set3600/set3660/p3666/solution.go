package p3666

func minOperations(s string, k int) int {
	var a, b int
	for i := range len(s) {
		if s[i] == '0' {
			a++
		} else {
			b++
		}
	}
	var res int
	for a >= 2*k {
		a -= k
		b += k
		res++
	}

	for b >= 2*k {
		b -= k
	}
	n := a + b

	getExtArea := func(a int) (int, int) {
		b := n - a
		l := max(0, k-b)
		r := min(k, a)
		l, r = a+k-2*r, a+k-2*l
		return l, r
	}

	type pair struct {
		first  int
		second int
	}

	marked := make(map[pair]bool)
	marked[pair{a, a}] = true

	l, r := a, a

	for l > 0 {
		if l <= k && k <= r && l&1 == k&1 {
			res++
			break
		}
		u, v := getExtArea(l)
		x, y := getExtArea(r)
		l = min(u, x)
		r = max(v, y)
		if marked[pair{l, r}] {
			return -1
		}
		marked[pair{l, r}] = true
		res++
	}

	return res
}
