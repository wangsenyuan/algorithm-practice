package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	res := solve(n)
	fmt.Println(res)
}

func solve(n int) int {
	mem := make(map[int]int)
	var f func(x int) int

	f = func(x int) int {
		if x <= 10 {
			return min(x, 13-x)
		}
		if v, ok := mem[x]; ok {
			return v
		}
		// 得到一个更小的x
		y := 1
		cnt := 1
		for y*10+1 <= x {
			y = y*10 + 1
			cnt++
		}
		if y == x {
			mem[x] = cnt
			return cnt
		}
		// y <= x
		w := x / y
		res := w*cnt + f(x-w*y)
		y2 := y*10 + 1
		cnt2 := cnt + 1
		// x = y2 - w * y - x2
		// x2 < x
		// x2 = y2 - w * y - x
		w = (y2 - x) / y
		// y2 - w * y >= x
		x2 := y2 - w*y - x
		if x2 == 0 {
			res = min(res, cnt2+w*cnt)
		} else {
			// x2 > 0
			res = min(res, cnt2+w*cnt+f(x2))
		}
		mem[x] = res
		return res
	}

	return f(n)
}
