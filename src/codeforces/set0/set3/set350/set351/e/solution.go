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
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

type num struct {
	id  int
	val int
}

func solve(a []int) int {
	n := len(a)
	arr := make([]num, n)
	for i := range n {
		arr[i] = num{i, a[i]}
	}

	slices.SortFunc(arr, func(x, y num) int {
		return cmp.Or(abs(x.val)-abs(y.val), x.id-y.id)
	})

	aa := make([]int, n)
	ok := make([]bool, n)

	check := func(v int, p int) int {
		var res int
		for i := 0; i < p; i++ {
			if ok[i] && aa[i] > v {
				res++
			}
		}
		for i := p + 1; i < n; i++ {
			if ok[i] && v > aa[i] {
				res++
			}
		}
		return res
	}

	var res int

	for _, cur := range arr {
		i := cur.id

		w1 := check(cur.val, i)
		w2 := check(cur.val*-1, i)
		if w1 < w2 {
			aa[i] = cur.val
		} else {
			aa[i] = -cur.val
		}
		res += min(w1, w2)
		ok[i] = true
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}
