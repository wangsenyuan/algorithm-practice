package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, b := range res {
		if b {
			writer.WriteString("YES\n")
		} else {
			writer.WriteString("NO\n")
		}
	}
}

func drive(reader *bufio.Reader) []bool {
	var n int
	fmt.Fscan(reader, &n)
	queries := make([][]int, n)
	for i := 0; i < n; i++ {
		var c string
		var x, y int
		fmt.Fscan(reader, &c, &x, &y)
		if c == "+" {
			queries[i] = []int{1, x, y}
		} else {
			queries[i] = []int{2, x, y}
		}
	}
	return solve(queries)
}

type pair struct {
	first  int
	second int
}

func solve(queries [][]int) []bool {
	best := []int{0, 0}
	var ans []bool
	for i := 0; i < len(queries); i++ {
		if queries[i][0] == 1 {
			x, y := queries[i][1], queries[i][2]
			x, y = min(x, y), max(x, y)
			best[0] = max(best[0], x)
			best[1] = max(best[1], y)
		} else {
			h, w := queries[i][1], queries[i][2]
			if best[0] <= h && best[1] <= w || best[0] <= w && best[1] <= h {
				ans = append(ans, true)
			} else {
				ans = append(ans, false)
			}
		}
	}

	return ans
}
