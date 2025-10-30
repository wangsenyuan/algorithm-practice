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

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	return solve(m, a, b)
}

func solve(m int, a []int, b []int) []int {
	getFreq := func(arr []int) []int {
		res := make([]int, m)
		for _, v := range arr {
			res[v]++
		}
		return res
	}

	fa := getFreq(a)
	fb := getFreq(b)

	ans := make([]int, m)

	var f []int

	for i := range m {
		for range fa[i] {
			f = append(f, i+1)
		}
		for range fb[m-1-i] {
			if len(f) > 0 && f[len(f)-1] > 0 {
				x := f[len(f)-1] - 1
				f = f[:len(f)-1]
				y := m - 1 - i
				ans[(x+y)%m]++
			} else {
				f = append(f, -(m - 1 - i + 1))
			}
		}
	}

	tot := len(f)

	for i := 0; i < tot/2; i++ {
		x := -f[i] - 1
		y := f[tot-1-i] - 1
		ans[(x+y)%m]++
	}

	var res []int
	for i := m - 1; i >= 0; i-- {
		for range ans[i] {
			res = append(res, i)
		}
	}

	return res
}
