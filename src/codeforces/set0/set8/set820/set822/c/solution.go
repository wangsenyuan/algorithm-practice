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
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, x int
	fmt.Fscan(reader, &n, &x)
	vouchers := make([][]int, n)
	for i := 0; i < n; i++ {
		var l, r, cost int
		fmt.Fscan(reader, &l, &r, &cost)
		vouchers[i] = []int{l, r, cost}
	}
	return solve(x, vouchers)
}

const inf = 1 << 60

type data struct {
	cost int
	r    int
}

func solve(x int, vouchers [][]int) int {
	arr := make([][]data, x+1)

	slices.SortFunc(vouchers, func(x []int, y []int) int {
		return x[1] - y[1]
	})

	best := inf

	for _, cur := range vouchers {
		l, r, cost := cur[0], cur[1], cur[2]
		x1 := r - l + 1
		if x1 >= x {
			continue
		}
		// x1 < x
		x2 := x - x1
		if len(arr[x2]) > 0 {
			j := sort.Search(len(arr[x2]), func(j int) bool {
				return arr[x2][j].r >= l
			})
			if j > 0 {
				best = min(best, arr[x2][j-1].cost+cost)
			}
		}
		if len(arr[x1]) == 0 || arr[x1][len(arr[x1])-1].cost > cost {
			arr[x1] = append(arr[x1], data{cost, r})
		}
	}

	if best == inf {
		return -1
	}

	return best
}
