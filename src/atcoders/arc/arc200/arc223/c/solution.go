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

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	res := make([]int, t)
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

func solve(a []int) int {
	n := len(a)
	c := slices.Clone(a)
	slices.Sort(c)

	b := make([]int, n)
	vis := make([]bool, n)
	for i, v := range c {
		b[i] = v % n
		if vis[b[i]] {
			return 0
		}
		vis[b[i]] = true
	}

	res, fact := 1, 1
	for i := 1; i < n; i++ {
		fact = fact * i % n
		res = res * fact % n
	}

	clear(vis)
	cycles := 0
	for i := range n {
		if !vis[i] {
			cycles++
			for j := i; !vis[j]; j = b[j] {
				vis[j] = true
			}
		}
	}

	if (n-cycles)&1 == 1 {
		res = (n - res) % n
	}

	return res
}
