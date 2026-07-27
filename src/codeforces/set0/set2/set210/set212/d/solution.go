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

	res := drive(reader)
	for i, v := range res {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprintf(writer, "%.15f", v)
	}
	fmt.Fprintln(writer)
}

func drive(reader *bufio.Reader) []float64 {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	ks := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &ks[i])
	}
	return solve(a, ks)
}

func solve(a []int, ks []int) []float64 {
	n := len(a)
	L := make([]int, n)
	R := make([]int, n)
	stack := make([]int, n)
	var top int
	for i, v := range a {
		L[i] = -1
		R[i] = n
		for top > 0 && a[stack[top-1]] >= v {
			R[stack[top-1]] = i
			top--
		}
		if top > 0 {
			// a[stack[top-1]] < a[i]
			L[i] = stack[top-1]
		}
		stack[top] = i
		top++
	}

	// change[k] records how the slope sum[k] - sum[k-1] changes at k.
	change := make([]int64, n+3)

	for i := range n {
		l, r := L[i], R[i]
		w := min(i-l, r-i)
		v := max(i-l, r-i)
		d := r - l - 1
		// 1...w, 贡献是 a[i], 2 * a[i], ... w * a[i]
		// w+1...v, 贡献 = w * a[i]
		// v+1...d, 贡献 = (w - 1) * a[i], ... 1 * a[i]
		value := int64(a[i])
		change[1] += value
		change[w+1] -= value
		change[v+1] -= value
		change[d+2] += value
	}

	sum := make([]int64, n+1)
	var slope int64
	for k := 1; k <= n; k++ {
		slope += change[k]
		// sum[k] is the sum of the minima of all subarrays of length k.
		sum[k] = sum[k-1] + slope
	}

	res := make([]float64, len(ks))
	for i, k := range ks {
		res[i] = float64(sum[k]) / float64(n-k+1)
	}
	return res
}
