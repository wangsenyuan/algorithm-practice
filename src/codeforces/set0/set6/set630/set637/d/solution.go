package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("IMPOSSIBLE")
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, row := range res {
		fmt.Fprintln(writer, row)
	}
}

func drive(reader *bufio.Reader) (m int, s int, d int, obstacles []int, res []string) {
	var n int
	fmt.Fscan(reader, &n, &m, &s, &d)
	obstacles = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &obstacles[i])
	}
	res = solve(m, s, d, obstacles)
	return
}

func solve(m int, s int, d int, obstacles []int) []string {
	if d == 1 {
		// no way
		return nil
	}

	// 如果目前的位置可以jump，它的前面至少有s的距离，没有障碍物
	// 那么在接下来的位置到d之前的，都可以做为落地的地方
	// 所以需要知道哪里是一个落地的地方
	var pos []int
	pos = append(pos, 0)
	pos = append(pos, m)
	for _, x := range obstacles {
		pos = append(pos, x-1)
		pos = append(pos, x)
		pos = append(pos, x+1)
	}
	slices.Sort(pos)
	pos = slices.Compact(pos)

	slices.Sort(obstacles)
	n := len(pos)

	dp := make([]int, n+2)
	dp[0]++
	dp[1]--

	runFrom := make([]int, n)
	runFrom[0] = 0

	var lastObstacle = -1
	var r int
	for i, j := 1, 0; i < n; i++ {
		if i > 0 {
			dp[i] += dp[i-1]
		}
		next := sort.SearchInts(obstacles, pos[i])
		if next < len(obstacles) && obstacles[next] == pos[i] {
			lastObstacle = i
		}

		runFrom[i] = -1
		for j < i && (j <= lastObstacle || dp[j] == 0) {
			// 中间有block, 或者j不是一个安全的落点
			j++
		}
		if pos[i]-pos[j] >= s {
			// 可以从j开始跑, 且位置j是一个安全的落点
			runFrom[i] = j
			for r < n && pos[r] <= pos[i]+d {
				r++
			}
			// pos[i1] >= pos[i] + d
			dp[i+1]++
			dp[r]--
		}
	}

	if runFrom[n-1] == -1 && dp[n-1] == 0 {
		// no way to run to n-1 or jump to n-1
		return nil
	}

	var res []string

	nextPos := m

	for i := n - 1; i > 0; {
		for i > 0 && runFrom[i] == -1 {
			i--
		}
		if i == 0 {
			// should not happen
			return nil
		}
		// runFrom[i] >= 0
		if pos[i] < nextPos {
			res = append(res, fmt.Sprintf("JUMP %d", nextPos-pos[i]))
		}
		j := runFrom[i]
		res = append(res, fmt.Sprintf("RUN %d", pos[i]-pos[j]))
		nextPos = pos[j]
		i = j - 1
	}

	slices.Reverse(res)
	return res
}
