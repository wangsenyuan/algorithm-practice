package p3816

func lexSmallestAfterDeletion(s string) string {
	cnt := make([]int, 26)
	for _, x := range s {
		cnt[x - 'a']++
	}

	var st []rune

	for _, ch := range s {
		for len(st) > 0 && st[len(st)-1] > ch && cnt[st[len(st)-1]-'a'] > 1 {
			cnt[st[len(st)-1]-'a']--
			st = st[:len(st)-1]
		}
		st = append(st, ch)
	}

	for len(st) > 0 && cnt[st[len(st)-1]-'a'] > 1 {
		cnt[st[len(st)-1]-'a']--
		st = st[:len(st)-1]
	}

	return string(st)
}
