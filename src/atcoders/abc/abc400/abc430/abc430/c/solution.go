package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int64 {
	var n, a, b int
	fmt.Fscan(reader, &n, &a, &b)
	var s string
	fmt.Fscan(reader, &s)
	return solve(n, a, b, s)
}

func solve(n, a, b int, s string) int64 {
	var ans int
	cnt := [2]int{}
	var l1, l2 int

	for _, ch := range s {
		cnt[ch-'a']++
		for cnt[0] >= a {
			if s[l1] == 'a' {
				cnt[0]--
			}
			l1++
		}
		for cnt[1] >= b {
			if s[l2] == 'b' {
				cnt[1]--
			}
			l2++
		}
		ans += max(l1-l2, 0)
	}

	return int64(ans)
}
func solve1(n, a, b int, s string) int64 {
	// s1[r] - s1[l] >= a
	// s2[r] - s2[l] <= b - 1 => s2[l] >= s2[r] - b + 1

	bit := make([]int, n+2)

	add := func(p int, v int) {
		p++
		for p < len(bit) {
			bit[p] += v
			p += p & -p
		}
	}

	get := func(p int) int {
		p++
		var res int
		for p > 0 {
			res += bit[p]
			p -= p & -p
		}
		return res
	}

	var l int
	var s1, s2 int
	var res int

	add(0, 1)
	for r := range n {
		if s[r] == 'a' {
			s1++
		}

		for s1-s2 >= a {
			if s[l] == 'a' {
				s2++
			}
			// l + 1 - s2 = b的数量
			if s1-s2 < a {
				if s[l] == 'a' {
					s2--
				}
				break
			}
			add(l+1-s2, 1)
			l++
		}
		if s1-s2 >= a {
			w := (r + 1 - s1) - (b - 1)
			res += get(n)
			if w >= 0 {
				res -= get(w - 1)
			}
		}
	}

	return int64(res)
}
