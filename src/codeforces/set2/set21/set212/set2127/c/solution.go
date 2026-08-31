package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
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

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(k, a, b)
}

type pair struct {
	first  int
	second int
}

func solve(k int, a, b []int) int {
	n := len(a)
	// 这个结果不会影响结果
	var sum int
	arr := make([]pair, n)
	for i := range n {
		l, r := min(a[i], b[i]), max(a[i], b[i])
		sum += r - l
		arr[i] = pair{l, r}
	}
	// 如果 b[i] < a[j], dv = (a[j] - a[i]) + (b[j] - b[i]) - (b[j] - a[j]) - (b[i] - a[i])
	// = a[j] - a[i] + b[j] - b[i] - b[j] + a[j] - b[i] + a[i]
	// = 2 * a[j] - 2 * b[i]
	// 如果 a[i] > b[j], 那么合理猜 dv = 2 * a[i] - 2 * b[j]
	//  如果 b[i] >= a[j], 那么就不会变了, 比如 (1, 3), (2, 4)

	slices.SortFunc(arr, func(a pair, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})
	// 如果存在重叠的, 就直接返回sum
	mx := -1
	var best = 1 << 60
	for i, cur := range arr {
		if mx >= cur.first {
			return sum
		}
		mx = cur.second
		if i > 0 {
			best = min(best, 2*(cur.first-arr[i-1].second))
		}
	}
	// 都不重叠
	return sum + best
}
