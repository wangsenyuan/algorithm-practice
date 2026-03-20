package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)

	var n int
	fmt.Fscan(reader, &n)
	res := solve(n)
	fmt.Fprintln(w, len(res))
	for _, cur := range res {
		s := fmt.Sprintf("%v", cur)
		fmt.Fprintln(w, len(cur), s[1:len(s)-1])
	}
}

func solve(n int) [][]int {
	var res [][]int

	// when n = 10
	// (1, 3, 5, 7, 9) vs (2, 4, 6, 8)
	// (1, 3, 5, 2, 4) vs (7, 9, 6, 8)
	for i := 0; 1<<i < n; i++ {
		var cur []int
		for j := range n {
			if (j+1)&(1<<i) != 0 {
				cur = append(cur, j+1)
			}
		}
		res = append(res, cur)
	}

	return res
}
