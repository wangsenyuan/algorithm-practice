package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	x, y := drive(reader)
	fmt.Println(x, y)
}

func drive(reader *bufio.Reader) (int, int) {
	var n int
	fmt.Fscan(reader, &n)
	return solve(n)
}

func solve(n int) (int, int) {
	if n == 0 {
		return 0, 0
	}
	// 第0圈，移动0步，第1圈移动6，第二圈移动12
	// 所以第i圈，移动 6 * i. 全部 = 6 * (0 + 1 + 2 + .. + i)

	count := func(i int) int {
		if i*(i+1) > n/3 || i*(i+1)*3 > n {
			return n + 1
		}
		return i * (i + 1) * 3
	}

	k := sort.Search(1e9, func(i int) bool {
		return count(i) >= n
	})

	// count(k) >= n
	w := count(k)
	if w == n {
		return k * 2, 0
	}
	// w > count(k-1)
	n -= count(k - 1)
	x, y := k*2, 0

	for _, cur := range [][]int{{-1, 2}, {-2, 0}, {-1, -2}, {1, -2}, {2, 0}, {1, 2}} {
		dx, dy := cur[0], cur[1]
		if n <= k {
			x += dx * n
			y += dy * n
			break
		}
		x += dx * k
		y += dy * k
		n -= k
	}

	return x, y
}
