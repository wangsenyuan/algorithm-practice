package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, line := range drive(reader) {
		fmt.Fprintln(writer, line)
	}
}

func drive(reader *bufio.Reader) []string {
	var n, m, k, t int
	fmt.Fscan(reader, &n, &m, &k, &t)
	waste := make([][]int, k)
	for i := range k {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		waste[i] = []int{a, b}
	}
	queries := make([][]int, t)
	for i := range t {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		queries[i] = []int{x, y}
	}
	return solve(n, m, waste, queries)
}

func solve(n, m int, waste [][]int, queries [][]int) []string {
	type data struct {
		r  int
		c  int
		id int
	}

	arr := make([]data, len(waste))
	for i := range waste {
		arr[i] = data{
			r:  waste[i][0] - 1,
			c:  waste[i][1] - 1,
			id: i,
		}
	}

	slices.SortFunc(arr, func(a data, b data) int {
		return cmp.Or(a.r-b.r, a.c-b.c)
	})

	todo := make([]data, len(queries))
	for i := range queries {
		todo[i] = data{
			r:  queries[i][0] - 1,
			c:  queries[i][1] - 1,
			id: i,
		}
	}
	slices.SortFunc(todo, func(a data, b data) int {
		return cmp.Or(a.r-b.r, a.c-b.c)
	})

	ans := make([]string, len(queries))
	for i := range ans {
		ans[i] = "Waste"
	}

	var foods = []string{"Carrots", "Kiwis", "Grapes"}

	var rowSum int
	var j int
	for i := 0; i < len(arr); {
		for j < len(todo) && todo[j].r < arr[i].r {
			pos := m*todo[j].r - rowSum + todo[j].c
			pos %= 3
			ans[todo[j].id] = foods[pos]
			j++
		}

		var row []data
		i1 := i
		for i < len(arr) && arr[i].r == arr[i1].r {
			row = append(row, arr[i])
			i++
		}
		tmp := len(row)
		var colSum int
		for j < len(todo) && todo[j].r == arr[i1].r {
			for len(row) > 0 && row[0].c < todo[j].c {
				colSum++
				row = row[1:]
			}
			if len(row) == 0 || todo[j].c < row[0].c {
				pos := m*arr[i1].r - rowSum + todo[j].c - colSum
				pos %= 3
				ans[todo[j].id] = foods[pos]
			}
			j++
		}
		rowSum += tmp
	}

	for j < len(todo) {
		pos := m*todo[j].r - rowSum + todo[j].c
		pos %= 3
		ans[todo[j].id] = foods[pos]
		j++
	}

	return ans
}
