package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	var buf bytes.Buffer
	for range tc {
		res := drive(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}
	return solve(nums)
}

func solve(nums []int) int {
	sort.Ints(nums)
	slices.Reverse(nums)

	x := nums[0]

	n := len(nums)

	var res int
	for c := 0; c < n; c++ {
		a := n
		y := max(nums[c], x-nums[c])
		for b := c + 1; b+1 < n; b++ {
			for a-1 > b && nums[a-1]+nums[b] <= y {
				a--
			}
			if a <= b+1 {
				break
			}
			res += max(0, a-b-1)
		}
	}
	return res
}
