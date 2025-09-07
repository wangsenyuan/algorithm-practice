package p3677

import "math/bits"

func countBinaryPalindromes(n int64) int {
	if n <= 1 {
		return 1 + int(n)
	}
	m := bits.Len(uint(n))
	k := (m - 1) / 2

	// 二进制长度小于 m
	ans := 2<<k - 1
	if m%2 == 0 {
		ans += 1 << k
	}

	// 二进制长度等于 m，且回文数的左半小于 n 的左半
	left := n >> (m / 2)
	ans += int(left) - 1<<k

	// 二进制长度等于 m，且回文数的左半等于 n 的左半
	right := bits.Reverse32(uint32(left>>(m%2))) >> (32 - m/2)
	if left<<(m/2)|int64(right) <= n {
		ans++
	}

	return ans
}
