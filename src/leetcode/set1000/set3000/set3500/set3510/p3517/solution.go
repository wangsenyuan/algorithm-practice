package p3517

func smallestPalindrome(s string) string {
	cnt := make([]int, 26)
	for _, x := range []byte(s) {
		cnt[x-'a']++
	}
	n := len(s)
	// s 是个回文串
	buf := make([]byte, n)

	if n&1 == 1 {
		for i := range 26 {
			if cnt[i]&1 == 1 {
				buf[n/2] = byte(i + 'a')
			}
			cnt[i] = cnt[i] / 2 * 2
		}
	}
	for i, l, r := 0, 0, n-1; i < 26; i++ {
		for cnt[i] > 0 {
			buf[l] = byte(i + 'a')
			buf[r] = byte(i + 'a')
			l++
			r--
			cnt[i] -= 2
		}
	}
	return string(buf)
}
