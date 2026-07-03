package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	if drive(reader) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func drive(reader *bufio.Reader) bool {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) bool {
	type data struct {
		pos int
		val int
	}

	todo := make([][]data, k)
	for i, v := range a {
		todo[i%k] = append(todo[i%k], data{i, v})
	}
	arr := make([]int, len(a))

	for i := range k {
		poss := make([]int, len(todo[i]))
		for j, v := range todo[i] {
			poss[j] = v.pos
		}
		slices.SortFunc(todo[i], func(x data, y data) int {
			return x.val - y.val
		})

		for j, v := range todo[i] {
			arr[poss[j]] = v.val
		}
	}

	return sort.IntsAreSorted(arr)
}
