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
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	points := make([][]int, n)
	for i := range n {
		points[i] = make([]int, 2)
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}
	return solve(points)
}

const inf = 1 << 60

func solve(points [][]int) int {
	var pos []int
	pos = append(pos, 0)
	for _, p := range points {
		d := max(p[0], p[1])
		pos = append(pos, d)
	}

	slices.Sort(pos)
	pos = slices.Compact(pos)

	m := len(pos)
	levels := make([][][]int, m)
	levels[0] = append(levels[0], []int{0, 0})
	for _, p := range points {
		d := max(p[0], p[1])
		i := sort.SearchInts(pos, d)
		levels[i] = append(levels[i], p)
	}

	for i := range m {
		slices.SortFunc(levels[i], func(a []int, b []int) int {
			return cmp.Or(a[0]-b[0], b[1]-a[1])
		})
	}

	dp := make([]int, 2)
	ndp := make([]int, 2)
	ndp[0] = inf
	ndp[1] = inf

	distance := func(arr [][]int) int {
		first := arr[0]
		second := last(arr)
		return abs(first[0]-second[0]) + abs(first[1]-second[1])
	}

	for i := 1; i < m; i++ {
		w := distance(levels[i])
		for d := range 2 {
			// d = 0, 表示在上一层的开始节点,
			// d = 1, 表示在上一层的结束节点
			px, py := levels[i-1][0][0], levels[i-1][0][1]
			if d == 1 {
				tmp := last(levels[i-1])
				px, py = tmp[0], tmp[1]
			}
			for d1 := range 2 {
				nx, ny := levels[i][0][0], levels[i][0][1]
				if d1 == 1 {
					tmp := last(levels[i])
					nx, ny = tmp[0], tmp[1]
				}
				v := abs(px-nx) + abs(py-ny)
				// 先到nx, ny, 然后从这层到达另外一端
				ndp[d1^1] = min(ndp[d1^1], dp[d]+v+w)
			}
		}
		for d := range 2 {
			dp[d] = ndp[d]
			ndp[d] = inf
		}
	}

	return min(dp[0], dp[1])
}

func last(arr [][]int) []int {
	return arr[len(arr)-1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
