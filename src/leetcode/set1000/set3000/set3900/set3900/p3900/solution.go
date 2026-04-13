package p3900

import "slices"

func longestBalanced(s string) int {
	n := len(s)

	cnt := make([]int, 2)
	for i := range n {
		cnt[int(s[i]-'0')]++
	}
	play := func(s string) int {
		pos := make([]int, 2*n+1)
		for i := range pos {
			pos[i] = -2
		}
		pos[n] = -1

		var best int
		var sum int
		for r := range n {
			if s[r] == '0' {
				sum--
			} else {
				sum++
			}
			if pos[sum+n] != -2 {
				l := pos[sum+n]
				if (r-l)%2 == 0 {
					best = max(best, r-l)
				}
			}
			// 如果中间多了一个0
			if sum+2 <= n && pos[sum+2+n] != -2 {
				l := pos[sum+2+n]
				// freq[1] + freq[0] = r - l
				// freq[1] - freq[0] = -2
				if (r-l)%2 == 0 {
					c1 := (r - l - 2) / 2
					// c0 := c1 + 2
					if cnt[1] > c1 {
						best = max(best, r-l)
					}
				}
			}
			if sum-2 >= -n && pos[sum-2+n] != -2 {
				l := pos[sum-2+n]
				if (r-l)%2 == 0 {
					c1 := (r - l + 2) / 2
					c0 := c1 - 2
					if cnt[0] > c0 {
						best = max(best, r-l)
					}
				}
			}

			if pos[sum+n] == -2 {
				pos[sum+n] = r
			}
		}
		return best
	}

	res := play(s)
	buf := []byte(s)
	slices.Reverse(buf)
	s = string(buf)
	res = max(res, play(s))
	return res
}
