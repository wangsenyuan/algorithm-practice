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
	ans := make([]int, t)
	for i := range t {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, 2*n)
		for j := range a {
			fmt.Fscan(reader, &a[j])
		}
		ans[i] = solve(n, a)
	}
	return ans
}

func solve(n int, a []int) int {
	freq := make([]int, n+1)

	play := func(l int, r int) int {
		clear(freq)
		var mex int
		for l >= 0 && r < len(a) && a[l] == a[r] {
			freq[a[l]]++
			for freq[mex] > 0 {
				mex++
			}
			l--
			r++
		}
		return mex
	}

	var pos []int

	for i, v := range a {
		if v == 0 {
			pos = append(pos, i)
		}
	}
	l, r := pos[0], pos[1]
	mid := (l + r) / 2
	var ans int
	if (r-l+1)%2 == 0 {
		ans = play(mid, mid+1)
	} else {
		ans = play(mid, mid)
	}

	ans = max(ans, play(l, l), play(r, r))

	return ans
}
