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

func drive(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	blocked := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &blocked[i])
	}
	costs := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &costs[i])
	}
	return solve(n, blocked, costs)
}

func solve(n int, s []int, costs []int) int {
	// TODO: solve by hand first.
	blocked := make([]bool, n)
	for _, i := range s {
		blocked[i] = true
	}

	if blocked[0] {
		return -1
	}

	next := make([]int, n+1)
	next[n] = n

	for i := n - 1; i >= 0; i-- {
		next[i] = next[i+1]
		if blocked[i] {
			next[i] = i
		}
	}

	before := make([]int, n)
	for i := range n {
		before[i] = i
		if blocked[i] {
			before[i] = before[i-1]
		}
	}

	k := len(costs)

	best := 1 << 60

	for t := 1; t <= k; t++ {
		var cnt int
		for i := 0; i < n; {
			cnt++
			// next[i] 是下一个被block的位置, 或者是n
			if i+t >= n {
				break
			}
			// i + t < n - 1
			// j1是下一个开始的位置
			j1 := before[i+t]
			if i == j1 {
				cnt = -1
				break
			}
			// j2是下一个block的位置
			j2 := next[i]
			if j2 < j1 {
				i = j1
				continue
			}
			// j1 < j2 (j1 != j2)
			// i, i + t, i + 2 * t, ... i + w * t < j2
			w := (j2 - i) / t
			if i+w*t == j2 {
				w--
			}
			j3 := i + w*t
			if j3 < j1 {
				i = j1
				continue
			}
			cnt += w - 1
			i = j3
		}
		if cnt != -1 {
			best = min(best, costs[t-1]*cnt)
		}
	}

	if best == 1<<60 {
		return -1
	}

	return best
}
