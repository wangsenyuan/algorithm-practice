package main

import (
	"bufio"
	"fmt"
	"math/bits"
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
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, 2)
		fmt.Fscan(reader, &a[i][0], &a[i][1])
	}
	b := make([][]int, m)
	for i := 0; i < m; i++ {
		b[i] = make([]int, 2)
		fmt.Fscan(reader, &b[i][0], &b[i][1])
	}
	return solve(a, b)
}

func solve(a [][]int, b [][]int) int {

	check := func(p []int, arr [][]int) int {
		// 如果对方拿到的是p，在自己的set中，是否能够判断出p中，有一个数字和自己是common的
		// 如果对方的p数字在arr中没有出现，那么返回-2
		// 如果p[0]或者p[1]出现了，（只有一个），那么就是对应的数
		// 如果 p[0] 或者 p[1] 都出现了，那么就是 -1 （因为两个都有可能）
		x, y := min(p[0], p[1]), max(p[0], p[1])
		// x != y
		var flag int
		for _, cur := range arr {
			a, b := min(cur[0], cur[1]), max(cur[0], cur[1])
			if a == x && b == y {
				// 这个肯定是无效的
				continue
			}
			if a == x || b == x {
				flag |= 1
			}
			if a == y || b == y {
				flag |= 2
			}
		}
		if flag == 0 {
			return 0
		}
		if flag == 1 {
			return x
		}
		if flag == 2 {
			return y
		}
		return -1
	}
	// 当有多个值，但是这两个值不在同一个pair中时，就是0
	var flag1 int
	for _, cur := range a {
		tmp := check(cur, b)
		if tmp < 0 {
			return -1
		}
		if tmp > 0 {
			flag1 |= 1 << tmp
		}
	}

	var flag2 int
	for _, cur := range b {
		tmp := check(cur, a)
		if tmp < 0 {
			return -1
		}
		if tmp > 0 {
			flag2 |= 1 << tmp
		}
	}

	if flag1 != flag2 {
		return -1
	}
	if flag1&(flag1-1) == 0 {
		return bits.Len(uint(flag1)) - 1
	}
	return 0
}
