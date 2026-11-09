package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	var k1 int
	fmt.Fscan(reader, &k1)
	a := make([]int, k1)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	var k2 int
	fmt.Fscan(reader, &k2)
	b := make([]int, k2)
	for i := range b {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

func solve(a []int, b []int) string {
	seen := make(map[string]bool)
	var cnt int
	for len(a) > 0 && len(b) > 0 {
		key := fmt.Sprint(a) + "|" + fmt.Sprint(b)
		if seen[key] {
			return "-1"
		}
		seen[key] = true

		cnt++
		u, v := a[0], b[0]
		a = a[1:]
		b = b[1:]
		if u > v {
			a = append(a, v, u)
		} else {
			b = append(b, u, v)
		}
	}

	if len(a) > 0 {
		return fmt.Sprintf("%d %d", cnt, 1)
	}
	return fmt.Sprintf("%d %d", cnt, 2)
}
