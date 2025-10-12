package p3714

type data struct {
	ab int
	bc int
	ca int
}

func longestBalanced(s string) int {
	ans := max(solve1(s), solve3(s))
	ch := "abc"
	for i := range 3 {
		ans = max(ans, solve2(s, ch[i]))
	}
	return ans
}

func solve1(s string) int {
	var res int
	for i := 0; i < len(s); {
		j := i
		for i < len(s) && s[i] == s[j] {
			i++
		}
		res = max(res, i-j)
	}
	return res
}

func solve2(s string, avoid byte) int {
	pos := make(map[int]int)
	pos[0] = -1

	cnt := make([]int, 3)

	var best int

	for i := range len(s) {
		if s[i] == avoid {
			clear(cnt)
			pos = make(map[int]int)
			pos[0] = i
			continue
		}
		x := int(s[i] - 'a')
		cnt[x]++
		var val int
		switch avoid {
		case 'a':
			val = cnt[1] - cnt[2]
		case 'b':
			val = cnt[0] - cnt[2]
		default:
			val = cnt[0] - cnt[1]
		}
		if j, ok := pos[val]; ok {
			best = max(best, i-j)
		} else {
			pos[val] = i
		}
	}
	return best
}

func solve3(s string) int {
	pos := make(map[data]int)
	pos[data{0, 0, 0}] = -1

	// ar - al = br - bl
	// ar - br = al - bl

	var best int
	cnt := make([]int, 3)
	for i := range len(s) {
		x := int(s[i] - 'a')
		cnt[x]++
		cur := data{cnt[0] - cnt[1], cnt[1] - cnt[2], cnt[2] - cnt[0]}
		if j, ok := pos[cur]; ok {
			best = max(best, i-j)
		} else {
			pos[cur] = i
		}
	}
	return best
}
