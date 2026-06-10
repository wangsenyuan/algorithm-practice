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
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, a)
}

func solve(n int, paints []int) int {
	slices.Sort(paints)
	m := len(paints)

	var res int

	for i := 1; i < n; i++ {
		j := sort.SearchInts(paints, i)
		// paints[j] >= i
		k := sort.SearchInts(paints, n-i)
		// paints[k] >= m - i
		cur := (m-j)*(m-k) - min(m-j, m-k)

		res += cur
	}

	return res
}
