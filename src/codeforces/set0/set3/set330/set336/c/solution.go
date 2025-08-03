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

	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	res := solve(a)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

const H = 30

func solve(a []int) []int {
	sort.Ints(a)
	// n := len(a)
	a = slices.Compact(a)
	for v := H - 1; v >= 0; v-- {
		// 希望找一些数字 v是set的，比v小的位数要and = 0
		var arr []int
		and := 1<<v - 1
		for _, x := range a {
			if (x>>v)&1 == 1 {
				arr = append(arr, x)
				and &= x
			}
		}
		if len(arr) > 0 && and&(1<<v-1) == 0 {
			return arr
		}
	}
	return nil
}
