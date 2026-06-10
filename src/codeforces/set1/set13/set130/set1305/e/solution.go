package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	res := solve(n, m)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func solve(n int, m int) []int {
	// Maximum achievable balance with n elements is reached by the
	// consecutive sequence 1,2,...,n, where value i contributes (i-1)/2
	// triples. If m exceeds that maximum, there is no solution.
	maxBalance := 0
	for i := 1; i <= n; i++ {
		maxBalance += (i - 1) / 2
	}
	if m > maxBalance {
		return nil
	}

	res := make([]int, 0, n)
	count := 0
	// Greedily extend the consecutive prefix 1,2,3,...; placing value v=len+1
	// adds (v-1)/2 triples (pairs a+b=v with a<b inside the prefix).
	for len(res) < n {
		v := len(res) + 1
		add := (v - 1) / 2
		if count+add > m {
			break
		}
		res = append(res, v)
		count += add
	}

	// If we still owe some triples, insert one value that contributes exactly
	// the remainder. The prefix is 1..p (p = len(res)); a value w in (p, 2p]
	// pairs as (w-p,p),(w-p+1,p-1),... so it forms exactly need pairs when
	// w = 2*p + 1 - 2*need. Since the greedy stopped with need < p/2, we have
	// w > p, keeping the array strictly increasing.
	if count < m {
		p := len(res)
		need := m - count
		res = append(res, 2*p+1-2*need)
		count += need
	}

	// Pad the remaining slots with large ascending values, all strictly above
	// limit/2, so any two of them sum to more than limit and can never equal a
	// third. The gap exceeds the largest small value, so no padding value is the
	// difference of two others either. Hence padding introduces no new triple.
	const limit = 1_000_000_000
	if r := n - len(res); r > 0 {
		gap := limit / (2 * n)
		base := limit - (r-1)*gap
		for j := 0; j < r; j++ {
			res = append(res, base+j*gap)
		}
	}

	return res
}
