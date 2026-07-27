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
	return solve(n)
}

func solve(n int) string {
	if n%2 == 1 {
		return "black"
	}
	return "white\n1 2"
}
