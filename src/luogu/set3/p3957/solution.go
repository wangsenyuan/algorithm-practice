package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n, d, k := readThreeNums(reader)
	items := make([][]int, n)
	for i := 0; i < n; i++ {
		items[i] = readNNums(reader, 2)
	}
	return solve(d, k, items)
}

const inf = 1 << 60

type pair struct {
	first  int
	second int
}

func solve(d int, k int, items [][]int) int {
	// x[i] < x[i+1]
	if items[0][0] != 0 {
		items = append([][]int{{0, 0}}, items...)
	}
	n := len(items)

	que := make([]pair, n)
	dp := make([]int, n)
	check := func(g int) bool {
		clear(dp)
		clear(que)
		l, r := max(1, d-g), d+g
		var head, tail int
		for i, j := n-1, n-1; i >= 0; i-- {
			for j > i && items[j][0]-items[i][0] >= l {
				for head > tail && dp[j] >= que[head-1].first {
					head--
				}
				que[head] = pair{dp[j], j}
				head++
				j--
			}
			for head > tail && items[que[tail].second][0]-items[i][0] > r {
				tail++
			}
			if head == tail {
				dp[i] = items[i][1]
			} else {
				// 可以在i这里结束
				dp[i] = items[i][1] + max(0, que[tail].first)
			}
		}
		return dp[0] >= k
	}

	if !check(inf) {
		return -1
	}

	return sort.Search(inf, check)
}
