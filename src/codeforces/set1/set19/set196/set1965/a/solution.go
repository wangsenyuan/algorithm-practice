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
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const Alice = "Alice"
const Bob = "Bob"

func solve(a []int) string {
	slices.Sort(a)
	a = slices.Compact(a)

	if len(a) == 1 {
		return Alice
	}

	// 第一个遇到差值超过1的
	players := []string{Alice, Bob}
	var w int
	for i := 0; i < len(a); i++ {
		if i == len(a)-1 {
			// 当前用户可以全部取走
			return players[w]
		}
		// 如果diff = 1， 那么当前player没有选择，只能把主动权给对方
		// 但是diff > 1的时候，如果从下一个状态开始，第一个move的player会输掉，那么就let m = diff
		// 否则的话，只要选择 diff - 1, 就可以扭转战局
		diff := a[i]
		if i > 0 {
			diff -= a[i-1]
		}
		if diff > 1 {
			return players[w]
		}
		w ^= 1
	}
	// not reachable
	return Alice
}
