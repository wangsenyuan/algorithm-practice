package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		buf.WriteString(fmt.Sprintf("%d\n", drive(reader)))
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	writers := make([][]int, n)
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Fscan(reader, &a, &b, &c)
		writers[i] = []int{a, b, c}
	}
	return solve(writers)
}

func solve(writers [][]int) int {
	// 先最大化x, 然后最大化y
	var res int
	var x1 int
	var y2 int
	for _, cur := range writers {
		a, b, c := cur[0], cur[1], cur[2]
		x1 += min(a, b)
		y2 += min(b, c)

		d := min(a, b)
		res += d
		res += min(b-d, c)
	}

	return min(res/2, x1, y2)
}
