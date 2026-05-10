package p3922

func minFlips(s string) int {
	n := len(s)
	if n < 3 {
		return 0
	}

	cnt := make([]int, 2)

	for _, x := range s {
		cnt[int(x-'0')]++
	}
	res := min(cnt[0], cnt[1])

	if cnt[1] > 0 {
		// 只剩下一个1
		res = min(res, cnt[1]-1)
	}
	if cnt[1] >= 2 && s[0] == '1' && s[n-1] == '1' {
		res = min(res, cnt[1]-2)
	}

	return res
}
