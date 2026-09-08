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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	var s string
	fmt.Fscan(reader, &s)
	_ = n
	return solve(s)
}

func solve(s string) int {
	// 只有一个是2, 其他都是1
	x := s[0]
	var i int
	for i < len(s) && s[i] == x {
		i++
	}
	// 如果剩下的都和x不同, 就是2
	if i < len(s) {
		y := s[i]
		for i < len(s) && s[i] == y {
			i++
		}
		if i == len(s) {
			return 2
		}
	}

	return 1
}
