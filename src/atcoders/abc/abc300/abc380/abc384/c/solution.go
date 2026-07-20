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

	for _, name := range drive(reader) {
		fmt.Fprintln(writer, name)
	}
}

func drive(reader *bufio.Reader) []string {
	scores := make([]int, 5)
	for i := range 5 {
		fmt.Fscan(reader, &scores[i])
	}
	return solve(scores)
}

func solve(scores []int) []string {
	type data struct {
		id    string
		score int
	}

	var arr []data

	for mask := 1; mask < (1 << 5); mask++ {
		var sum int
		var id []byte
		for i := range 5 {
			if (mask>>i)&1 == 1 {
				sum += scores[i]
				id = append(id, byte('A'+i))
			}
		}
		arr = append(arr, data{string(id), sum})
	}

	slices.SortFunc(arr, func(a, b data) int {
		return cmp.Or(b.score-a.score, cmp.Compare(a.id, b.id))
	})

	res := make([]string, len(arr))
	for i, v := range arr {
		res[i] = v.id
	}
	return res
}
