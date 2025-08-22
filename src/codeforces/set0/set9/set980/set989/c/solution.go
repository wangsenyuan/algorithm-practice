package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b, c, d int
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
	res := solve(a, b, c, d)
	fmt.Println(len(res), len(res[0]))
	for _, s := range res {
		fmt.Println(s)
	}
}

var flowers = "ABCD"

func solve(a, b, c, d int) []string {
	// 分成4块, 固定50列

	arr := []int{a, b, c, d}

	doIt := func(x byte, y byte, n int) []string {
		// 以x为背景，放置y的块
		border := strings.Repeat(string(x), 50)
		var res []string
		res = append(res, border)
		// 至少3行
		row := (n + 23) / 24
		for i := range row {
			buf := make([]byte, 50)
			for j := range buf {
				buf[j] = x
				// 最后一列使用x
				if j < 48 && i%2 == j%2 && n > 0 {
					buf[j] = y
					n--
				}
			}
			res = append(res, string(buf))
			res = append(res, border)
		}
		return res

	}
	var res []string

	for i := 0; i < 4; i++ {
		tmp := doIt(flowers[i], flowers[(i+1)%4], arr[(i+1)%4]-1)
		res = append(res, tmp...)
	}
	return res
}
