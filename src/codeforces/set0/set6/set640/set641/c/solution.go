package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	qs := make([][]int, q)
	for i := 0; i < q; i++ {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			var x int
			fmt.Fscan(reader, &x)
			qs[i] = []int{t, x}
		} else {
			qs[i] = []int{t}
		}
	}
	return solve(n, qs)
}

func solve(n int, q [][]int) []int {
	moves := []int{0, 0}
	var swap int

	add := func(x *int, y int) {
		if y < 0 {
			y += n
		}
		*x += y
		if *x >= n {
			*x -= n
		}
	}

	for _, cur := range q {
		if cur[0] == 1 {
			x := cur[1]
			add(&moves[0], x)
			add(&moves[1], x)
			if x&1 == 1 {
				moves[0], moves[1] = moves[1], moves[0]
				swap ^= 1
			}
		} else {
			// swap
			swap ^= 1
			x, y := moves[0], moves[1]
			add(&x, 1)
			add(&y, -1)
			moves[1] = x
			moves[0] = y
		}
	}

	ans := make([]int, n)

	for i := range n {
		j := i & 1
		j ^= swap
		ni := i
		add(&ni, moves[j])
		ans[ni] = i + 1
	}

	return ans
}
