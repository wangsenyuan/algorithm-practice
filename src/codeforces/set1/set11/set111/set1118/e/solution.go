package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)
	res := solve(n, k)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, "YES")
	for _, edge := range res {
		fmt.Fprintln(writer, edge[0], edge[1])
	}
}

func solve(n int, k int) [][]int {
	t := k * (k - 1)
	if t < n {
		return nil
	}
	var res [][]int
	id := 1
	for len(res) < n {
		for i := id + 1; i <= k && len(res) < n; i++ {
			res = append(res, []int{id, i})
			if len(res) < n {
				res = append(res, []int{i, id})
			}
		}
		id++
	}

	return res
}
