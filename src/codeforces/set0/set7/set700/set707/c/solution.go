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
		fmt.Println("-1")
	} else {
		fmt.Println(res[0], res[1])
	}
}

func drive(reader *bufio.Reader) (n int, res []int) {
	fmt.Fscan(reader, &n)
	res = solve(n)
	return
}

func solve(w int) []int {
	if w <= 2 {
		return nil
	}
	if w%2 == 0 {
		m := w / 2
		n := 1
		return []int{m*m - n*n, m*m + n*n}
	}
	// w是奇数
	// let m - n = 1, m + n = w
	m := (w + 1) / 2
	n := m - 1
	return []int{2 * m * n, m*m + n*n}
}
