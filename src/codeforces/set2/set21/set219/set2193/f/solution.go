package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, ax, ay, bx, by int
	fmt.Fscan(reader, &n, &ax, &ay, &bx, &by)
	a := []int{ax, ay}
	b := []int{bx, by}
	customers := make([][]int, n)
	for i := range n {
		customers[i] = make([]int, 2)
		fmt.Fscan(reader, &customers[i][0])
	}
	for i := range n {
		fmt.Fscan(reader, &customers[i][1])
	}
	return solve(a, b, customers)
}

func solve(a []int, b []int, customers [][]int) int {
	slices.SortFunc(customers, func(f []int, s []int) int {
		return cmp.Or(f[0]-s[0], f[1]-s[1])
	})

	// dp[0]
	dp := make([]int, 2)
	last := [][]int{a, a}
	for i := 0; i < len(customers); {
		j := i
		for i < len(customers) && customers[i][0] == customers[j][0] {
			i++
		}
		cur := [][]int{customers[j], customers[i-1]}
		move := cur[1][1] - cur[0][1]

		ndp := []int{inf, inf}
		for k1 := range 2 {
			for k2 := range 2 {
				cur := dp[k1] + distance(last[k1], cur[k2]) + move
				ndp[k2^1] = min(ndp[k2^1], cur)
			}
		}
		dp = ndp
		last = cur
	}

	res := dp[0] + distance(last[0], b)
	res = min(res, dp[1]+distance(last[1], b))

	return res
}

func distance(a []int, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

const inf = 1 << 60
