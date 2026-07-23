package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int64 {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	w := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &w[i])
	}
	return solve(b, w)
}

func solve(b, w []int) int64 {
	// 选择黑色的, 至少比白色的多
	slices.Sort(b)
	slices.Reverse(b)
	slices.Sort(w)
	slices.Reverse(w)
	// 假设选择了k个黑色的, 那么就是前k个白色种的最大值

	var sum1, sum2 int64
	var best int64
	var ans int64

	for i, j := 0, 0; i < len(b); i++ {
		sum1 += int64(b[i])
		for j < len(w) && j <= i {
			sum2 += int64(w[j])
			best = max(best, sum2)
			j++
		}
		ans = max(ans, sum1+best)
	}

	return ans
}
