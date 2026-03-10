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
		var n int
		fmt.Fscan(reader, &n)
		res := solve(n)
		fmt.Fprintln(writer, len(res))
		for _, cur := range res {
			fmt.Fprintln(writer, cur[0], cur[1], cur[2])
		}
	}
}

func solve(n int) [][]int {
	var res [][]int
	// 第一个需要反转掉
	res = append(res, []int{1, 1, n})
	for r := 2; r <= n; r++ {
		k := r - 1
		// shift(k)
		if k > 1 {
			res = append(res, []int{r, 1, k})
		}
		if n-k > 1 {
			res = append(res, []int{r, k + 1, n})
		}
	}

	return res
}
