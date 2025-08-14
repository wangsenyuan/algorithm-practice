package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (a [][]int, k int, res []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m, &k)
	a = make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	res = solve(a, k)
	return
}

type pair struct {
	first  int
	second int
}

func last[T any](arr []T) T {
	return arr[len(arr)-1]
}

func solve(a [][]int, k int) []int {
	n := len(a)
	m := len(a[0])
	stack := make([][]pair, m)

	add := func(i int) {
		for j := 0; j < m; j++ {
			for len(stack[j]) > 0 && last(stack[j]).first <= a[i][j] {
				stack[j] = stack[j][:len(stack[j])-1]
			}
			stack[j] = append(stack[j], pair{a[i][j], i})
		}
	}

	count := func() int {
		var sum int
		for j := range m {
			sum += stack[j][0].first
		}
		return sum
	}

	ans := make([]int, m+1)

	update := func(best int) {
		for j := range m {
			ans[j] = stack[j][0].first
		}
		ans[m] = best
	}

	for l, r := 0, 0; r < n; r++ {
		add(r)
		for l <= r && count() > k {
			for j := 0; j < m; j++ {
				if stack[j][0].second == l {
					stack[j] = stack[j][1:]
				}
			}
			l++
		}
		if r-l+1 > ans[m] {
			update(r - l + 1)
		}
	}
	return ans[:m]
}
