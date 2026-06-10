package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		_, res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) (a []int, ord []int) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	ord = solve(a)
	return
}

type lecturer struct {
	id  int
	val int
}

func solve(a []int) []int {
	n := len(a)
	arr := make([]lecturer, n)
	for i := range n {
		arr[i] = lecturer{i, a[i]}
	}

	slices.SortFunc(arr, func(a, b lecturer) int {
		return b.val - a.val
	})

	res := make([]int, n)
	var pos int
	var marker int
	for l, r := 0, n-1; l <= r; r-- {

		// marker < arr[r].val
		if l < r && marker < arr[r].val {
			for l < r && marker < arr[r].val {
				res[pos] = arr[l].id + 1
				pos++
				marker++
				l++
			}
		}
		res[pos] = arr[r].id + 1
		pos++
		marker = 1
	}

	return res
}
