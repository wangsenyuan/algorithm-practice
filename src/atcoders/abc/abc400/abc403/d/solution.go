package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, d int
	fmt.Fscan(reader, &n, &d)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, d)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, d int) int {
	slices.Sort(a)
	n := len(a)
	if d == 0 {
		a = slices.Compact(a)
		return n - len(a)
	}
	// d > 0
	// 相差d的那些先找出来
	play := func(arr []pair) int {
		if len(arr) == 1 {
			return arr[0].second
		}
		// 到i为止，且没有两个位置是相邻的最大和
		// x,y 表示保留最多的部分
		x, y := arr[0].second, arr[1].second
		y = max(x, y)
		for i := 2; i < len(arr); i++ {
			z := max(y, arr[i].second+x)
			x, y = y, z
		}
		return max(x, y)
	}

	dp := make(map[int][]pair)

	var res int

	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		if arr, ok := dp[a[j]%d]; ok {
			last := arr[len(arr)-1]
			if last.first+d == a[j] {
				// 它们相差为d
				dp[a[j]%d] = append(arr, pair{a[j], i - j})
			} else {
				// 已经超过d了
				res += play(arr)
				dp[a[j]%d] = []pair{{a[j], i - j}}
			}
		} else {
			dp[a[j]%d] = []pair{{a[j], i - j}}
		}
	}

	for _, arr := range dp {
		res += play(arr)
	}

	return n - res
}
