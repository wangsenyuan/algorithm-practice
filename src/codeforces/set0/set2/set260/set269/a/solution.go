package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	boxes := make([][]int, n)
	for i := range n {
		boxes[i] = make([]int, 2)
		fmt.Fscan(reader, &boxes[i][0], &boxes[i][1])
	}
	return solve(boxes)
}

func solve(boxes [][]int) int {
	var mk int
	var p int
	for _, cur := range boxes {
		k, a := cur[0], cur[1]
		mk = max(mk, k)
		m, s := 0, 1
		for s < a {
			s *= 4
			m++
		}
		p = max(p, k+m)
	}

	if p == mk {
		p++
	}

	return p
}
