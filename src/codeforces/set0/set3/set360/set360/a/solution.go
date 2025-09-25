package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	} else {
		fmt.Println("YES")
		s := fmt.Sprintf("%v", res)
		fmt.Println(s[1 : len(s)-1])
	}
}

func drive(reader *bufio.Reader) (ops [][]int, res []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	ops = make([][]int, m)
	for i := 0; i < m; i++ {
		ops[i] = make([]int, 4)
		fmt.Fscan(reader, &ops[i][0], &ops[i][1], &ops[i][2], &ops[i][3])
	}
	res = solve(n, ops)
	return
}

const inf = 1e9

func solve(n int, ops [][]int) []int {
	offset := make([]int, n)

	for _, op := range ops {
		if op[0] == 1 {
			l, r, d := op[1], op[2], op[3]
			l--
			r--
			for j := l; j <= r; j++ {
				offset[j] += d
			}
		}
	}

	a := make([]int, n)

	for i := range n {
		a[i] = inf
	}
	m := len(ops)
	for i := m - 1; i >= 0; i-- {
		t, l, r, d := ops[i][0], ops[i][1], ops[i][2], ops[i][3]
		l--
		r--
		if t == 1 {
			for j := l; j <= r; j++ {
				offset[j] -= d
			}
		} else {
			for j := l; j <= r; j++ {
				v := offset[j]
				// d-v是a[j]的上限
				a[j] = min(a[j], d-v)
			}
		}
	}

	// offset must be all 0

	for _, op := range ops {
		t, l, r, d := op[0], op[1], op[2], op[3]
		l--
		r--
		if t == 1 {
			for j := l; j <= r; j++ {
				offset[j] += d
			}
		} else {
			var x int = -inf
			for j := l; j <= r; j++ {
				x = max(x, a[j]+offset[j])
			}
			if x != d {
				return nil
			}
		}
	}

	return a
}
