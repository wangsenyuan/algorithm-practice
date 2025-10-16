package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	solve(writer)
}

func solve(writer *bufio.Writer) {
	// thief在偶数移动后无法改变他所在位置的奇偶性
	// 所以先排除不能出现在左侧的奇数位
	// 然后再排除他不能出现在右侧的偶数位
	fmt.Fprintf(writer, "2000\n")
	for i := 1; i <= 1000; i++ {
		fmt.Fprintf(writer, "%d %d %d %d\n", i, 1, i, 2)
	}
	for i := 1000; i >= 1; i-- {
		fmt.Fprintf(writer, "%d %d %d %d\n", i, 1, i, 2)
	}
}
