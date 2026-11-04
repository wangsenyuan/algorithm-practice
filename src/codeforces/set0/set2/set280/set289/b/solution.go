package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, m, d int
	fmt.Fscan(reader, &n, &m, &d)
	a := make([]int, 0, n*m)
	for range n * m {
		var x int
		fmt.Fscan(reader, &x)
		a = append(a, x)
	}
	return solve(a, d)
}

func solve(a []int, d int) int {
	// 假设最后的结果是x
	// a[i][j] % d = x % d
	// 所以对于所有格子的差 % d = 0
	n := len(a)
	for i := range n - 1 {
		diff := abs(a[0] - a[i+1])
		if diff%d != 0 {
			return -1
		}
	}
	// 可以, 那么现在的问题就是应该是多少呢?
	// 如果x太大, 有很多添加, 如果x太小, 有很多-
	play := func(x int) int {
		var sum int
		for _, v := range a {
			sum += abs(v-x) / d
		}
		return sum
	}
	x0 := slices.Min(a)
	x1 := slices.Max(a)

	lo, hi := 0, (x1-x0)/d+1
	for hi-lo > 2 {
		// lo, mid, hi
		diff := (hi - lo + 1) / 3
		m1 := lo + diff
		m2 := hi - diff
		w1 := play(m1*d + x0)
		w2 := play(m2*d + x0)
		if w1 >= w2 {
			lo = m1
		} else {
			hi = m2
		}
	}

	ans := 1 << 30

	for i := lo; i <= hi; i++ {
		ans = min(ans, play(i*d+x0))
	}

	return ans
}

func abs(num int) int {
	return max(num, -num)
}
