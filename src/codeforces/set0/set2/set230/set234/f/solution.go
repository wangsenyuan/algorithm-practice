package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, _ := os.Open("input.txt")
	defer r.Close()
	w, _ := os.Create("output.txt")
	defer w.Close()
	reader := bufio.NewReader(r)
	res := process(reader)
	fmt.Fprintln(w, res)
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
	n := readNum(reader)
	a, b := readTwoNums(reader)
	h := readNNums(reader, n)
	return solve(a, b, h)
}

const inf = 1 << 60

func solve(a int, b int, h []int) int {
	n := len(h)
	if a > b {
		a, b = b, a
	}
	dp := make([][]int, a+1)
	fp := make([][]int, a+1)
	for i := range dp {
		dp[i] = make([]int, 2)
		fp[i] = make([]int, 2)
		for j := range dp[i] {
			dp[i][j] = inf
			fp[i][j] = inf
		}
	}
	// dp[i][j]表示到目前为止，使用了i个红色，(sum - j)个蓝色，且上一个使用的是红色（j = 0)
	// 或蓝色(j = 1)， 所需要的最小代价
	if a >= h[0] {
		dp[h[0]][0] = 0
	}
	if b >= h[0] {
		dp[0][1] = 0
	}
	sum := h[0]
	for i := 1; i < n; i++ {
		x := h[i]

		for u := 0; u <= min(a, sum); u++ {
			// u is for red, v is for blue
			v := sum - u
			if dp[u][0] < inf {
				if u+x <= a {
					// 前后使用红色
					fp[u+x][0] = min(fp[u+x][0], dp[u][0])
				}
				if v+x <= b {
					fp[u][1] = min(fp[u][1], dp[u][0]+min(h[i-1], x))
				}
			}
			if dp[u][1] < inf {
				if u+x <= a {
					fp[u+x][0] = min(fp[u+x][0], dp[u][1]+min(h[i-1], x))
				}
				if v+x <= b {
					fp[u][1] = min(fp[u][1], dp[u][1])
				}
			}
		}
		sum += x
		for u := 0; u <= min(a, sum); u++ {
			for j := 0; j < 2; j++ {
				dp[u][j] = fp[u][j]
				fp[u][j] = inf
			}
		}
	}
	ans := inf
	for u := 0; u <= a; u++ {
		for j := 0; j < 2; j++ {
			ans = min(ans, dp[u][j])
		}
	}
	if ans == inf {
		return -1
	}
	return ans
}
