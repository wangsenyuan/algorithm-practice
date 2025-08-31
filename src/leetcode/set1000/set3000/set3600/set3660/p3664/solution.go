package p3664

func score(cards []string, x byte) int {
	// 要么是 x?, 要么是?x, 还有就是xx
	var cnt int
	a1 := make([]int, 26)
	a2 := make([]int, 26)

	for _, card := range cards {
		if card[0] == x && card[1] == x {
			cnt++
		} else if card[0] == x {
			a1[int(card[1]-'a')]++
		} else if card[1] == x {
			a2[int(card[0]-'a')]++
		}
	}
	work := func(freq []int, cnt int) int {
		mv := cnt
		sum := cnt
		for i := range 26 {
			mv = max(mv, freq[i])
			sum += freq[i]
		}
		if 2*mv <= sum {
			return sum / 2
		}
		return sum - mv
	}

	var best int

	for u := 0; u <= cnt; u++ {
		best = max(best, work(a1, u)+work(a2, cnt-u))
	}

	return best
}
