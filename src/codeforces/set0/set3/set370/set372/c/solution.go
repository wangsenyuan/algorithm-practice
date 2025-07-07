package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) int {
	n, m, d := readThreeNums(reader)
	events := make([][]int, m)
	for i := 0; i < m; i++ {
		events[i] = readNNums(reader, 3)
	}
	return solve(n, m, d, events)
}

const inf = 1 << 60

func solve(n int, m int, d int, events [][]int) int {
	slices.SortFunc(events, func(a, b []int) int {
		return a[2] - b[2]
	})

	dp := make([]int, n)

	type pair struct {
		first  int
		second int
	}

	que := make([]pair, n)

	for i, cur := range events {
		time_diff := cur[2]
		if i > 0 {
			time_diff -= events[i-1][2]
		} else {
			time_diff -= 1
		}

		var head, tail int

		for j, r := 0, 0; j < n; j++ {
			for r < n && r-j <= d*time_diff {
				for head > tail && que[head-1].first < dp[r] {
					head--
				}
				que[head] = pair{dp[r], r}
				head++
				r++
			}

			for tail < head && que[tail].second < j-d*time_diff {
				tail++
			}
			dp[j] = que[tail].first + cur[1] - abs(j+1-cur[0])
		}
	}

	return slices.Max(dp)
}

func abs(num int) int {
	return max(num, -num)
}
