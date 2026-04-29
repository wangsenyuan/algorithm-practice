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
	c := make([]int, n)
	for _, v := range a {
		c[v-1]++
	}
	slices.Sort(c)
	res := make([]int, n)
	for i := range n {
		res[i] = n
	}

	for j, k := 0, 0; k <= n; k++ {
		for j < n && c[j] <= k {
			j++
		}
		// 这个s是剩余的元素个数
		var s int
		for i := j; i <= n && s < n; i++ {
			d := k + n - i
			res[s] = min(res[s], d)
			if i < n {
				s += c[i] - k
			}
		}
	}

	for i := 1; i < n; i++ {
		res[i] = min(res[i], res[i-1])
	}

	return res
}
