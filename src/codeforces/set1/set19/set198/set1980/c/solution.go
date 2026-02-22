package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		if res {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	var m int
	fmt.Fscan(reader, &m)
	d := readNNums(reader, m)
	return solve(a, b, d)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, b []int, d []int) bool {
	n := len(a)
	// m := len(d)
	vis := make(map[int]bool)

	diff := make(map[int]int)
	for i := range n {
		if a[i] != b[i] {
			// a[i] 不重要，它需要被替换成d中的某个数
			diff[b[i]]++
		}
		vis[b[i]] = true
	}

	var stack []int

	for _, v := range d {
		if !vis[v] {
			stack = append(stack, v)
		} else {
			stack = stack[:0]
			if diff[v] > 0 {
				diff[v]--
				if diff[v] == 0 {
					delete(diff, v)
				}
			}
		}
	}

	return len(diff) == 0 && len(stack) == 0
}
