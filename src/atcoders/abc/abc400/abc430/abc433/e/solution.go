package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		_, _, a := drive(reader)
		if len(a) == 0 {
			buf.WriteString("No\n")
		} else {
			buf.WriteString("Yes\n")
			for _, row := range a {
				for _, v := range row {
					buf.WriteString(fmt.Sprintf("%d ", v))
				}
				buf.WriteByte('\n')
			}
		}
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) (X []int, Y []int, a [][]int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	X = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &X[i])
	}
	Y = make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &Y[i])
	}
	a = solve(X, Y)
	return
}

type pair struct {
	first  int
	second int
}

func solve(X []int, Y []int) [][]int {
	n := len(X)
	m := len(Y)
	p1 := make([]int, n*m)
	p2 := make([]int, n*m)
	for i := range n * m {
		p1[i] = -1
		p2[i] = -1
	}
	for i, v := range X {
		v--
		if p1[v] != -1 {
			return nil
		}
		p1[v] = i
	}
	for i, v := range Y {
		v--
		if p2[v] != -1 {
			return nil
		}
		p2[v] = i
	}

	ok := make([][][]int, n*m)
	ans := make([][]int, n)
	for i, v := range X {
		ans[i] = make([]int, m)
		for j, w := range Y {
			ans[i][j] = -1
			ok[min(v, w)-1] = append(ok[min(v, w)-1], []int{i, j})
		}
	}

	var q [][]int

	for v := n*m - 1; v >= 0; v-- {
		if p1[v] < 0 && p2[v] < 0 {
			if len(q) == 0 {
				return nil
			}
			cur := q[0]
			q = q[1:]
			i, j := cur[0], cur[1]
			ans[i][j] = v + 1
			q = append(q, ok[v]...)
		} else if p1[v] < 0 || p2[v] < 0 {
			if len(ok[v]) == 0 {
				return nil
			}
			cur := ok[v][0]
			i, j := cur[0], cur[1]
			ans[i][j] = v + 1
			q = append(q, ok[v][1:]...)
		} else {
			i, j := p1[v], p2[v]
			ans[i][j] = v + 1
			for _, cur := range ok[v] {
				if cur[0] != i || cur[1] != j {
					q = append(q, cur)
				}
			}
		}
	}
	return ans
}
