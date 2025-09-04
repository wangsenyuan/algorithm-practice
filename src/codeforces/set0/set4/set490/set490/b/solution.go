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
	var n int
	fmt.Fscan(reader, &n)
	lines := make([][]int, n)
	for i := 0; i < n; i++ {
		lines[i] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &lines[i][0], &lines[i][1])
	}
	return solve(n, lines)
}

func solve(n int, lines [][]int) []int {
	res := make([]int, n)
	next := make(map[int]int)
	prev := make(map[int]int)
	var first, second int
	for _, cur := range lines {
		if cur[0] == 0 {
			second = cur[1]
		} else {
			next[cur[0]] = cur[1]
		}
		prev[cur[1]] = cur[0]
	}

	for k := range next {
		if _, ok := prev[k]; !ok {
			first = k
			break
		}
	}

	for i := 0; i < n; i += 2 {
		res[i] = first
		first = next[first]
	}

	for i := 1; i < n; i += 2 {
		res[i] = second
		second = next[second]
	}

	return res
}
