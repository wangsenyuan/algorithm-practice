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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		for _, ans := range drive(reader) {
			fmt.Fprintln(writer, ans)
		}
	}
}

func drive(reader *bufio.Reader) []int {
	var s int
	var q int
	fmt.Fscan(reader, &s, &q)
	queries := make([][2]int, q)
	for i := range queries {
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(s, queries)
}

type rect struct {
	h int
	w int
}

func solve(s int, queries [][2]int) []int {
	var arr []rect
	arr = append(arr, rect{s + 1, 0})

	for w := 1; w <= s/w; w++ {
		if s%w == 0 {
			arr = append(arr, rect{s / w, w})
			if w != s/w {
				arr = append(arr, rect{w, s / w})
			}
		}
	}
	arr = append(arr, rect{0, s + 1})

	slices.SortFunc(arr, func(a rect, b rect) int {
		return a.w - b.w
	})

	n := len(arr)
	pref := make([]int, n)
	for i := 1; i < n; i++ {
		pref[i] = pref[i-1] + (arr[i].w-arr[i-1].w)*arr[i].h
	}

	ans := make([]int, len(queries))

	for pos, cur := range queries {
		x, y := cur[0], cur[1]
		x = min(x, s)
		y = min(y, s)
		i := sort.Search(n, func(i int) bool {
			return arr[i].w > x
		})
		j := sort.Search(n, func(j int) bool {
			return arr[j].h <= y
		})
		j--
		// arr[j].h > y
		if j < i {
			ans[pos] = pref[i] - pref[j] - (arr[i].w-x)*arr[i].h
			ans[pos] += arr[j].w * y
		} else {
			// j > i
			ans[pos] = x * y
		}
	}

	return ans
}
