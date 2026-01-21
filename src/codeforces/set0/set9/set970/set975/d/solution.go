package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, a, b int
	fmt.Fscan(reader, &n, &a, &b)
	ghosts := make([][]int, n)
	for i := range n {
		ghosts[i] = make([]int, 3)
		fmt.Fscan(reader, &ghosts[i][0], &ghosts[i][1], &ghosts[i][2])
	}
	return solve(a, b, ghosts)
}

type pair struct {
	first  int
	second int
}

func solve(a int, b int, ghosts [][]int) int {
	// n := len(ghosts)
	freq := make(map[int]int)
	para := make(map[pair]int)
	var ans int
	for _, cur := range ghosts {
		vx, vy := cur[1], cur[2]
		d := a*vx - vy
		ans += freq[d] - para[pair{vx, vy}]
		freq[d]++
		para[pair{vx, vy}]++
	}

	return ans * 2
}
