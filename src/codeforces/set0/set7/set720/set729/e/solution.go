package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, s int
	fmt.Fscan(reader, &n, &s)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(s, a)
}

func solve(s int, a []int) int {
	n := len(a)
	s--
	var w int
	if a[s] > 0 {
		w++
		a[s] = 0
	}
	cnt := make([]int, n)
	var bad int
	for _, v := range a {
		if v >= n {
			bad++
		} else {
			cnt[v]++
		}
	}
	cnt[0]--
	bad += cnt[0]

	// 如果高度最高的高度为h时，最优的结果是什么
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] + cnt[i]
	}

	best := n - 1
	var need int
	var extra int
	for i := 1; i < n; i++ {
		if cnt[i] == 0 {
			need++
		} else {
			extra += cnt[i] - 1
		}

		has := bad + suf[i+1]
		if has >= need {
			// 后面的全部在说谎
			best = min(best, has)
		} else if need-has <= extra {
			// 多出来的那些人可以补上，但只需要补充need即可
			best = min(best, need)
		}
	}

	return best + w
}
