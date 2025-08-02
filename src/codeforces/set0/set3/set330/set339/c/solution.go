package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans, _, _ := process(reader)
	if len(ans) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	s := fmt.Sprintf("%v", ans)
	fmt.Println(s[1 : len(s)-1])
}

func process(reader *bufio.Reader) ([]int, int, string) {
	w, _ := reader.ReadString('\n')
	w = strings.TrimSpace(w)
	var m int
	fmt.Fscan(reader, &m)
	return solve(w, m), m, w
}

type data struct {
	ok bool
	w  int
	d  int
}

func solve(weights string, m int) []int {
	dp := make([][][]data, m)
	for i := range m {
		dp[i] = make([][]data, 11)
		for w := range 11 {
			dp[i][w] = make([]data, 11)
			for d := range 10 {
				dp[i][w][d] = data{false, -1, -1}
			}
		}
	}

	for i := range 10 {
		if weights[i] == '1' {
			// 一开始放在左边的情况
			dp[0][i+1][i+1] = data{true, -1, -1}
		}
	}

	for i := range m - 1 {
		for d := 1; d <= 10; d++ {
			for w := 1; w <= 10; w++ {
				if !dp[i][d][w].ok {
					continue
				}
				// 另外一边放置的时候，必须要比d大，才能保证向另外一边倾斜
				for nw := d + 1; nw <= 10; nw++ {
					if weights[nw-1] == '0' || nw == w || nw-d > 10 {
						continue
					}
					if !dp[i+1][nw-d][nw].ok {
						dp[i+1][nw-d][nw] = data{true, w, d}
					}
				}
			}
		}
	}

	construct := func(d int, w int) []int {
		var res []int
		for i := m - 1; i >= 0; i-- {
			res = append(res, w)
			d1, w1 := dp[i][d][w].d, dp[i][d][w].w
			d, w = d1, w1
		}
		slices.Reverse(res)
		return res
	}

	for d := 1; d <= 10; d++ {
		for w := 1; w <= 10; w++ {
			if dp[m-1][d][w].ok {
				return construct(d, w)
			}
		}
	}

	return nil
}
