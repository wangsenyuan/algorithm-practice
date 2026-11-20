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

	for _, pair := range drive(reader) {
		fmt.Fprintln(writer, pair[0], pair[1])
	}
}

func drive(reader *bufio.Reader) [][]int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, a)
}

func solve(n int, a []int) [][]int {
	pos := make([]int, n+1)

	hi := make([]int, n)
	lo := make([]int, n)
	arr := make([]int, n)
	for i := range n {
		pos[i] = i
		hi[i] = i
		lo[i] = i
		arr[i] = i
	}

	like := func(v int) {
		v--
		if pos[v] == 0 {
			return
		}
		w := arr[pos[v]-1]
		pos[w], pos[v] = pos[v], pos[w]
		arr[pos[v]] = v
		arr[pos[w]] = w
		if hi[v] > pos[v] {
			hi[v] = pos[v]
		}
		if lo[w] < pos[w] {
			lo[w] = pos[w]
		}
	}

	for _, v := range a {
		like(v)
	}

	ans := make([][]int, n)

	for i := range n {
		lo[i]++
		hi[i]++
		ans[i] = []int{hi[i], lo[i]}
	}

	return ans
}
