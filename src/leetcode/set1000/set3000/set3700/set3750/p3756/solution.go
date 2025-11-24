package p3756

const mod = 1e9 + 7

func add(a int, b int) int {
	return (a + b) % mod
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func sumAndMultiply(s string, queries [][]int) []int {
	n := len(s)
	pw := make([]int, n+1)
	pw[0] = 1
	for i := range n {
		pw[i+1] = mul(pw[i], 10)
	}
	s1 := make([]int, n+1)
	s2 := make([]int, n+1)
	cnt := make([]int, n+1)

	for i := range n {
		s1[i+1] = s1[i]
		s2[i+1] = s2[i]
		cnt[i+1] = cnt[i]
		if s[i] != '0' {
			x := int(s[i] - '0')
			s1[i+1] = add(mul(s1[i], 10), x)
			s2[i+1] = add(s2[i], x)
			cnt[i+1]++
		}
	}

	ans := make([]int, len(queries))
	for i, cur := range queries {
		l, r := cur[0], cur[1]
		d := cnt[r+1] - cnt[l]
		if d == 0 {
			ans[i] = 0
			continue
		}
		x := sub(s1[r+1], mul(s1[l], pw[d]))
		y := sub(s2[r+1], s2[l])
		ans[i] = mul(x, y)
	}

	return ans
}
