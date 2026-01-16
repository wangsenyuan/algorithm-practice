package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a []int, b []int) int {

	n := len(a)
	for i := 1; i < n; i++ {
		a[i] += a[i-1]
	}

	m := len(b)

	// a[?] - b[i] = score

	freq := make(map[int]int)

	mem := map[int]bool{}
	check := func(score int) bool {
		if _, ok := mem[score]; ok {
			return false
		}
		mem[score] = true

		clear(freq)

		for _, v := range b {
			freq[v]++
		}

		var cnt int

		for _, v := range a {
			v += score
			cnt += freq[v]
			if freq[v] == 1 {
				freq[v]--
			}
		}

		return cnt == m
	}

	var res int

	for i := range n {
		score := b[0] - a[i]
		if check(score) {
			res++
		}
	}

	return res
}
