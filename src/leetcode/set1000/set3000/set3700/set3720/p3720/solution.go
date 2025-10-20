package p3720

import (
	"math/bits"
	"strings"
)

func lexGreaterPermutation(s, target string) string {
	left := make([]int, 26)
	for i, b := range s {
		left[b-'a']++
		left[target[i]-'a']--
	}

	neg, mask := 0, 0
	for i, cnt := range left {
		if cnt < 0 {
			neg++ // 统计 left 中的负数个数
		} else if cnt > 0 {
			mask |= 1 << i
		}
	}

	ans := []byte(target)
	for i := len(s) - 1; i >= 0; i-- {
		b := target[i] - 'a'
		left[b]++ // 撤销消耗

		if left[b] == 0 {
			neg--
		} else if left[b] == 1 {
			mask |= 1 << b
		}

		// left 有负数 or 没有大于 target[i] 的字母
		if neg > 0 || mask>>(b+1) == 0 {
			continue
		}

		// target[i] 增大到 j
		mask &^= 1<<(b+1) - 1
		j := bits.TrailingZeros(uint(mask)) // 也可以写循环找 j
		left[j]--
		ans[i] = 'a' + byte(j)
		ans = ans[:i+1]

		for k, c := range left {
			ch := string('a' + byte(k))
			ans = append(ans, strings.Repeat(ch, c)...)
		}
		return string(ans)
	}
	return ""
}

func lexGreaterPermutation1(s string, target string) string {
	f1 := make([]int, 26)
	n := len(s)
	for i := range n {
		f1[int(s[i]-'a')]++
	}

	m := len(target)

	next_diff := make([]int, m+1)
	next_diff[m] = m

	for i := m - 1; i >= 0; i-- {
		if i == m-1 || target[i] == target[i+1] {
			next_diff[i] = next_diff[i+1]
		} else {
			next_diff[i] = i + 1
		}
	}

	buf := make([]byte, n)

	build := func(p1 int, c int) string {
		buf[p1] = byte(c + 'a')
		f1[c]--
		var j int
		for i := p1 + 1; i < n; i++ {
			for f1[j] == 0 {
				j++
			}
			buf[i] = byte(j + 'a')
			f1[j]--
		}
		return string(buf)
	}

	f2 := make([]int, 26)

	check := func(pos int) bool {
		copy(f2, f1)
		w := 25
		for pos < m {
			for w >= 0 && f2[w] == 0 {
				w--
			}
			if w < 0 {
				return false
			}
			if w > int(target[pos]-'a') {
				return true
			}
			if w < int(target[pos]-'a') {
				return false
			}

			next := next_diff[pos]
			cnt := min(f2[w], next-pos)
			pos += cnt
			f2[w] -= cnt
		}

		return false
	}

	// 不是找刚好比target大的序列
	for i := range n {
		buf[i] = '-'
		for j := range 26 {
			if f1[j] > 0 {
				if i == m || int(target[i]-'a') < j {
					return build(i, j)
				}
				if int(target[i]-'a') == j {
					f1[j]--
					if check(i + 1) {
						buf[i] = target[i]
						break
					}
					f1[j]++
				}
			}
		}
		if buf[i] == '-' {
			return ""
		}
	}
	// 那么使用那个 next-permutation 交换，如果可以，就ok

	return string(buf)
}
