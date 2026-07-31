package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	slices.Sort(a)
	n := len(a)
	freq := make([]int, n+1)

	check := func(w int) bool {
		// 能否得到mex(a) = w
		var todo []int
		clear(freq)
		for _, v := range a {
			if v < w {
				if freq[v] == 0 {
					freq[v]++
				} else {
					todo = append(todo, v)
				}
			} else {
				// v >= w
				todo = append(todo, v)
			}
		}
		var mex int
		for mex < w {
			if freq[mex] == 0 {
				for len(todo) > 0 && mex*2 >= todo[0] {
					todo = todo[1:]
				}
				if len(todo) == 0 {
					return false
				}
				// todo[0] 可以得到 mex
				todo = todo[1:]
			}
			mex++
		}
		return mex == w
	}

	res := sort.Search(n+1, func(w int) bool {
		return !check(w)
	})
	return res - 1
}
