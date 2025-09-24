package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	return solve(c)
}

func solve(c []int) bool {
	n := len(c)
	if n == 1 {
		return true
	}
	if n == 2 {
		return false
	}

	slices.Sort(c)
	if c[n-1] != n {
		return false
	}

	var pos int
	var roots []int

	for pos < n && c[pos] == 1 {
		roots = append(roots, 1)
		pos++
	}

	find := func(v int) int {
		// 找到一个sum = v，且节点数量 >= 2的组合
		res := -1
		m := len(roots)
		for mask := 1; mask < 1<<m; mask++ {
			var sum int
			var cnt int
			for i := range m {
				if (mask>>i)&1 == 1 {
					sum += roots[i]
					cnt++
				}
			}
			if sum == v && cnt >= 2 {
				if res < 0 || cnt < bits.OnesCount(uint(res)) {
					res = mask
				}
			}
		}

		return res
	}

	for pos < n-1 {
		v := c[pos]
		// 要找到sum = v-1的节点组合
		flag := find(v - 1)
		if flag < 0 {
			return false
		}
		var buf []int
		for i, v := range roots {
			if (flag>>i)&1 == 0 {
				buf = append(buf, v)
			}
		}
		buf = append(buf, v)
		roots = buf
		pos++
	}

	return true
}
