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
	_, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, cur := range res {
		writer.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
	}
}

func drive(reader *bufio.Reader) (n int, edges [][]int, res [][]int) {
	var m int
	fmt.Fscan(reader, &n, &m)
	edges = make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	res = solve(n, edges)
	return
}

type Edge struct {
	id     int
	weight int
	used   int
}

func solve(n int, edges [][]int) [][]int {
	m := len(edges)
	arr := make([]Edge, m)
	for i, cur := range edges {
		arr[i] = Edge{i, cur[0], cur[1]}
	}
	slices.SortFunc(arr, func(a, b Edge) int {
		return cmp.Or(a.weight-b.weight, b.used-a.used)
	})

	res := make([][]int, m)
	var que [][]int
	var k int
	for j, cur := range arr {
		if cur.used == 1 {
			res[cur.id] = []int{k + 1, k + 2}
			k++
			for i := k - 2; i >= 0 && len(que) < m-j; i-- {
				que = append(que, []int{i + 1, k + 1})
			}
		} else {
			if len(que) == 0 {
				return nil
			}
			res[cur.id] = que[0]
			que = que[1:]
		}
	}

	return res
}
