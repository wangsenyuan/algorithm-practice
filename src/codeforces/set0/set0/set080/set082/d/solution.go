package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	score, res, _ := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", score))
	for _, cur := range res {
		for _, x := range cur {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) (score int, res [][]int, a []int) {
	n := readNum(reader)
	a = readNNums(reader, n)
	score, res = solve(a)
	return
}

type state struct {
	val int
	u   int
	v   int
	x   int
}

const inf = 1 << 60

func solve(a []int) (int, [][]int) {
	n := len(a)
	if n == 1 {
		return a[0], [][]int{{1}}
	}
	if n == 2 {
		return max(a[0], a[1]), [][]int{{1}, {2}}
	}
	// n > 2
	dp := make([][]state, (n+1)/2+1)
	for i := range dp {
		dp[i] = make([]state, n+1)
		for j := range dp[i] {
			dp[i][j] = state{val: inf, u: -1, v: -1, x: -1}
		}
	}
	clear(dp[0])
	// dp[j][i] 表示处理完前j轮后，剩余的是i
	dp[1][0] = state{max(a[1], a[2]), 1, 2, 0}
	dp[1][1] = state{max(a[0], a[2]), 0, 2, 1}
	dp[1][2] = state{max(a[0], a[1]), 0, 1, 2}
	for j := 1; j+1 < len(dp); j++ {
		w := 2*j + 1
		// 这次要处理w, w+1
		if w == n {
			// n = 3
			for i := 0; i < n; i++ {
				if dp[j][i].val+a[i] < dp[j+1][n].val {
					dp[j+1][n].val = dp[j][i].val + a[i]
					dp[j+1][n].u = i
					dp[j+1][n].x = i
				}
			}
		} else if w+1 == n {
			// n = 4
			for i := 0; i < n; i++ {
				if dp[j][i].val+max(a[i], a[w]) < dp[j+1][n].val {
					dp[j+1][n] = state{dp[j][i].val + max(a[i], a[w]), i, w, i}
				}
			}
		} else {
			for i := 0; i < n; i++ {
				if dp[j][i].val+max(a[w], a[w+1]) < dp[j+1][i].val {
					dp[j+1][i] = state{dp[j][i].val + max(a[w], a[w+1]), w, w + 1, i}
				}
				if dp[j][i].val+max(a[i], a[w+1]) < dp[j+1][w].val {
					dp[j+1][w] = state{dp[j][i].val + max(a[i], a[w+1]), i, w + 1, i}
				}
				if dp[j][i].val+max(a[i], a[w]) < dp[j+1][w+1].val {
					dp[j+1][w+1] = state{dp[j][i].val + max(a[i], a[w]), i, w, i}
				}
			}
		}
	}
	var res [][]int
	pos := n
	for j := len(dp) - 1; j > 0; j-- {
		s := dp[j][pos]
		cur := []int{s.u + 1}
		if s.v != -1 {
			cur = append(cur, s.v+1)
		}
		res = append(res, cur)
		pos = s.x
	}

	reverse(res)

	return dp[len(dp)-1][n].val, res
}

func reverse(arr [][]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
