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
	var n, x, y int
	fmt.Fscan(reader, &n, &x, &y)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(x, y, p)
}

func solve(x, y int, p []int) []int {
	// n := len(p)
	first := p[x:y]
	second := slices.Clone(p[:x])
	second = append(second, p[y:]...)

	w := slices.Min(first)
	for i, v := range first {
		if w == v {
			shift(first, i)
			break
		}
	}
	// 然后找到second里面, 第一个比w大的数, 在这个数插入

	for i, v := range second {
		if v > w {
			ans := slices.Clone(second[:i])
			ans = append(ans, first...)
			ans = append(ans, second[i:]...)
			return ans
		}
	}

	// second < first
	return append(second, first...)
}

func shift(arr []int, k int) {
	slices.Reverse(arr[:k])
	slices.Reverse(arr[k:])
	slices.Reverse(arr)
}
