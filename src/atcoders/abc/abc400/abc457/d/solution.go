package main

import (
	"bufio"
	"fmt"
	"os"
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
	return solve(a, k)
}

func solve(a []int, k int) int {

	n := len(a)

	check := func(w int) bool {
		var tot int
		for i := 1; i <= n; i++ {
			// x = m * i + a[i] >= w
			if a[i-1] >= w {
				continue
			}
			cnt := (w - a[i-1] + i - 1) / i
			if tot > k-cnt {
				return true
			}
			tot += cnt
		}

		return false
	}

	best := a[0] + k + 1

	res := sort.Search(best, check)
	return res - 1
}
