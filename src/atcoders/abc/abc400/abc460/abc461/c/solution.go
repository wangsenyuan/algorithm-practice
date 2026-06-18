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
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int64 {
	var n, k, m int
	fmt.Fscan(reader, &n, &k, &m)
	gems := make([]gem, n)
	for i := range n {
		fmt.Fscan(reader, &gems[i].color, &gems[i].value)
	}
	return solve(n, k, m, gems)
}

type gem struct {
	color int
	value int
}

func solve(n, k, m int, gems []gem) int64 {
	// TODO
	best := make([]int, n)
	for i := range n {
		best[i] = -1
	}

	for i, gem := range gems {
		c := gem.color - 1
		if best[c] < 0 || gem.value > gems[best[c]].value {
			best[c] = i
		}
	}

	slices.SortFunc(best, func(x int, y int) int {
		if x < 0 && y < 0 {
			return 0
		}
		if x < 0 {
			// move x to end
			return 1
		}
		if y < 0 {
			return -1
		}
		return cmp.Or(gems[y].value-gems[x].value, x-y)
	})

	marked := make([]bool, n)
	var sum int
	for i := range m {
		sum += gems[best[i]].value
		marked[best[i]] = true
	}

	var rest []int
	for i, gem := range gems {
		if !marked[i] {
			rest = append(rest, gem.value)
		}
	}

	slices.Sort(rest)
	slices.Reverse(rest)
	for i := range k - m {
		sum += rest[i]
	}

	return int64(sum)
}
