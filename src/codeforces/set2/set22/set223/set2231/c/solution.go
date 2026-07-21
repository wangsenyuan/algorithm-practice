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

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	res := make([]int, t)
	for i := range t {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[j])
		}
		res[i] = solve(a)
	}
	return res
}

func solve(a []int) int {
	x := slices.Max(a)
	if x == 1 {
		return 0
	}

	freq := make(map[int]int)
	sum := make(map[int]int)

	for _, v := range a {
		if v == 1 {
			freq[2]++
			sum[2]++
		}
		var ops int
		for v > 1 {
			freq[v]++
			sum[v] += ops
			if v&1 == 1 {
				v += 1
			} else {
				v /= 2
			}
			ops++
		}
		freq[1]++
		sum[1] += ops
	}

	n := len(a)

	res := sum[1]
	for v := x; v > 1; {
		if freq[v] == n {
			res = min(res, sum[v])
		}
		if v&1 == 1 {
			v++
		} else {
			v /= 2
		}
	}

	return res
}
