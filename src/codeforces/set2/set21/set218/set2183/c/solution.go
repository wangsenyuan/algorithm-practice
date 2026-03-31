package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var n, m, k int
		fmt.Fscan(reader, &n, &m, &k)
		res := solve(n, m, k)
		fmt.Fprintln(writer, res)
	}
}

func solve(n int, m int, k int) int {
	l := k - 1
	r := n - k
	if l < r {
		l, r = r, l
	}
	// l >= r
	// 假设覆盖了左边w个区域
	best := 1
	for w := 1; w <= l; w++ {
		// 覆盖左边w的区域, 覆盖右边v的区域
		// 第一个士兵移动w个位置，同时出现了w个新的士兵，然后他们依次移动w - 1次
		if 2*w-1 > m {
			break
		}
		m1 := m - (2*w - 1)
		// 比如w = 3, 那么3天后，在k处有3个士兵，第4天移动2个士兵，第5天移动k-1处1个士兵
		// 但是在k处，此时又有3个士兵
		// 还剩余 m1的时间，此时在k处有w个士兵
		// 这些士兵，可以覆盖w个位置,位置k可以不管，因为它可以再长出来
		v := min(r, w, m1)
		best = max(best, w+v+1)
	}

	return best
}
