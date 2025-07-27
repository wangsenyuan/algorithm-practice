package p3630

import (
	"math/bits"
	"slices"
)

// 线性基模板
type xorBasis []int

func (b xorBasis) insert(x int) {
	for i := len(b) - 1; i >= 0; i-- {
		if x>>i&1 == 0 {
			continue
		}
		if b[i] == 0 {
			b[i] = x
			return
		}
		x ^= b[i]
	}
}

func (b xorBasis) maxXor() (res int) {
	for i := len(b) - 1; i >= 0; i-- {
		if res^b[i] > res {
			res ^= b[i]
		}
	}
	return
}

func maximizeXorAndXor(nums []int) int64 {
	n := len(nums)
	// 预处理所有子集的 AND 和 XOR（刷表法）
	type pair struct{ and, xor int }
	subSum := make([]pair, 1<<n)
	subSum[0].and = -1
	for i, x := range nums {
		highBit := 1 << i
		for mask, p := range subSum[:highBit] {
			subSum[highBit|mask] = pair{p.and & x, p.xor ^ x}
		}
	}
	subSum[0].and = 0

	sz := bits.Len(uint(slices.Max(nums)))
	b := make(xorBasis, sz)
	maxXor2 := func(sub uint) (res int) {
		clear(b)
		xor := subSum[sub].xor
		for ; sub > 0; sub &= sub - 1 {
			x := nums[bits.TrailingZeros(sub)]
			b.insert(x &^ xor) // 只考虑有偶数个 1 的比特位（xor 在这些比特位上是 0）
		}
		return xor + b.maxXor()*2
	}

	ans := 0
	u := 1<<n - 1
	for i, p := range subSum {
		ans = max(ans, p.and+maxXor2(uint(u^i)))
	}
	return int64(ans)
}
