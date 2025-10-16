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

	_, ok, res := drive(reader)
	if !ok {
		fmt.Fprintln(writer, "NO")
		return
	}
	fmt.Fprintln(writer, "YES")
	for _, cur := range res {
		fmt.Fprintln(writer, cur[0], cur[1])
	}
}

func drive(reader *bufio.Reader) (p []int, ok bool, res [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	p = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	ok, res = solve(n, slices.Clone(p))
	return
}

func solve(n int, p []int) (bool, [][]int) {
	if n == 1 {
		return true, nil
	}
	for i := range n {
		p[i]--
	}

	for i := range n {
		if p[i] == i {
			return solve1(n, i)
		}
	}
	for i := range n {
		if p[p[i]] == i {
			return solve2(n, p, i, p[i])
		}
	}

	return false, nil
}

func solve1(n int, i int) (bool, [][]int) {
	var res [][]int
	for j := range n {
		if i != j {
			res = append(res, []int{i + 1, j + 1})
		}
	}
	return true, res
}

func solve2(n int, p []int, l int, r int) (bool, [][]int) {
	marked := make([]bool, n)
	var res [][]int
	for i := range n {
		if i == l || i == r {
			continue
		}
		if !marked[i] {
			j := i
			u := l
			var cnt int
			for !marked[j] {
				res = append(res, []int{u + 1, j + 1})
				u ^= l ^ r
				marked[j] = true
				cnt++
				j = p[j]
			}
			if cnt&1 == 1 {
				return false, nil
			}
		}
	}
	res = append(res, []int{l + 1, r + 1})

	return true, res
}
