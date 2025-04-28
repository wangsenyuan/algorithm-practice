package p3533

import "fmt"

type data struct {
	val  int
	next int
	rem  int
	id   int
}

func concatenatedDivisibility(nums []int, k int) []int {
	n := len(nums)
	N := 1 << n
	dp := make([][]data, N)
	for i := range N {
		dp[i] = make([]data, k)
		for j := range k {
			dp[i][j] = data{-1, n, 0, -1}
		}
	}

	base := make([]int, n*6+10)
	base[0] = 1
	for i := 1; i < len(base); i++ {
		base[i] = base[i-1] * 10 % k
	}

	L := make([]int, n)
	for i := range n {
		L[i] = len(fmt.Sprintf("%d", nums[i]))
	}

	mul := func(x int, y int) int {
		return x * y % k
	}

	add := func(x int, y int) int {
		return (x + y) % k
	}

	update := func(state int, k int, v int, prev int, x int, i int) {
		if dp[state][k].val == -1 || v < dp[state][k].val {
			dp[state][k] = data{v, prev, x, i}
		}
	}

	for i := 0; i < n; i++ {
		dp[1<<i][nums[i]%k] = data{nums[i], 0, 0, i}
	}

	for state := 0; state < N; state++ {
		var sum_len int
		for i := range n {
			if (state>>i)&1 == 1 {
				sum_len += L[i]
			}
		}
		for x := 0; x < k; x++ {
			if dp[state][x].val == -1 {
				continue
			}
			for i := range n {
				if (state>>i)&1 == 0 {
					next := state | 1<<i
					nx := mul(nums[i], base[sum_len])
					nx = add(nx, x)
					update(next, nx, nums[i], state, x, i)
				}
			}
		}
	}

	if dp[N-1][0].val == -1 {
		return nil
	}

	var res []int
	cur_state := N - 1
	cur_k := 0

	for len(res) < n {
		s := dp[cur_state][cur_k]
		res = append(res, nums[s.id])
		cur_state = s.next
		cur_k = s.rem
	}

	return res
}
