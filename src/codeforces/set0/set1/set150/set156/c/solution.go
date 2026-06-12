package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1_000_000_007

var ways [101][2501]int

func init() {
	ways[0][0] = 1
	for n := 1; n <= 100; n++ {
		for sum := 0; sum <= 2500; sum++ {
			for c := 0; c <= 25; c++ {
				if sum >= c {
					ways[n][sum] = add(ways[n][sum], ways[n-1][sum-c])
				}
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	tc := readNum(reader)
	for range tc {
		s := readString(reader)
		fmt.Fprintln(writer, solve(s))
	}
}

func readNum(reader *bufio.Reader) int {
	var x int
	fmt.Fscan(reader, &x)
	return x
}

func readString(reader *bufio.Reader) string {
	var s string
	fmt.Fscan(reader, &s)
	return s
}

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func solve(s string) int {
	n := len(s)
	if n == 1 {
		return 0
	}
	sum := 0
	for i := range n {
		sum += int(s[i] - 'a')
	}
	return sub(ways[n][sum], 1)
}
