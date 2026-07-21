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

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int64 {
	var t int
	fmt.Fscan(reader, &t)
	res := make([]int64, t)
	for i := range t {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[j])
		}
		res[i] = solve(a)
	}
	return res
}

func solve(a []int) int64 {
	var res int64
	n := len(a)
	for i := 0; i < n; {
		if a[i] == 0 {
			i++
			continue
		}
		j := i
		var sum int
		var mx int
		for i < n && a[i] > 0 {
			sum += a[i]
			if i&1 == j&1 {
				mx = max(mx, a[i])
			}
			i++
		}
		res += 2 * int64(sum)
		if (i-j)&1 == 1 {
			// 奇数长度
			res -= int64(mx)
		}
	}

	return res
}
