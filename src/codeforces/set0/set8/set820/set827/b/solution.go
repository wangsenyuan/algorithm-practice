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

	var n, k int
	fmt.Fscan(reader, &n, &k)
	dist, res := solve(n, k)
	fmt.Fprintln(writer, dist)
	for _, cur := range res {
		fmt.Fprintln(writer, cur[0], cur[1])
	}
}

func solve(n int, k int) (int, [][]int) {
	// make n as the root
	x, y := (n-1)/k, (n-1)%k

	dist := 2 * x
	if y > 1 {
		dist += 2
	} else if y > 0 {
		dist++
	}

	var res [][]int
	for i := 0; i < n-1; {
		res = append(res, []int{n, i + 1})
		j := i + x - 1
		if y > 0 {
			// 这里需要x+1个节点
			j++
			y--
		}

		i1 := j

		for j > i {
			res = append(res, []int{j, j + 1})
			j--
		}

		i = i1 + 1
	}
	return dist, res
}
