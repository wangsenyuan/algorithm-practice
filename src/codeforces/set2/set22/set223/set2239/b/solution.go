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

func drive(reader *bufio.Reader) []int64 {
	var tc int
	fmt.Fscan(reader, &tc)
	res := make([]int64, tc)
	for i := range tc {
		var n, d int
		fmt.Fscan(reader, &n, &d)
		a := make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[j])
		}
		res[i] = solve(n, d, a)
	}
	return res
}

func solve(n, d int, a []int) int64 {
	// Total happiness expands to sum_i s_i * (2*d*a_i - A_i), where A_i is the
	// sum of a over the 2*d neighbors in i's field of view. Cross terms cancel,
	// so each s_i can be chosen independently: take i iff 2*d*a_i > A_i.
	pref := make([]int64, 2*n+1)
	for i := range 2 * n {
		pref[i+1] = pref[i] + int64(a[i%n])
	}
	rangeSum := func(l, r int) int64 {
		// sum a[l] + ... + a[r-1] on the doubled circle
		return pref[r] - pref[l]
	}

	var ans int64
	for i := range n {
		// neighbors [i-d, i) and (i, i+d] on the circle
		A := rangeSum(i+n-d, i+n) + rangeSum(i+1, i+1+d)
		contrib := int64(2*d)*int64(a[i]) - A
		if contrib > 0 {
			ans += contrib
		}
	}
	return ans
}
