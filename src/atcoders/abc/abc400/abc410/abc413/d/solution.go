package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		if res {
			buf.WriteString("Yes\n")
		} else {
			buf.WriteString("No\n")
		}
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(arr []int) bool {
	n := len(arr)
	if n <= 2 {
		return true
	}
	mn := slices.Min(arr)
	mx := slices.Max(arr)
	if mx == mn {
		// 1
		return true
	}
	if mn > 0 || mx < 0 {
		slices.Sort(arr)
		// factor > 0
		// arr[1] / arr[0] = arr[2] / arr[1] = ... = arr[n-1] / arr[n-2]
		for i := 1; i+1 < n; i++ {
			if arr[i]*arr[i] != arr[i-1]*arr[i+1] {
				return false
			}
		}
		return true
	}
	// factor < 0
	// -2, 4, -8, 16 所以可以处理绝对值
	// -4, 4, 4, 4, 这个不行

	slices.SortFunc(arr, func(x int, y int) int {
		return abs(x) - abs(y)
	})

	if abs(arr[0]) == abs(arr[n-1]) {
		// factor -1
		cnt := make([]int, 2)
		for _, x := range arr {
			if x < 0 {
				cnt[0]++
			} else {
				cnt[1]++
			}
		}
		return abs(cnt[0]-cnt[1]) <= 1
	}

	for i := 1; i+1 < n; i++ {
		if arr[i]*arr[i] != arr[i-1]*arr[i+1] {
			return false
		}
	}
	return true
}

func abs(x int) int {
	return max(x, -x)
}
