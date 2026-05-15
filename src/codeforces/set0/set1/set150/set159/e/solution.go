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
	_, best, res := drive(reader)
	fmt.Println(best)
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (cubes [][]int, best int, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	cubes = make([][]int, n)
	for i := 0; i < n; i++ {
		cubes[i] = make([]int, 2)
		fmt.Fscan(reader, &cubes[i][0], &cubes[i][1])
	}
	best, res = solve(cubes)
	return
}

type cube struct {
	id    int
	color int
	size  int
}

func solve(cubes [][]int) (best int, res []int) {
	n := len(cubes)
	arr := make([]cube, n)
	for i := range n {
		arr[i] = cube{i + 1, cubes[i][0], cubes[i][1]}
	}

	slices.SortFunc(arr, func(a, b cube) int {
		return cmp.Or(a.color-b.color, b.size-a.size)
	})

	// 数量也有关系
	dp := make([]int, n+1)
	fp := make([]int, n+1)
	for i := range n + 1 {
		fp[i] = -1
		dp[i] = -(1 << 60)
	}

	dp[0] = 0

	pickColors := []int{-1, -1}

	for i := 0; i < n; {
		j := i
		var sum int
		for i < n && arr[i].color == arr[j].color {
			sum += arr[i].size
			i++
			for w := max(0, i-j-1); w <= min(i-j+1, n); w++ {
				if best < dp[w]+sum {
					best = dp[w] + sum
					pickColors[0] = fp[w]
					pickColors[1] = arr[i-1].color
				}
			}
		}

		sum = 0
		for i1 := j; i1 < i; i1++ {
			sum += arr[i1].size
			if sum > dp[i1-j+1] {
				dp[i1-j+1] = sum
				fp[i1-j+1] = arr[i1].color
			}
		}
	}

	get := func(color int) []cube {
		var res []cube
		for i := range n {
			if arr[i].color == color {
				res = append(res, arr[i])
			}
		}
		return res
	}

	a := get(pickColors[0])
	b := get(pickColors[1])
	if len(a) < len(b) {
		a, b = b, a
	}
	// len(a) >= len(b)

	for len(b) > 0 {
		res = append(res, a[0].id)
		res = append(res, b[0].id)
		a = a[1:]
		b = b[1:]
	}
	if len(a) > 0 {
		res = append(res, a[0].id)
	}

	return best, res
}
