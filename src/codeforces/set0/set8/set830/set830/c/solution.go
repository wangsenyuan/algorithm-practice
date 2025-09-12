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
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

const inf = 1e18

func solve(k int, a []int) int {
	// x := slices.Max(a)
	var arr []int
	arr = append(arr, 1)
	for _, v := range a {
		k += v
		for j := 1; j*j <= v; j++ {
			arr = append(arr, j, (v+j-1)/j)
		}
	}

	arr = append(arr, inf)

	sort.Ints(arr)
	arr = slices.Compact(arr)

	var ans int
	for i := 0; i+1 < len(arr); i++ {
		l := arr[i]
		// r := arr[i+1]
		// 在d = [l...r)中间的部分， a[?] / d 都是相同的 = a[?] / d
		var sum int
		for _, v := range a {
			sum += (v + l - 1) / l
		}
		d := k / sum
		if l <= d {
			ans = max(ans, d)
		}
	}

	return ans
}
