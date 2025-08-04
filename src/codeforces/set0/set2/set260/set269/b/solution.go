package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	plants := make([]int, n)
	pos := make([]float64, n)
	for i := range n {
		fmt.Fscan(reader, &plants[i], &pos[i])
	}
	return solve(m, plants, pos)
}

func solve(m int, plants []int, pos []float64) int {
	n := len(plants)

	stack := make([]int, n)
	var top int
	for _, v := range plants {
		i := sort.Search(top, func(i int) bool {
			return stack[i] > v
		})
		stack[i] = v
		if i == top {
			top++
		}
	}
	return n - top
}
