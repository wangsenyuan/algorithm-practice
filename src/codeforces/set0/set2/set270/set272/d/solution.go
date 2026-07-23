package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	return solve(a, b, m)
}

func solve(a, b []int, m int) int {
	freq := make(map[int]int)
	for _, v := range a {
		freq[v]++
	}
	for _, v := range b {
		freq[v]++
	}

	d := 0
	for i := range a {
		if a[i] == b[i] {
			d++
		}
	}

	res := 1
	for _, c := range freq {
		for i := 2; i <= c; i++ {
			x := i
			for d > 0 && x%2 == 0 {
				x /= 2
				d--
			}
			res = res * x % m
		}
	}
	return res
}
