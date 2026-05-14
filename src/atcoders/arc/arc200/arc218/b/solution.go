package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) string {
	slices.Sort(a)
	n := len(a)

	player := []string{"Alice", "Bob"}

	var id int

	for i := 0; i < n; {
		j := i
		for i < n && a[i] == a[j] {
			i++
		}
		diff := a[j]
		if j > 0 {
			diff -= a[j-1]
		}
		if diff == 0 {
			// a[0] = 0
			if i-j > 1 || i == n {
				return player[id]
			}
			// i - j == 1
		} else {
			if diff > 1 {
				// 如果从0开始，Bob会输，那么就全部拿走
				// 如果从0开始，Bob会赢，那么就选择diff - 1
				return player[id]
			}
			// diff == 1
			id ^= 1
			// 现在轮到Bob了
			if i-j > 1 || i == n {
				// 也有反转的机会
				return player[id]
			}
		}

		id ^= 1
	}
	return player[id]
}
