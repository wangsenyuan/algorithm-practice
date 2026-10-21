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
		ok, res := drive(reader)
		if !ok {
			fmt.Fprintln(writer, -1)
			continue
		}
		fmt.Fprintln(writer, len(res))
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) (bool, []int) {
	var n int
	fmt.Fscan(reader, &n)
	var s string
	fmt.Fscan(reader, &s)
	return solve(s)
}

func solve(s string) (ok bool, res []int) {
	n := len(s)
	pos := make([][]int, 2)

	for i := range n {
		x := int(s[i] - '0')
		pos[x] = append(pos[x], i+1)
	}

	if len(pos[1]) == 0 {
		return true, nil
	}
	best := n + 1

	if len(pos[1])&1 == 0 {
		// 偶数个1, 除了i, 其他的1都操作的时候, 正好反转i
		best = len(pos[1])
	}

	if len(pos[0])&1 == 1 {
		best = min(best, len(pos[0]))
	}

	if best > n {
		return false, nil
	}

	if best == len(pos[1]) && len(pos[1])&1 == 0 {
		return true, pos[1]
	}
	return true, pos[0]
}
