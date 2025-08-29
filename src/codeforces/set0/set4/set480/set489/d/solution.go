package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	roads := make([][]int, m)
	for i := range m {
		roads[i] = make([]int, 2)
		fmt.Fscan(reader, &roads[i][0], &roads[i][1])
	}
	return solve(n, roads)
}

func solve(n int, roads [][]int) int {
	// m := len(roads)
	g1 := make([][]int, n)
	g2 := make([][]int, n)
	for _, road := range roads {
		a, b := road[0]-1, road[1]-1
		g1[a] = append(g1[a], b)
		g2[b] = append(g2[b], a)
	}
	// 通过b
	for i := range n {
		sort.Ints(g1[i])
		sort.Ints(g2[i])
	}

	var res int

	for a := range n {
		for c := range n {
			if a != c {
				var cnt int
				for i, j := 0, 0; i < len(g1[a]) && j < len(g2[c]); {
					if g1[a][i] == g2[c][j] {
						cnt++
						i++
						j++
					} else if g1[a][i] < g2[c][j] {
						i++
					} else {
						j++
					}
				}
				res += cnt * (cnt - 1) / 2
			}
		}
	}

	return res
}
