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
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) int {
	x := play(a)
	// a被处理一次后，就是 0,1,2 mex, 4, 5, mex .... 类似这样子
	if k == 1 {
		return sum(x)
	}
	y := play(x)
	if k == 2 {
		return sum(y)
	}
	// y的模式肯定是 0, 1, 2, 3 3 3 3 3
	// 现在需要知道y的模式
	z := play(y)
	if k&1 == 1 {
		return sum(z)
	}
	w := play(z)
	return sum(w)
}

func sum(a []int) int {
	var res int
	for _, v := range a {
		res += v
	}
	return res
}

func play(a []int) []int {
	n := len(a)
	freq := make([]int, n+2)
	for _, v := range a {
		freq[v]++
	}
	var mex int
	for freq[mex] > 0 {
		mex++
	}

	b := make([]int, n)

	for i := range n {
		b[i] = mex
		if a[i] < mex && freq[a[i]] == 1 {
			b[i] = a[i]
		}
	}

	return b
}
