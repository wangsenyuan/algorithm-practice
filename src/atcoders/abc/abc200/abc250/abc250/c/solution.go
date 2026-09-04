package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	operations := make([]int, q)
	for i := range operations {
		fmt.Fscan(reader, &operations[i])
	}
	return solve(n, operations)
}

func solve(n int, operations []int) []int {
	pos := make([]int, n)
	arr := make([]int, n)
	for i := range pos {
		arr[i] = i
		pos[i] = i
	}

	play := func(x int) {
		i := pos[x]
		var j int
		if i == n-1 {
			// 如果它是最后一个
			j = i - 1
		} else {
			j = i + 1
		}
		y := arr[j]
		pos[x], pos[y] = j, i
		arr[i], arr[j] = y, x
	}

	for _, x := range operations {
		play(x - 1)
	}

	for i := range arr {
		arr[i]++
	}
	return arr
}
