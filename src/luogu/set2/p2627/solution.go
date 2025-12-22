package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
}

func process(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

type pair struct {
	first  int
	second int
}

func solve(k int, a []int) int {
	n := len(a)
	// 理解错了，不是只选k个，而是不能连续选k个
	// 如果选择了一段 l...r
	// 按么 res[r] = sum[l...r] + res[l-2]
	//     res[r] = max(sum[r] - sum[l-1] + res[l-2])
	//  res[r] = sum[r] + max(res[l-1] - sum[l])
	// 且 l >= r - k- 1
	que := make([]pair, n+1)
	var head, tail int
	var sum int
	var res int

	que[head] = pair{-a[0], 0}
	head++

	for r := 1; r <= n; r++ {
		sum += a[r-1]
		for tail < head && que[tail].second < r-k-1 {
			tail++
		}
		cur := sum
		if r > k && tail < head {
			cur += que[tail].first
		}
		res = max(res, cur)
		if r+1 <= n {
			cur -= sum + a[r]
		}
		for head > tail && que[head-1].first <= cur {
			head--
		}
		que[head] = pair{cur, r}
		head++
	}

	return res
}
