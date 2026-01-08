package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) (cnt []int, res string) {
	cnt = make([]int, 4)
	for i := range 4 {
		fmt.Fscan(reader, &cnt[i])
	}
	res = solve(cnt[0], cnt[1], cnt[2], cnt[3])
	return
}

func calc(n int) int {
	// x * (x - 1) / 2 = n
	x := sort.Search(n+2, func(x int) bool {
		return x*(x-1)/2 > n
	})
	return x - 1
}

func solve(a00, a01, a10, a11 int) string {
	// x * (x - 1) / 2 = a00
	x := calc(a00)
	y := calc(a11)

	if x*(x-1)/2 != a00 || y*(y-1)/2 != a11 {
		return "Impossible"
	}
	if a01+a10 != x*y {
		if a01+a10 > 0 || x > 1 && y > 1 {
			return "Impossible"
		}
		if x == 1 {
			x--
		} else {
			y--
		}
	}

	buf := make([]byte, x+y)
	// 0000111,
	for i := range x {
		buf[i] = '0'
	}
	for i := x; i < x+y; i++ {
		buf[i] = '1'
	}
	b01 := x * y
	// 最后一个0
	pos := x - 1
	for b01 > a01 {
		d := min(y, b01-a01)
		buf[pos], buf[pos+d] = buf[pos+d], buf[pos]
		b01 -= d
		pos--
	}

	return string(buf)
}

func abs(num int) int {
	return max(num, -num)
}
