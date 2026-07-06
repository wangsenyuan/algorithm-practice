package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	var s string
	fmt.Fscan(reader, &s)
	w := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &w[i])
	}
	return solve(s, w)
}

func solve(s string, w []int) int {
	// TODO: solve by hand first.
	n := len(s)

	type data struct {
		weight int
		adult  int
	}

	arr := make([]data, n)
	for i := range n {
		arr[i] = data{weight: w[i], adult: int(s[i] - '0')}
	}

	slices.SortFunc(arr, func(a, b data) int {
		return a.weight - b.weight
	})

	suf := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		suf[i] = suf[i+1] + arr[i].adult
	}

	// 如果X = arr[0].weight
	ans := suf[0]
	var pref int
	for i := 0; i < n; {
		j := i
		for i < n && arr[i].weight == arr[j].weight {
			pref += 1 - arr[i].adult
			i++
		}
		ans = max(ans, pref+suf[i])
	}

	return ans
}
