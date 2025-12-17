package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	q := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &q[i])
	}
	s := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &s[i])
	}
	return solve(k, q, s)
}

func solve(k int, q []int, s []int) bool {
	n := len(q)
	a := make([]int, n)
	for i := range n {
		q[i]--
		s[i]--
		a[i] = i
	}

	apply := func(a []int) []int {
		res := slices.Clone(a)
		for i := range n {
			res[i] = a[q[i]]
		}
		return res
	}

	inv := func(a []int) []int {
		res := slices.Clone(a)
		for i := range n {
			res[q[i]] = a[i]
		}
		return res
	}

	check := func(a []int) bool {
		for i := range n {
			if s[i] != a[i] {
				return false
			}
		}
		return true
	}

	if check(a) {
		return false
	}

	t := inv(a)
	x := check(t)
	y := check(apply(a))

	if k == 1 {
		return x || y
	}

	for i := range k {
		a = apply(a)
		if check(a) {
			if (i > 0 || !x) && (k-i-1)%2 == 0 {
				return true
			}
			break
		}
	}

	for i := range k {
		if check(t) {
			if i > 0 || !y {
				return (k-1-i)%2 == 0
			}
			return false
		}
		t = inv(t)
	}

	return false
}
