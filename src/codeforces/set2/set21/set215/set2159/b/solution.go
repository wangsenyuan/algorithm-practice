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

	for _, grid := range drive(reader) {
		for _, row := range grid {
			s := fmt.Sprintf("%v", row)
			fmt.Fprintln(writer, s[1:len(s)-1])
		}
	}
}

func drive(reader *bufio.Reader) [][][]int {
	var tc int
	fmt.Fscan(reader, &tc)
	res := make([][][]int, tc)
	for i := range tc {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		g := make([]string, n)
		for j := range n {
			fmt.Fscan(reader, &g[j])
		}
		res[i] = solve(g)
	}
	return res
}

func solve(g []string) [][]int {
	// TODO
	return nil
}
