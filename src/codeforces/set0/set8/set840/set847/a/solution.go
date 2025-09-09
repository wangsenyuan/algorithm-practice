package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) (n int, l []int, r []int, res [][]int) {
	fmt.Fscan(reader, &n)
	l = make([]int, n)
	r = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &l[i], &r[i])
	}
	return n, l, r, solve(n, l, r)
}

func solve(n int, l []int, r []int) [][]int {
	var head []int
	next := make([]int, n+1)

	for i := range n {
		if l[i] == 0 {
			head = append(head, i+1)
		} else {
			next[l[i]] = i + 1
		}

		// 这个貌似会有点多余
		if r[i] != 0 {
			next[i+1] = r[i]
		}
	}

	cur := head[0]
	head = head[1:]

	for cur != 0 {
		tmp := next[cur]
		if tmp == 0 {
			if len(head) == 0 {
				break
			}
			tmp = head[0]
			head = head[1:]
		}
		r[cur-1] = tmp
		l[tmp-1] = cur
		cur = tmp
	}

	res := make([][]int, n)
	for i := range n {
		res[i] = []int{l[i], r[i]}
	}
	return res
}
