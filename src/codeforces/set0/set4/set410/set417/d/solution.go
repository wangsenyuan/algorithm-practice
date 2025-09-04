package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m, b int
	fmt.Fscan(reader, &n, &m, &b)
	requirements := make([][]int, n)
	problems := make([][]int, n)
	for i := range n {
		requirements[i] = make([]int, 3)
		fmt.Fscan(reader, &requirements[i][0], &requirements[i][1], &requirements[i][2])
		problems[i] = make([]int, requirements[i][2])
		for j := range requirements[i][2] {
			fmt.Fscan(reader, &problems[i][j])
		}
	}
	return solve(m, b, requirements, problems)
}

const inf = 1 << 60

type friend struct {
	x        int
	k        int
	problems []int
}

func solve(m int, b int, requirements [][]int, problems [][]int) int {
	M := 1 << m
	dp := make([]int, M)
	for i := range M {
		dp[i] = inf
	}
	dp[0] = 0
	n := len(requirements)

	friends := make([]friend, n)
	for i := range n {
		friends[i] = friend{
			x:        requirements[i][0],
			k:        requirements[i][1],
			problems: problems[i],
		}
	}

	slices.SortFunc(friends, func(a, b friend) int {
		return b.k - a.k
	})

	for _, cur := range friends {
		x, k := cur.x, cur.k
		var mask int
		for _, v := range cur.problems {
			mask |= 1 << (v - 1)
		}
		for s := 0; s < M; s++ {
			if s == 0 {
				dp[s|mask] = min(dp[s|mask], dp[s]+x+k*b)
			} else {
				dp[s|mask] = min(dp[s|mask], dp[s]+x)
			}
		}
	}
	if dp[M-1] == inf {
		return -1
	}
	return dp[M-1]
}
