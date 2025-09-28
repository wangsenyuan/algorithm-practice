package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	var n, m, l, f int
	fmt.Fscan(reader, &n, &m, &l, &f)
	steps := make([][]int, m)
	for i := 0; i < m; i++ {
		steps[i] = make([]int, 3)
		fmt.Fscan(reader, &steps[i][0], &steps[i][1], &steps[i][2])
	}
	return solve(n, l, f, steps)
}

func solve(n int, s int, f int, steps [][]int) string {
	// 永远超一个方向
	var ans []byte

	play := func(x int, y int) {
		// 可以移动这么远
		t := y - x
		if s < f {
			for s < f && t > 0 {
				ans = append(ans, 'R')
				s++
				t--
			}
		} else {
			for s > f && t > 0 {
				ans = append(ans, 'L')
				s--
				t--
			}
		}
	}

	m := len(steps)
	if steps[0][0] > 1 {
		play(1, steps[0][0])
	}

	for i := 0; i < m && s != f; i++ {
		x := steps[i][0]
		y := inf
		if i+1 < m {
			y = steps[i+1][0]
		}
		l, r := steps[i][1], steps[i][2]
		if l <= s && s <= r || s == l-1 && s < f || s == r+1 && s > f {
			// 只需要等一秒
			ans = append(ans, 'X')
			x++
		}
		if x < y {
			play(x, y)
		}
	}

	return string(ans)
}

const inf = 1 << 60

func abs(a int) int {
	return max(a, -a)
}
