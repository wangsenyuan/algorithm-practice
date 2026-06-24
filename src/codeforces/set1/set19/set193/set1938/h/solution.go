package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	tables := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &tables[i])
	}
	return solve(n, tables)
}

type table struct {
	cnt0 int
	cnt1 int
}

func solve(n int, tables []string) int {
	var arr []table
	buf := make([]table, n)
	min_pos := []int{-1, -1}
	for i, cur := range tables {
		var tmp table
		for j := range cur {
			if cur[j] == '0' {
				tmp.cnt0++
			} else {
				tmp.cnt1++
			}
		}
		if tmp.cnt0 > 0 && tmp.cnt1 > 0 {
			arr = append(arr, tmp)
		} else if tmp.cnt0 == 0 {
			if min_pos[1] == -1 || buf[min_pos[1]].cnt1 > tmp.cnt1 {
				min_pos[1] = i
			}
		} else if tmp.cnt1 == 0 {
			if min_pos[0] == -1 || buf[min_pos[0]].cnt0 > tmp.cnt0 {
				min_pos[0] = i
			}
		}
		buf[i] = tmp
	}

	if len(arr) == 0 {
		return 0
	}

	const inf = 1 << 60
	dp := []int{0, inf, inf, inf}
	for _, cur := range arr {
		ndp := []int{inf, inf, inf, inf}
		for mask, v := range dp {
			if v == inf {
				continue
			}
			// mixed table keeps 0, so all 1s on it move
			ndp[mask|1] = min(ndp[mask|1], v+cur.cnt1)
			// mixed table keeps 1, so all 0s on it move
			ndp[mask|2] = min(ndp[mask|2], v+cur.cnt0)
		}
		dp = ndp
	}

	best := inf
	for mask, v := range dp {
		if v == inf {
			continue
		}
		if mask&2 > 0 && mask&1 == 0 && min_pos[0] < 0 {
			if min_pos[1] < 0 {
				continue
			}
			v += buf[min_pos[1]].cnt1
		}
		if mask&1 > 0 && mask&2 == 0 && min_pos[1] < 0 {
			if min_pos[0] < 0 {
				continue
			}
			v += buf[min_pos[0]].cnt0
		}
		best = min(best, v)
	}
	return best
}
