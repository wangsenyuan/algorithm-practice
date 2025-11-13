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
	res := drive(reader)
	if len(res) == 0 {
		fmt.Println(":(")
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

type pair struct {
	first  int
	second int
}

func solve(a []int) []int {
	n := len(a)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{a[i] + i, i}
	}

	slices.SortFunc(arr, func(x, y pair) int {
		return cmp.Or(x.first-y.first, y.second-x.second)
	})

	res := make([]int, n)

	for i := range n {
		res[i] = arr[i].first - i
	}

	if !slices.IsSorted(res) {
		return nil
	}

	return res
}
