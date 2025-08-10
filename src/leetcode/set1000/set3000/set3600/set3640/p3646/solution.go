package p3646

import (
	"fmt"
	"slices"
)

func specialPalindrome(n int64) int64 {
	var digits []int
	for tmp := n; tmp > 0; tmp /= 10 {
		digits = append(digits, int(tmp%10))
	}

	slices.Reverse(digits)

	type pair struct {
		first  int
		second int
	}

	calc := func(buf []int, cnt []pair, l int, r int) int {
		// 在l和r中间，尽量填最小的数，满足回文的条件
		for l <= r {
			if l < r {
				for i := range cnt {
					if cnt[i].second >= 2 {
						buf[l] = cnt[i].first
						buf[r] = cnt[i].first
						cnt[i].second -= 2
						break
					}
				}
			} else {
				for i := range cnt {
					if cnt[i].second == 1 {
						buf[l] = cnt[i].first
						cnt[i].second -= 1
						break
					}
				}
			}
			l++
			r--
		}

		var res int
		for _, v := range buf {
			res = res*10 + v
		}
		return res
	}

	ans := inf

	// 2468
	val := []int{2, 4, 6, 8}
	f := func(state int, mid int) int {
		var ln int
		cnt := make([]pair, 0, 5)
		for i := range 4 {
			if (state>>i)&1 == 1 {
				ln += val[i]
				cnt = append(cnt, pair{val[i], val[i]})
			}
		}
		if mid > 0 {
			cnt = append(cnt, pair{mid, mid})
		}
		// 1, 3, 5, 7
		ln += mid
		if ln < len(digits) || ln > len(fmt.Sprintf("%d", ans)) {
			return inf
		}

		slices.SortFunc(cnt, func(a, b pair) int {
			return a.first - b.first
		})

		// ln >= digits
		// 然后排列选中的数，让其满足比n大, 且组成回文
		buf := make([]int, ln)
		best := inf

		if ln == len(digits) {
			ok := false
			for i := 0; i < ln/2; i++ {
				d := digits[i]
				// 如果能够在这里放置一个数使buf[i] > d
				// 剩余的位置就可以随便放置
				// 否则这里必须放置d
				for j := range cnt {
					if cnt[j].second >= 2 && cnt[j].first > d {
						// 必须至少有两个
						cnt[j].second -= 2
						buf[i] = cnt[j].first
						buf[ln-1-i] = cnt[j].first
						best = min(best, calc(buf, slices.Clone(cnt), i+1, ln-i-2))
						cnt[j].second += 2
						break
					}
				}
				//如果要放置d, 则d必须要出现
				ok = false
				for j := range cnt {
					if cnt[j].first == d && cnt[j].second >= 2 {
						buf[i] = d
						buf[ln-1-i] = d
						cnt[j].second -= 2
						ok = true
						break
					}
				}

				if !ok {
					break
				}
			}

			if ok {
				if mid > 0 && digits[ln/2] <= mid {
					buf[ln/2] = mid
				}
				tmp := calc(buf, cnt, ln/2+1, ln/2-1)
				if tmp > int(n) {
					best = min(best, tmp)
				}
			}
		} else {
			// ln > len(digits)
			best = calc(buf, cnt, 0, ln-1)
		}

		return best
	}

	for state := 0; state < 1<<4; state++ {
		for _, mid := range []int{0, 1, 3, 5, 7, 9} {
			ans = min(ans, f(state, mid))
		}
	}

	return int64(ans)
}

const inf = 1 << 62
