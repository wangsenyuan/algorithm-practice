package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader io.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, a)
}

func solve(n int, a []int) int {
	var res int
	var pos int

	for i := 0; i < n; {
		if a[i] == 0 {
			i++
			continue
		}
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		res += min(j-pos, i-j)
		pos += i - j
	}
	return res
}
