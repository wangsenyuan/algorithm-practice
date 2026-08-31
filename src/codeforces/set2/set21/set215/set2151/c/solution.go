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

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		ans := drive(reader)
		for i, v := range ans {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) []int64 {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int64, 2*n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int64) []int64 {
	n := len(a) / 2
	m := len(a)

	res := make([]int64, n)
	pref := make([]int64, m)
	for i := 1; i < m; i++ {
		gap := a[i] - a[i-1]
		pref[i] = pref[i-1] + gap
		if i%2 == 1 {
			res[0] += gap
		} else if n >= 2 {
			res[1] += 2 * gap
		}
	}
	if n >= 2 {
		res[1] += res[0]
	}

	for k := 3; k <= n; k++ {
		// Compared with capacity k-2, the occupancy increases by two on
		// every gap from k-1 through 2n-k+1 (one-based gap indices).
		res[k-1] = res[k-3] + 2*(pref[m-k+1]-pref[k-2])
	}

	return res
}
