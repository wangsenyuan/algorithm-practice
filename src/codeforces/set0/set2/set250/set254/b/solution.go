package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, _ := os.Open("input.txt")
	reader := bufio.NewReader(r)

	w, _ := os.Create("output.txt")

	writer := bufio.NewWriter(w)

	defer writer.Flush()

	res := drive(reader)
	fmt.Fprintln(writer, res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	tasks := make([][]int, n)
	for i := range n {
		tasks[i] = make([]int, 4)
		fmt.Fscan(reader, &tasks[i][0], &tasks[i][1], &tasks[i][2], &tasks[i][3])
	}
	return solve(tasks)
}

func solve(tasks [][]int) int {
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	sum := make([]int, 13)
	for i := range 12 {
		sum[i+1] = sum[i] + days[i]
	}
	diff := make([]int, 400)

	for _, task := range tasks {
		m, d, p, t := task[0], task[1], task[2], task[3]
		t1 := max(0, sum[m-1]+d-t)
		diff[t1] += p
		diff[sum[m-1]+d] -= p
	}

	var res int

	for d := 0; d <= sum[12]; d++ {
		if d > 0 {
			diff[d] += diff[d-1]
		}
		res = max(res, diff[d])
	}

	return res
}
