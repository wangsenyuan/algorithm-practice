package p3757

import "math/bits"

const mod = 1e9 + 7

func add(a, b int) int {
	return (a + b) % mod
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func countEffective(nums []int) int {
	var or int
	for _, v := range nums {
		or |= v
	}

	h := bits.Len(uint(or))

	T := 1 << h
	freq := make([]int, T)
	for _, v := range nums {
		freq[v]++
	}

	for i := range h {
		if or>>i&1 == 0 {
			continue
		}
		bit := 1 << i
		for s := 0; s < T; s++ {
			s |= bit
			freq[s] += freq[s^bit]
		}
	}

	n := len(nums)

	pw := make([]int, n+1)
	pw[0] = 1
	for i := 1; i <= n; i++ {
		pw[i] = mul(pw[i-1], 2)
	}
	ans := pw[n]

	for s := or; s >= 0; {
		diff := or ^ s
		cnt := bits.OnesCount(uint(diff))
		if cnt&1 == 0 {
			ans = sub(ans, pw[freq[s]])
		} else {
			ans = add(ans, pw[freq[s]])
		}
		if s == 0 {
			break
		}
		s = (s - 1) & or
	}
	return ans
}
