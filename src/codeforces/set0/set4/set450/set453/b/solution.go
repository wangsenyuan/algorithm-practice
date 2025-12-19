package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (a []int, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(a)
	return
}

const inf = 1 << 60

type data struct {
	val int
	w   int
}

func solve(a []int) []int {
	n := len(a)
	x := slices.Max(a)

	lpf := make([]int, 2*x)
	var primes []int
	for i := 2; i < 2*x; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= 2*x {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}

	fs := make([]int, 2*x+1)

	for pos, i := range primes {
		for j := i; j <= 2*x; j += i {
			fs[j] |= 1 << pos
		}
	}

	k := len(primes)
	K := 1 << k

	dp := make([][]data, n+1)
	for i := range n + 1 {
		dp[i] = make([]data, K)
		for j := range K {
			dp[i][j].val = inf
		}
	}
	dp[0][0].val = 0

	for j, v := range a {
		for state := range K {
			if dp[j][state].val < inf {
				for w := 1; w <= 2*x; w++ {
					if state&fs[w] == 0 {
						newState := state | fs[w]
						if dp[j+1][newState].val > dp[j][state].val+abs(v-w) {
							dp[j+1][newState].val = dp[j][state].val + abs(v-w)
							dp[j+1][newState].w = w
						}
					}
				}
			}
		}
	}

	best := dp[n][0].val
	var s int
	for state := range K {
		if dp[n][state].val < best {
			best = dp[n][state].val
			s = state
		}
	}

	var res []int

	for j := n; j > 0; j-- {
		c := dp[j][s].w
		res = append(res, c)
		s ^= fs[c]
	}

	slices.Reverse(res)
	return res
}

func abs(num int) int {
	return max(num, -num)
}
