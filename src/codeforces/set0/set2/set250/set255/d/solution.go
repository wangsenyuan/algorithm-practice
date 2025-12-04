package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := drive(reader)
	fmt.Println(ans)
}

func drive(reader *bufio.Reader) int {
	var n, x, y, c int
	fmt.Fscan(reader, &n, &x, &y, &c)
	return solve(n, x, y, c)
}

func solve(n int, x int, y int, c int) int {
	x--
	y--

	calc := func(w int) int {
		return 1 + 4*w*(w+1)/2
	}

	calc1 := func(w int, x int) int {
		s1 := calc(w - x - 1)
		s2 := (s1 - 1) / 4
		s2 *= 2
		s2 += w - x - 1
		s2 += 1
		return s2
	}

	// 传播1/4圆周
	calc2 := func(w int, x int) int {
		s1 := calc(w - x)
		s2 := (s1 - 1) / 4
		s2 -= w - x - 1
		s2 -= 1
		return s2
	}

	play := func(w int) int {
		// 1 + 4 + 8 + ... + 4 * (w + 1) - 4
		// 1, 4, 8, 12, ... 4 * w
		// 1 + 4 * (1 + w) * w / 2
		// sum := 1 + 4*(1+w)*w/2
		sum := calc(w)
		// 这个算起来好麻烦

		// 上部
		if w > x {
			// 在位置(-1, y)处，先上传播 w - x - 1秒
			sum -= calc1(w, x)
		}

		// 左边
		if w > y {
			sum -= calc1(w, y)
		}

		// 左上角需要加回来
		if w > x+y {
			sum += calc2(w, x+y)
		}

		// 往右边传播
		if w > n-1-y {
			sum -= calc1(w, n-1-y)
		}

		// 右上角
		if w > x+n-1-y {
			sum += calc2(w, x+n-1-y)
		}

		if w > n-1-x {
			sum -= calc1(w, n-1-x)
		}

		if w > y+n-1-x {
			sum += calc2(w, y+n-1-x)
		}

		if w > n-1-y+n-1-x {
			sum += calc2(w, n-1-y+n-1-x)
		}

		return sum
	}

	lo, hi := 0, n*2
	for lo < hi {
		mid := (lo + hi) / 2
		if play(mid) >= c {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return hi
}
