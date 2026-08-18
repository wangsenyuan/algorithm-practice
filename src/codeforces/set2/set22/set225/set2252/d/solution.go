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
		res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) []int {
	n := len(a)

	d := make([]int, n-1)
	for i := range n - 1 {
		d[i] = a[i+1] - a[i]
	}
	// 连续相同parity的可以swap
	ans := make([]int, n)
	ans[0] = a[0]

	for i := 0; i < n-1; {
		j := i
		for i < n-1 && d[i]&1 == d[j]&1 {
			i++
		}
		slices.Sort(d[j:i])
		for j < i {
			ans[j+1] = ans[j] + d[j]
			j++
		}
	}

	return ans
}
