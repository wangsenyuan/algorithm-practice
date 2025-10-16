package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)

	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	children := make([][]int, n)
	for i := range n {
		var v, d, p int
		fmt.Fscan(reader, &v, &d, &p)
		children[i] = []int{v, d, p}
	}
	return solve(children)
}

func solve(children [][]int) []int {
	n := len(children)

	var res []int

	next := make([]int, n)
	for i := range n {
		next[i] = i + 1
	}

	cry := make([]int, n+1)

	for i := 0; i < n; i = next[i] {
		v := children[i][0]
		res = append(res, i+1)
		prev := i
		var cur int
		for j := next[i]; j < n; j = next[j] {
			cry[j] += cur + max(v, 0)
			v--
			d, p := children[j][1], children[j][2]
			if p < cry[j] {
				next[prev] = next[j]
				cur += d
			} else {
				// change when j is still in line
				prev = j
			}
		}
	}

	return res
}
