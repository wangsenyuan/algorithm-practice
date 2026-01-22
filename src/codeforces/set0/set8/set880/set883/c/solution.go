package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var f, T, t0, a1, t1, p1, a2, t2, p2 int
	fmt.Fscan(reader, &f, &T, &t0, &a1, &t1, &p1, &a2, &t2, &p2)
	return solve(f, T, t0, a1, t1, p1, a2, t2, p2)
}

const inf = 1 << 60

func solve(f int, T int, t0 int, a1 int, t1 int, p1 int, a2 int, t2 int, p2 int) int {
	if f*t0 <= T {
		return 0
	}

	if t1 > t2 {
		t1, t2 = t2, t1
		p1, p2 = p2, p1
		a1, a2 = a2, a1
	}

	best := inf

	find := func(T int, f int) int {
		if T < 0 {
			// 已经超时了
			return inf
		}
		// 使用0,2在时间T内下载f的量
		// t0 * f0 + f2 * f2 = T
		// f0 + f2 = f
		if f*t0 <= T {
			return 0
		}
		// 全部使用2下载，需要购买这么多次
		m := (f + a2 - 1) / a2
		if min(f, m*a2)*t2 > T {
			return inf
		}
		// 目标是使的f2最小
		f2 := sort.Search(m, func(i int) bool {
			return min(f, i*a2)*t2+max(0, (f-i*a2))*t0 <= T
		})
		if min(f, f2*a2)*t2+max(0, f-f2*a2)*t0 <= T {
			return f2 * p2
		}
		return inf
	}

	for f1 := 0; ; f1++ {
		sum := f1*p1 + find(T-min(f, f1*a1)*t1, f-f1*a1)
		best = min(best, sum)
		if a1*f1 >= f {
			break
		}
	}

	if best == inf {
		return -1
	}
	return best
}
