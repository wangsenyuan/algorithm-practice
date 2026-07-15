package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	res := make([]int, t)
	for i := range t {
		var n int
		var W int
		fmt.Fscan(reader, &n, &W)
		w := make([]int, n)
		v := make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &w[j], &v[j])
		}
		res[i] = solve(W, w, v)
	}
	return res
}

func solve(W int, w, v []int) int {

	n := len(w)
	pref1 := make([]int, n+1)
	pref2 := make([]int, n+1)
	for i := range n {
		pref1[i+1] = min(W+1, pref1[i]+w[i])
		pref2[i+1] = pref2[i] + v[i]
	}

	if pref1[n] <= W {
		return pref2[n]
	}

	var f func(i int, rest int) int
	f = func(i int, rest int) (res int) {
		if i < 0 {
			return 0
		}

		if w[i] <= rest {
			res = max(res, pref2[i], v[i]+f(i-1, rest-w[i]))
		} else {
			res = f(i-1, rest)
		}
		return res
	}

	return f(n-1, W)
}
