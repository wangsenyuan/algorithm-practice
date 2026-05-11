package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 3)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return solve(p, queries)
}

func solve(p []int, queries [][]int) []int {
	n := len(p)
	pos := make([]int, n+1)
	for i, v := range p {
		pos[v] = i + 1
	}

	var f func(l int, r int, exp int, lt *int, lt1 *int, gt *int, gt1 *int)

	f = func(l int, r int, exp int, lt *int, lt1 *int, gt *int, gt1 *int) {
		mid := (l + r) / 2
		if mid == exp {
			return
		}
		if mid < exp {
			if p[mid-1] > p[exp-1] {
				*lt++
			} else {
				*lt1++
			}
			f(mid+1, r, exp, lt, lt1, gt, gt1)
		} else {
			// mid > exp
			if p[mid-1] < p[exp-1] {
				*gt++
			} else {
				*gt1++
			}
			f(l, mid-1, exp, lt, lt1, gt, gt1)
		}
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		l, r, k := cur[0], cur[1], cur[2]
		exp := pos[k]
		if exp < l || r < exp {
			ans[i] = -1
		} else {
			var lt, gt, lt1, gt1 int
			f(l, r, exp, &lt, &lt1, &gt, &gt1)
			// lt是要交换进来比k小的数, lt1是已经比k小的数
			if lt+lt1 >= k || gt+gt1 > n-k {
				ans[i] = -1
			} else {
				ans[i] = lt + gt + abs(lt-gt)
			}
		}
	}

	return ans
}

func abs(a int) int {
	return max(a, -a)
}
