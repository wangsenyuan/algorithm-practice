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
		var a, b int
		fmt.Fscan(reader, &a, &b)
		res := solve(a, b)
		fmt.Fprintf(writer, "%.10f\n", res)
	}
}

func solve(a int, b int) float64 {
	tot := float64(a) * float64(b*2)
	if b == 0 {
		return 1.0
	}
	if a == 0 {
		// 一半的区域？
		return 0.5
	}

	// 第3象限
	s := float64(a * b)

	// y = 4 * x的上方
	if b*4 >= a {
		// 一个三角形
		s += float64(a) * (float64(a) / 4) / 2
	} else {
		// 一个梯形
		s += float64(a+a-4*b) * float64(b) / 2
	}

	return s / tot
}
