package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	incorrect, correct := solve(n)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	if len(incorrect) > 0 {
		for _, edge := range incorrect {
			fmt.Fprintln(writer, edge[0], edge[1])
		}
	} else {
		fmt.Fprintln(writer, "-1")
	}

	for _, edge := range correct {
		fmt.Fprintln(writer, edge[0], edge[1])
	}
}

func solve(n int) (incorrect [][]int, correct [][]int) {
	// 那种中间两个点，其他的围一圈，
	// 好像至少要6个点，才能构造一个错误的结构出来
	if n >= 6 {
		// 最佳答案是(1, 2)
		incorrect = append(incorrect, []int{1, 2})

		for i := 3; i <= n; i++ {
			if i&1 == 1 {
				incorrect = append(incorrect, []int{i, 1})
			} else {
				incorrect = append(incorrect, []int{i, 2})
			}
		}
	}

	for i := 1; i < n; i++ {
		correct = append(correct, []int{i, i + 1})
	}

	return
}
