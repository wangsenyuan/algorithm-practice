package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	y := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &y[i])
	}
	return solve(y)
}

func solve(y []int) bool {
	n := len(y)

	get := func(i int, j int) pair {
		// first >= 0
		dy := y[j] - y[i]
		dx := j - i
		g := gcd(abs(dy), dx)
		dy /= g
		dx /= g
		return pair{dy, dx}
	}

	slops := make([]pair, n)

	for i := 1; i < n; i++ {
		slops[i] = get(0, i)
	}

	marked := make([]bool, n)

	for i := 1; i < n; i++ {
		if marked[i] {
			// 如果它和前面的某个点已经共线了，那么它已经被处理了
			continue
		}
		// 这是一个新的斜率
		var cnt int
		for j := 1; j < n; j++ {
			if slops[j] == slops[i] {
				marked[j] = true
				cnt++
			}
		}
		if cnt == n-1 {
			// 全部在一条线上
			return false
		}
		// 不在一条上线
		var first = -1
		var w pair
		ok := true
		for j := 1; j < n; j++ {
			if slops[j] != slops[i] {
				if first < 0 {
					first = j
					continue
				}

				cur := get(first, j)

				if w.second == 0 {
					w = cur
					continue
				}

				if cur != w {
					ok = false
					break
				}
			}
		}

		if ok && (w.second == 0 || w == slops[i]) {
			return true
		}
	}

	w := get(1, 2)
	// 还有就是第一个是单独的，其他的所有的点在一条线上
	for j := 3; j < n; j++ {
		if get(1, j) != w {
			return false
		}
	}

	return true
}

type pair struct {
	first  int
	second int
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}

	return a
}

func sign(num int) int {
	if num > 0 {
		return 1
	}
	if num < 0 {
		return -1
	}
	return 0
}

func abs(num int) int {
	return max(num, -num)
}
