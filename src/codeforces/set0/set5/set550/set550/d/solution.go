package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var k int
	fmt.Fscan(reader, &k)
	n, res := solve(k)
	if n == 0 {
		fmt.Println("NO")
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, "YES")
	fmt.Fprintln(writer, n, len(res))
	for _, edge := range res {
		fmt.Fprintln(writer, edge[0], edge[1])
	}
}

func solve(k int) (n int, res [][]int) {
	if k&1 == 0 {
		return 0, nil
	}

	if k == 1 {
		n = 2
		res = append(res, []int{1, 2})
		return
	}

	n1, res1 := play(k)

	_, res2 := play(k)

	for i := range res2 {
		for j := range res2[i] {
			res2[i][j] += n1
		}
	}

	n = 2 * n1

	res = append(res1, res2...)
	res = append(res, []int{1, n1 + 1})

	return
}

func play(k int) (n int, res [][]int) {
	// 1, k - 1, k - 1
	n = 2*k - 1
	for i := 2; i <= k; i++ {
		res = append(res, []int{1, i})
	}
	for i := 2; i <= k; i++ {
		for j := k + 1; j <= n; j++ {
			res = append(res, []int{i, j})
		}
	}

	for j := k + 1; j <= n; j += 2 {
		res = append(res, []int{j, j + 1})
	}

	return
}
