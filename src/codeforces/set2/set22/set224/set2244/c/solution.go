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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var n, x, y int
	fmt.Fscan(reader, &n, &x, &y)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(x, y, p)
}

func convert(b bool) string {
	if b {
		return "YES"
	}
	return "NO"
}

func solve(x, y int, p []int) string {
	n := len(p)
	w := gcd(x, y)
	if w >= n {
		return convert(sort.IntsAreSorted(p))
	}

	type pair struct {
		first  int
		second int
	}

	arr := make([][]pair, w)

	for i, v := range p {
		arr[i%w] = append(arr[i%w], pair{v, i})
	}

	for i := range w {
		slices.SortFunc(arr[i], func(x pair, y pair) int {
			return x.first - y.first
		})
	}

	buf := make([]int, n)
	for i := range w {
		var ids []int
		for _, cur := range arr[i] {
			ids = append(ids, cur.second)
		}
		slices.Sort(ids)
		for i, cur := range arr[i] {
			buf[ids[i]] = cur.first
		}
	}

	return convert(sort.IntsAreSorted(buf))
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
