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
	var n, T int
	fmt.Fscan(reader, &n, &T)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i])
	}
	return solve(T, t)
}

type pair struct {
	first  int
	second int
}

func solve(T int, t []int) int {
	n := len(t)

	// 假设dog在开始位置等待了s秒，开始work，那么 t[i] <= s + i 的部分，都可以被吃到
	// t[i] - i <= s, 同时 s + i < T
	// 那么这里就有n个时间点, 当 t[i] - i == s的时候，
	// 所有 t[j] - j < s 的部分，就可以被吃到，但是其中要排除掉 i > T - s 的部分
	// 如果迭代arr, 当等待的时间越长，那么能够进入冷却状态的碗就更多
	// 但是，越会来不及吃完

	s1 := make([]pair, n)
	s2 := make([]pair, n)

	var arr []int

	for i := range n {
		s1[i] = pair{t[i] - (i + 1), i}
		s2[i] = pair{T - i - 2, i}
		arr = append(arr, t[i]-(i+1), T-i-2)
	}

	slices.SortFunc(s1, func(a, b pair) int {
		return a.first - b.first
	})

	slices.SortFunc(s2, func(a, b pair) int {
		return a.first - b.first
	})

	arr = append(arr, 0)
	sort.Ints(arr)
	arr = slices.Compact(arr)
	var best int

	var sum int

	var i, j int
	for _, s := range arr {
		for i < n && s1[i].first == s {
			sum++
			i++
		}

		if s >= 0 {
			best = max(best, sum)
		}

		for j < n && s2[j].first == s {
			sum--
			j++
		}
	}

	return best
}
