package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for _, ok := range drive(reader) {
		if ok {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func drive(reader *bufio.Reader) []bool {
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]byte, q)
	for i := range q {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			var c string
			fmt.Fscan(reader, &c)
			queries[i] = []byte{c[0]}
		} else {
			queries[i] = []byte{'2'}
		}
	}
	return solve(queries)
}

type state struct {
	level int
	ok    bool
}

func solve(queries [][]byte) []bool {
	n := len(queries)
	ans := make([]bool, n)

	var stack []state
	stack = append(stack, state{0, true})

	for i, cur := range queries {
		if cur[0] == '2' {
			stack = stack[:len(stack)-1]
		} else {
			level := stack[len(stack)-1].level
			ok := stack[len(stack)-1].ok
			if cur[0] == '(' {
				level++
			} else {
				level--
			}
			if level < 0 {
				ok = false
			}
			stack = append(stack, state{level, ok})
		}
		last := stack[len(stack)-1]
		ans[i] = last.ok && last.level == 0
	}

	return ans
}
