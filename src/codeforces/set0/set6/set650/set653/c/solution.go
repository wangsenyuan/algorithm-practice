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
	var n int
	fmt.Fscan(reader, &n)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i])
	}
	return solve(t)
}

func solve(t []int) int {
	n := len(t)

	check := func(i int) int {
		if i < 0 || i >= n-1 {
			return 1
		}
		if i&1 == 0 && t[i] >= t[i+1] {
			return 0
		}
		if i&1 == 1 && t[i] <= t[i+1] {
			return 0
		}
		return 1
	}

	var bad []int

	for i := 0; i+1 < n; i++ {
		if check(i) == 0 {
			bad = append(bad, i)
		}
	}

	if len(bad) > 4 {
		return 0
	}

	play := func(p1 int, p2 int) bool {
		if p2 >= n {
			return false
		}
		t[p1], t[p2] = t[p2], t[p1]

		defer func() {
			t[p1], t[p2] = t[p2], t[p1]
		}()

		for _, i := range bad {
			if check(i) == 0 {
				return false
			}
		}

		return check(p1)+check(p1-1)+check(p2)+check(p2-1) == 4
	}

	p := bad[0]

	var res int
	for i := 0; i < n; i++ {
		if i != p && play(i, p) {
			res++
		}
		if i != p+1 && play(i, p+1) {
			res++
		}
	}

	if play(p, p+1) {
		res--
	}

	return res
}
