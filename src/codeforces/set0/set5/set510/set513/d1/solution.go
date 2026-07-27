package main

import (
	"bufio"
	"fmt"
	"os"
)

type constraint struct {
	a         int
	b         int
	direction string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	if res == nil {
		fmt.Fprintln(writer, "IMPOSSIBLE")
		return
	}
	for i, label := range res {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, label)
	}
	fmt.Fprintln(writer)
}

func drive(reader *bufio.Reader) []int {
	var n, c int
	fmt.Fscan(reader, &n, &c)
	constraints := make([]constraint, c)
	for i := range constraints {
		fmt.Fscan(
			reader,
			&constraints[i].a,
			&constraints[i].b,
			&constraints[i].direction,
		)
	}
	return solve(n, constraints)
}

func solve(n int, constraints []constraint) []int {
	dp := make([][]int, n)

	for i := range n {
		dp[i] = make([]int, n)
		for j := range n {
			dp[i][j] = -2
		}
	}

	maxLeft := make([]int, n)
	minRight := make([]int, n)
	maxRequired := make([]int, n)
	invalid := make([]bool, n)
	for i := range n {
		maxLeft[i] = i
		minRight[i] = n
		maxRequired[i] = i
	}

	for _, cur := range constraints {
		a, b := cur.a-1, cur.b-1
		if b <= a {
			invalid[a] = true
			continue
		}
		maxRequired[a] = max(maxRequired[a], b)
		if cur.direction == "LEFT" {
			maxLeft[a] = max(maxLeft[a], b)
		} else {
			minRight[a] = min(minRight[a], b)
		}
	}

	var f func(l int, r int) bool
	f = func(l int, r int) bool {
		if dp[l][r] != -2 {
			return dp[l][r] >= l
		}
		dp[l][r] = -1
		if invalid[l] || maxRequired[l] > r {
			return false
		}

		// [l+1...split] is the left subtree, and
		// [split+1...r] is the right subtree.
		low := max(l, maxLeft[l])
		high := min(r, minRight[l]-1)
		for split := low; split <= high; split++ {
			leftOK := split == l || f(l+1, split)
			if !leftOK {
				continue
			}
			rightOK := split == r || f(split+1, r)
			if rightOK {
				dp[l][r] = split
				return true
			}
		}
		return false
	}

	ok := f(0, n-1)
	if !ok {
		return nil
	}

	var res []int
	var output func(l int, r int)

	output = func(l int, r int) {
		if l > r {
			return
		}
		split := dp[l][r]
		output(l+1, split)
		res = append(res, l+1)
		output(split+1, r)
	}

	output(0, n-1)

	return res
}
