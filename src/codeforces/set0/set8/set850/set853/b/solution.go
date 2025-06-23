package main

import (
	"bufio"
	"fmt"
	"os"
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
	n, m, k := readThreeNums(reader)
	flights := make([][]int, m)
	for i := range m {
		flights[i] = readNNums(reader, 4)
	}
	return solve(n, k, flights)
}

const inf = 1 << 60

func solve(n int, k int, flights [][]int) int {
	// dp[i]表示在第i天，将所有人聚到首都的最小cost
	var md int

	for _, cur := range flights {
		md = max(md, cur[0])
	}

	arrive := make([][]int, md+1)
	leave := make([][]int, md+1)

	for i, cur := range flights {
		d, f := cur[0], cur[1]
		if f == 0 {
			// 离开首都
			leave[d] = append(leave[d], i)
		} else {
			arrive[d] = append(arrive[d], i)
		}
	}

	// 到达首都的最优解
	best := make([]int, n+1)
	for i := 0; i <= n; i++ {
		best[i] = inf
	}
	var sum int
	var cnt int

	dp := make([]int, md+1)

	for d := range md + 1 {
		for _, i := range arrive[d] {
			flight := flights[i]
			f, c := flight[1], flight[3]
			if c < best[f] {
				if best[f] == inf {
					cnt++
				} else {
					sum -= best[f]
				}
				sum += c
				best[f] = c
			}
		}
		if cnt < n {
			dp[d] = inf
		} else {
			dp[d] = sum
		}
	}

	// 处理离开的事件
	sum = 0
	cnt = 0
	for i := range n + 1 {
		best[i] = inf
	}

	ans := inf

	for d := md; d > 0; d-- {
		for _, i := range leave[d] {
			flight := flights[i]
			t, c := flight[2], flight[3]
			if c < best[t] {
				if best[t] == inf {
					cnt++
				} else {
					sum -= best[t]
				}
				sum += c
				best[t] = c
			}
		}
		if cnt == n {
			// 可以在d天后离开
			if d-k-1 > 0 {
				ans = min(ans, sum+dp[d-k-1])
			}
		}
	}

	if ans == inf {
		return -1
	}
	return ans
}
