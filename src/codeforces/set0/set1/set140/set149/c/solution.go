package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	first, second := solve(a)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	output := func(arr []int) {
		fmt.Fprintln(writer, len(arr))
		for _, i := range arr {
			fmt.Fprintf(writer, "%d ", i)
		}
		fmt.Fprintln(writer)
	}
	output(first)
	output(second)
}

type player struct {
	score int
	id    int
}

func solve(a []int) (first []int, second []int) {
	n := len(a)
	players := make([]player, n)
	for i := range n {
		players[i] = player{score: a[i], id: i}
	}
	slices.SortFunc(players, func(i, j player) int {
		return j.score - i.score
	})

	sum := make([]int, 2)
	for i := 0; i < n; i += 2 {
		if i == n-1 {
			// 最后一个人
			if sum[0] > sum[1] {
				second = append(second, players[i].id+1)
			} else {
				first = append(first, players[i].id+1)
			}
		} else {
			d := i / 2
			if d%2 == 0 {
				sum[0] += players[i].score
				first = append(first, players[i].id+1)
				sum[1] += players[i+1].score
				second = append(second, players[i+1].id+1)
			} else {
				sum[1] += players[i].score
				second = append(second, players[i].id+1)
				sum[0] += players[i+1].score
				first = append(first, players[i+1].id+1)
			}
		}
	}
	return
}
