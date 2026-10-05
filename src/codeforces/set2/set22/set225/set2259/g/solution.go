package main

import (
	"bufio"
	"fmt"
	"os"
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
	var k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) []int {
	n := len(a)
	s := make([]int, n+1)
	for i := range a {
		s[i+1] = s[i] + a[i]
		a[i] -= i * k
	}

	ans := make([]int, n)
	j := n - 1
	for i := n - 1; i > 0; i-- {
		b := a[i-1] - k
		for a[j] <= b {
			j--
		}
		ans[i] = s[j+1] - s[i+1] - (i+j+1)*(j-i)/2*k - (j-i)*b
	}
	ans[0] = 0
	return ans
}
