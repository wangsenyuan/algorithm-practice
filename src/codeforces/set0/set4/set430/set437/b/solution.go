package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("-1")
		return
	}
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (sum int, limit int, res []int) {
	fmt.Fscan(reader, &sum, &limit)
	res = solve(sum, limit)
	return
}

func solve(sum int, limit int) []int {
	// let y = lowbit(x), sum(yi) = sum
	// y = 1, 2, 4, 8, 16, 是确定的
	// 1的数量最多，可以放到最后面处理
	h := bits.Len(uint(limit))

	var ans []int

	// 比如limit = 7, 当 y = 1的时候, 能够使用的数字包扩, 1, 3, 5, 7
	// 当 y = 2的时候，能够使用包括 2, 6
	// 当 y = 4的时候，能够使用的包括 4

	for j := h - 1; j >= 0; j-- {
		y := 1 << j
		for v := 0; sum >= y; v++ {
			w := y + (v << (j + 1))
			if w > limit {
				break
			}
			sum -= y
			ans = append(ans, w)
		}
	}

	if sum > 0 {
		return nil
	}
	return ans
}
