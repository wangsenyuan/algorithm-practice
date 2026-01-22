package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	x := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &x[i])
	}
	return solve(x)
}

func solve(x []int) int {
	slices.Sort(x)
	n := len(x)
	if n <= 2 {
		return 0
	}
	var d int
	for i := 0; i+1 < n; i++ {
		d = gcd(d, x[i+1]-x[i])
	}
	var res int
	for i := 0; i+1 < n; i++ {
		res += (x[i+1]-x[i])/d - 1
	}

	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
