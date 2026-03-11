package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, len(res))
	for _, v := range res {
		fmt.Fprintln(writer, v[0], v[1])
	}
}

func drive(reader *bufio.Reader) (n int, m int, res [][]int) {
	fmt.Fscan(reader, &n, &m)
	res = solve(n, m)
	return
}

func solve(n int, m int) [][]int {
	k := min(n, m)
	res := make([][]int, 0, k+1)
	for i := 0; i <= k; i++ {
		res = append(res, []int{i, k - i})
	}
	return res
}

func bruteForce(n int, m int) [][]int {
	var swap bool
	if n > m {
		swap = true
		n, m = m, n
	}
	// n <= m
	var res [][]int

	update := func(set [][]int) {
		if len(set) > len(res) {
			res = set
		}
	}

	check := func(set [][]int, x int, y int) bool {
		for _, p := range set {
			dx := x - p[0]
			dy := y - p[1]
			d := dx*dx + dy*dy
			d1 := int(math.Sqrt(float64(d)))
			if d1*d1 == d {
				return false
			}
		}
		return true
	}

	var f func(set [][]int, x int)
	f = func(set [][]int, x int) {
		if x > n {
			update(set)
			return
		}
		// y <= n
		var s int
		if x == 0 {
			s = 1
		}
		for y := s; y <= m; y++ {
			if check(set, x, y) {
				set = append(set, []int{x, y})
				f(set, x+1)
				set = set[:len(set)-1]
			}
		}
	}

	f(nil, 0)

	if swap {
		for i := range res {
			res[i][0], res[i][1] = res[i][1], res[i][0]
		}
	}

	return res
}
