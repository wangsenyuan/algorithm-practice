package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	birds := make([][]int, m)
	for i := range m {
		birds[i] = make([]int, 2)
		fmt.Fscan(reader, &birds[i][0], &birds[i][1])
	}
	return solve(n, birds)
}

type pair struct {
	first  int
	second int
}

func solve(n int, birds [][]int) int {
	m := len(birds)

	freq := make([]map[pair]int, n+1)
	for i := range n + 1 {
		freq[i] = make(map[pair]int)
	}

	for i := range m {
		for j := range i {
			x := play(birds[i], birds[j])
			if x < 1 || x > n {
				continue
			}

			if birds[i][0] == x {
				freq[x][pair{0, 1}]++
			} else {
				dx := birds[i][0] - birds[j][0]
				dy := birds[i][1] - birds[j][1]
				g := gcd(abs(dx), abs(dy))
				dx /= g
				dy /= g
				if sign(dx) != sign(dy) {
					freq[x][pair{-abs(dx), abs(dy)}]++
				} else {
					freq[x][pair{abs(dx), abs(dy)}]++
				}
			}

		}
	}

	best := make([]int, m*(m-1)/2+1)

	for k := 1; k <= m; k++ {
		sum := k * (k - 1) / 2
		best[sum] = k
	}

	var res int

	for i := range n {
		var tmp int
		for _, v := range freq[i+1] {
			tmp = max(tmp, v)
		}
		res += best[tmp]
	}

	return res
}

func play(a []int, b []int) int {
	if a[1] == b[1] {
		// 水平线
		return 0
	}
	if a[1] > b[1] {
		a, b = b, a
	}
	// a[1] < b[1]
	dy := b[1] - a[1]
	dx := b[0] - a[0]
	if dx == 0 {
		// 垂直线
		return a[0]
	}
	// y = k * x + b
	// b[1] = k * b[0] + b
	// a[1] = k * a[0] + b
	// k = dy / dx
	// b = b[1] - k * b[0] = (b[1] * dx - dy * b[0]) / dx
	w := b[1]*dx - dy*b[0]
	// v := dx
	// b = w / v
	// 当 y = 0时, x = -b/k = -w/(v * k) = -w / dy
	if w > 0 || (-w)%dy != 0 {
		return -1
	}
	return -w / dy
}

func abs(num int) int {
	return max(num, -num)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func sign(num int) int {
	if num < 0 {
		return -1
	}
	if num == 0 {
		return 0
	}
	return 1
}
