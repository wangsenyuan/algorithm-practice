package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a := make([]int, 6)
	for i := 0; i < 6; i++ {
		fmt.Fscanf(reader, "%d", &a[i])
	}
	res := solve(a)
	fmt.Println(res)
}

func solve(a []int) int {
	s := (a[0] + a[1] + a[2]) * (a[0] + a[1] + a[2])

	s -= a[0] * a[0]
	s -= a[2] * a[2]
	s -= a[4] * a[4]
	return s
}
