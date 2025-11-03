package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) []int {
	n := len(a)

	freq := make([]int, n+1)
	ans := make([]int, n+1)

	for l := 0; l < n; l++ {
		clear(freq)
		w := -1
		for r := l; r < n; r++ {
			v := a[r]
			freq[v]++
			if w == -1 || freq[v] > freq[w] || freq[v] == freq[w] && v < w {
				w = v
			}
			ans[w]++
		}
	}

	return ans[1:]
}
