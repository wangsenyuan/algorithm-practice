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
		_, res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) (a []int, res []int) {
	var n, k int

	fmt.Fscan(reader, &n, &k)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	res = solve(a, k)
	return
}

func solve(a []int, k int) []int {
	n := len(a)

	// 如果 suf[i:]是个回文，那么当前选择应该避免 = a[i-1]
	// 但是这样子似乎就变成了 O(n^2)，是不是选择没有出现的，或者是最远出现的那个数？

	marked := make([]bool, n+1)
	for _, v := range a {
		marked[v] = true
	}

	var missing []int
	for i := 1; i <= n; i++ {
		if !marked[i] {
			missing = append(missing, i)
		}
	}

	if len(missing) > 0 {
		a = append(a, missing[0])
	} else {
		a = append(a, a[0])
	}

	k--

	for k > 0 {
		x := a[len(a)-1]
		y := a[len(a)-2]
		z := 1
		for x == z || y == z {
			z++
		}
		a = append(a, z)
		k--
	}
	return a[n:]
}
