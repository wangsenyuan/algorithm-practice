package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func process(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	queries := make([][]int, m)
	for i := range m {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		queries[i] = []int{x, y}
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) []int {

	var mem [200][200]int

	var dfs func(x int, y int)
	dfs = func(x int, y int) {
		if mem[x][y] < 4 {
			return
		}
		tmp := mem[x][y]
		mem[x][y] %= 4
		if x == 1 {
			mem[x-1][y] += tmp / 4 * 2
			dfs(x-1, y)
		} else if x > 0 {
			mem[x-1][y] += tmp / 4
			dfs(x-1, y)
		}
		if y == 1 {
			mem[x][y-1] += tmp / 4 * 2
			dfs(x, y-1)
		} else if y > 0 {
			mem[x][y-1] += tmp / 4
			dfs(x, y-1)
		}
		mem[x+1][y] += tmp / 4
		dfs(x+1, y)
		mem[x][y+1] += tmp / 4
		dfs(x, y+1)
	}

	mem[0][0] = n
	dfs(0, 0)

	ans := make([]int, len(queries))

	for i, cur := range queries {
		x, y := cur[0], cur[1]
		x = abs(x)
		y = abs(y)
		if x >= 200 || y >= 200 {
			ans[i] = 0
		} else {
			ans[i] = mem[x][y]
		}
	}

	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
