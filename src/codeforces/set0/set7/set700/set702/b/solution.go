package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	res := solve(a)
	fmt.Println(res)
}

const H = 31

func solve(a []int) int {
	freq := make(map[int]int)
	var res int
	for _, v := range a {
		for i := range H {
			if v < 1<<i {
				w := 1<<i - v
				res += freq[w]
			}
		}
		freq[v]++
	}

	return res
}
