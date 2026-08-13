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

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		for i, x := range res {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, x)
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	blogs := make([][]int, n)
	for i := range n {
		var l int
		fmt.Fscan(reader, &l)
		blogs[i] = make([]int, l)
		for j := range l {
			fmt.Fscan(reader, &blogs[i][j])
		}
	}
	return solve(blogs)
}

func solve(blogs [][]int) []int {
	for i, cur := range blogs {
		slices.Reverse(cur)
		vis := make(map[int]bool)
		var j int
		for _, v := range cur {
			if !vis[v] {
				cur[j] = v
				j++
			}
			vis[v] = true
		}
		blogs[i] = cur[:j]
	}
	var res []int

	marked := make(map[int]bool)

	for len(blogs) > 0 {
		slices.SortFunc(blogs, func(a []int, b []int) int {
			// compare a + b vs b + a
			for i := 0; i < len(a)+len(b); i++ {
				var x, y int
				if i < len(a) {
					x = a[i]
				} else {
					x = b[i-len(a)]
				}
				if i < len(b) {
					y = b[i]
				} else {
					y = a[i-len(b)]
				}
				if x == y {
					continue
				}
				return x - y
			}
			return 0
		})

		first := blogs[0]
		for _, v := range first {
			res = append(res, v)
			marked[v] = true
		}

		var next [][]int

		for i := 1; i < len(blogs); i++ {
			cur := blogs[i]
			var j int
			for _, v := range cur {
				if !marked[v] {
					cur[j] = v
					j++
				}
			}
			if j > 0 {
				next = append(next, cur[:j])
			}
		}

		blogs = next
	}

	return res
}
