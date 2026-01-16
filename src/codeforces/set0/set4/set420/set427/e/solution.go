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
	var n, m int
	fmt.Fscan(reader, &n, &m)

	x := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &x[i])
	}
	return solve(m, x)
}

func solve(m int, x []int) int {
	n := len(x)

	// 当警局在位置pos时的移动距离
	find := func(pos int) int {
		w := x[pos]
		var res int

		// 在位置w处的罪犯不需要开车去抓捕
		l := pos
		for l >= 0 && x[l] == w {
			l--
		}

		r := pos
		for r < n && x[r] == w {
			r++
		}

		for i := 0; i <= l; {
			// 先要移动到这里, 然后一路带回去m个人(没有必要超过pos)
			res += (w - x[i]) * 2
			i = min(l+1, i+m)
		}

		for i := n - 1; i >= r; {
			res += (x[i] - w) * 2
			i = max(r-1, i-m)
		}

		return res
	}

	return find(n / 2)
}
