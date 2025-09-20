package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var d, n, m int
	fmt.Fscan(reader, &d, &n, &m)
	stations := make([][]int, m)
	for i := range stations {
		stations[i] = make([]int, 2)
		fmt.Fscan(reader, &stations[i][0], &stations[i][1])
	}
	return solve(d, n, stations)
}

const inf = 1 << 60

func solve(d int, n int, stations [][]int) int {
	stations = append(stations, []int{0, inf}, []int{d, inf})

	slices.SortFunc(stations, func(a, b []int) int {
		return a[0] - b[0]
	})

	var dfs func(gas int, l int, r int) []int

	dfs = func(gas int, l int, r int) []int {
		dist := stations[r][0] - stations[l][0]
		if dist <= gas {
			return []int{0, gas - dist}
		}
		if l+1 == r {
			// 中间没有加油的地方
			return nil
		}
		best := l + 1
		for i := l + 1; i < r; i++ {
			if stations[i][1] < stations[best][1] {
				best = i
			}
		}
		// 先到达位置best，然后再到位置r
		tmp := dfs(gas, l, best)
		if len(tmp) == 0 {
			return nil
		}
		cost := tmp[0]
		// 现在要在best处加油
		need := min(stations[r][0]-stations[best][0], n)
		add := max(0, need-tmp[1])
		cost += add * stations[best][1]

		tmp1 := dfs(tmp[1]+add, best, r)
		if len(tmp1) == 0 {
			return nil
		}
		cost += tmp1[0]
		return []int{cost, tmp1[1]}
	}

	res := dfs(n, 0, len(stations)-1)

	if len(res) == 0 {
		return -1
	}
	return res[0]
}
