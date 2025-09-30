package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res[0], res[1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) []int {
	sort.Ints(a)
	n := len(a)
	var sum int
	var res1 int
	var res2 int
	for i, v := range a {
		res1 += v*i - sum
		res2 += v
		sum += v
	}

	res1 *= 2
	num := res1 + res2
	div := n
	g := gcd(num, div)
	return []int{num / g, div / g}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
