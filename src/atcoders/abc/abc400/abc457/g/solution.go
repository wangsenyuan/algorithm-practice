package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	apples := make([][]int, n)
	for i := range n {
		apples[i] = make([]int, 2)
		fmt.Fscan(reader, &apples[i][0], &apples[i][1])
	}
	return solve(apples)
}

func solve(apples [][]int) int {
	slices.SortFunc(apples, func(a []int, b []int) int {
		return cmp.Or(a[0]+a[1]-(b[0]+b[1]), (b[1]-b[0])-(a[1]-a[0]))
	})

	n := len(apples)
	stack := make([]int, n)
	var top int
	for _, cur := range apples {
		v := cur[1] - cur[0]
		i := sort.Search(top, func(i int) bool {
			return stack[i] >= v
		})
		stack[i] = v
		if i == top {
			top++
		}
	}
	return top
}
