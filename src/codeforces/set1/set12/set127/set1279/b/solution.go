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
	for range tc {
		res := drive(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) int {
	var n, s int
	fmt.Fscan(reader, &n, &s)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, s)
}

func solve(a []int, s int) int {
	var sum int
	best := -1
	for i, v := range a {
		sum += v
		if best < 0 || v > a[best] {
			best = i
		}
		if sum > s && sum-a[best] <= s {
			return best + 1
		}
	}
	return 0
}
