package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n, &m)
	xs := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &xs[i])
	}
	segments := make([][]int, m)
	for i := range m {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		segments[i] = []int{l, r}
	}
	ans := solve(xs, segments)
	s := fmt.Sprintf("%v", ans)
	fmt.Println(s[1 : len(s)-1])
}

func solve(xs []int, segements [][]int) []int {
	n := len(xs)
	type pair struct {
		first  int
		second int
	}
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{xs[i], i}
	}
	slices.SortFunc(arr, func(a, b pair) int {
		return a.first - b.first
	})
	ans := make([]int, n)
	for i := range n {
		ans[arr[i].second] = i % 2
	}
	return ans
}
