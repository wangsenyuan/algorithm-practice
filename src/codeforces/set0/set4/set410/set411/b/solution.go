package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, v := range res {
		buf.WriteString(fmt.Sprintf("%d\n", v))
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) []int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(k, a)
}

func solve(k int, a [][]int) []int {
	n := len(a)
	m := len(a[0])
	s1 := make([]int, n)
	s2 := make([]int, k)
	for j := range m {
		target := make([][]int, k+1)
		for i := range n {
			if s1[i] > 0 || a[i][j] == 0 {
				continue
			}
			x := a[i][j]
			if s2[x-1] > 0 {
				// 这个马上被lock住了
				s1[i] = j + 1
				continue
			}
			target[x] = append(target[x], i)
		}
		for x := 1; x <= k; x++ {
			if len(target[x]) > 1 {
				for _, i := range target[x] {
					s1[i] = j + 1
				}
				s2[x-1] = j + 1
			}
		}
	}

	return s1
}
