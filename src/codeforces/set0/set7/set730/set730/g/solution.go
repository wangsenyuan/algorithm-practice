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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, cur := range res {
		fmt.Fprintln(writer, cur[0], cur[1])
	}
}

func drive(reader *bufio.Reader) [][]int {
	var n int
	fmt.Fscan(reader, &n)
	requests := make([][]int, n)
	for i := range n {
		requests[i] = make([]int, 2)
		fmt.Fscan(reader, &requests[i][0], &requests[i][1])
	}
	return solve(requests)
}

type pair struct {
	first  int
	second int
}

func solve(requests [][]int) [][]int {
	n := len(requests)

	var res [][]int

	var arr []pair

	add := func(s int, e int) {
		res = append(res, []int{s, e})
		arr = append(arr, pair{s, e})
	}

	for i := range n {
		s, d := requests[i][0], requests[i][1]
		// 检查是否可以从s到d
		e := s + d - 1
		ok := true
		for j := range i {
			if max(res[j][0], s) <= min(res[j][1], e) {
				ok = false
				break
			}
		}
		if ok {
			add(s, e)
			continue
		}

		slices.SortFunc(arr, func(a, b pair) int {
			return a.first - b.first
		})

		prev := 1

		for _, cur := range arr {
			if prev+d-1 < cur.first {
				add(prev, prev+d-1)
				ok = true
				break
			}
			prev = cur.second + 1
		}

		if !ok {
			add(prev, prev+d-1)
		}
	}

	return res
}
