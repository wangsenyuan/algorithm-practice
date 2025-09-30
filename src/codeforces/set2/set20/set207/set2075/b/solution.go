package main

import (
	"bufio"
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
		var n, k int
		fmt.Fscan(reader, &n, &k)
		a := readNNums(reader, n)
		fmt.Fprintln(writer, solve(a, k))
	}
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func solve(a []int, k int) int {
	// 1 2 4 3, k = 2
	// 考虑最后一个数，如果它作为最后一个数，那么它两边都必须有被选中的
	// 否则它就没法是最后一个数
	n := len(a)
	if k == 1 {
		if n == 2 {
			return a[0] + a[1]
		}
		x := slices.Max(a[1 : n-1])
		return max(x+max(a[0], a[n-1]), a[0]+a[n-1])
	}
	// k > 1
	// 只要有两个数可以选，那么就可以先选最大的k个数，然后去选剩下的那个best的数
	// 且总是有办法的
	slices.Sort(a)
	k++
	var res int
	for i := n - k; i < n; i++ {
		res += a[i]
	}
	return res
}
