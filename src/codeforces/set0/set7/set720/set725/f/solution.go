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
	var n int
	fmt.Fscan(reader, &n)
	piles := make([][]int, n)
	for i := range n {
		piles[i] = make([]int, 4)
		for j := range 4 {
			fmt.Fscan(reader, &piles[i][j])
		}
	}
	return solve(piles)
}

func solve(piles [][]int) int {
	n := len(piles)
	var res int

	// 收集参与排序的卡片（只收集 case1 堆的两张卡）
	// case2 堆（w_top < w_bottom）直接确定结果
	type item struct {
		a, b int
	}
	var items []item

	for i := range n {
		a1, b1, a2, b2 := piles[i][0], piles[i][1], piles[i][2], piles[i][3]
		w1, w2 := a1+b1, a2+b2
		if w1 >= w2 {
			// case 1: 正常堆，两张卡独立参与排序
			items = append(items, item{a1, b1}, item{a2, b2})
		} else {
			// case 2: 反转堆，两人都不想先拿
			// 如果 a1-b2 > 0，Alice 最终会先拿（收益 a1-b2）
			// 如果 a2-b1 < 0，Bonnie 最终会先拿（收益 a2-b1）
			// 否则两人都不动
			if a1-b2 > 0 {
				res += a1 - b2
			} else if a2-b1 < 0 {
				res += a2 - b1
			}
			// 否则 res += 0
		}
	}

	// 把 case1 的卡按 a+b 从大到小排序
	slices.SortFunc(items, func(u, v item) int {
		return (v.a + v.b) - (u.a + u.b)
	})

	// Alice 取奇数位（0,2,4,...），Bonnie 取偶数位（1,3,5,...）
	for i, it := range items {
		if i%2 == 0 {
			res += it.a
		} else {
			res -= it.b
		}
	}

	return res
}
