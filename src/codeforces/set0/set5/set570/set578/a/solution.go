package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b int
	fmt.Fscan(reader, &a, &b)
	res := solve(a, b)
	fmt.Printf("%.10f\n", res)
}

func solve(a int, b int) float64 {
	if a < b {
		return -1
	}

	return float64((a + b)) / float64(2*int(float64(a+b)/float64(2*b)))
}
