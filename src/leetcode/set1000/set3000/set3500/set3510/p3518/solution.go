package p3518

func smallestPalindrome(s string, k int) string {
	cnt := make([]int, 26)
	for _, x := range []byte(s) {
		cnt[x-'a']++
	}

	nCr := func(n int, r int) int {
		// n * (n - 1) * (n - 2) * ... * (n - r + 1) / r!
		// 主要的问题是， (n, r) 超出k的时候
		if n-r < r {
			r = n - r
		}
		res := 1
		for i := 0; i < r; i++ {
			res *= n - i
			res /= i + 1
			if res > k {
				return k + 1
			}
		}
		return res
	}

	get := func(m int) int {
		res := 1
		for i := 0; i < 26; i++ {
			// nCr(m, cnt[i])
			if cnt[i] == 0 {
				continue
			}
			res *= nCr(m, cnt[i])
			if res > k {
				return k + 1
			}
			m -= cnt[i]
		}
		return res
	}
	n := len(s)

	// s 肯定是回文串
	mid := -1
	for i := range 26 {
		if cnt[i]&1 == 1 {
			mid = i
		}
		cnt[i] /= 2
	}
	h := n / 2
	buf := make([]byte, h)
	for i := 0; i < h; i++ {
		var tmp int
		var j int
		for j < 26 {
			if cnt[j] != 0 {
				cnt[j]--
				tmp2 := get(h - i - 1)
				if tmp+tmp2 >= k {
					break
				}
				tmp += tmp2
				cnt[j]++
			}
			j++
		}
		if j == 26 {
			return ""
		}
		k -= tmp
		buf[i] = byte(j + 'a')
	}
	res := make([]byte, n)
	copy(res, buf)
	if mid != -1 {
		res[n/2] = byte(mid + 'a')
	}
	for i := n - h; i < n; i++ {
		res[i] = res[h-1-(i-(n-h))]
	}
	return string(res)
}
