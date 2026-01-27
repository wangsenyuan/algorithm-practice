package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var x, y, l, r int
	fmt.Fscan(reader, &x, &y, &l, &r)
	res := solve(x, y, l, r)
	fmt.Println(res)
}

func solve(x int, y int, l int, r int) int {
	if r == 1 {
		return 1
	}
	var a []int
	for i := 1; ; i *= x {
		a = append(a, i)
		if i > r/x || i*x > r {
			break
		}
	}
	var b []int
	for i := 1; ; i *= y {
		b = append(b, i)
		if i > r/y || i*y > r {
			break
		}
	}
	// len(a) & len(b) <= 60
	var arr []int
	for _, u := range a {
		for _, v := range b {
			if u+v <= r {
				arr = append(arr, u+v)
			}
		}
	}
	slices.Sort(arr)
	arr = slices.Compact(arr)
	var best int
	for i := 0; i < len(arr); i++ {
		u := max(arr[i]+1, l)
		v := r + 1
		if i+1 < len(arr) {
			v = arr[i+1]
		}
		best = max(best, v-u)
	}
	return best
}
