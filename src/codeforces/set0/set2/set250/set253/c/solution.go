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
	a := readNNums(reader, n)
	pos := readNNums(reader, 4)
	return solve(a, pos)
}

type pair struct {
	first  int
	second int
}

const inf = 1 << 60

func solve(a []int, pos []int) int {
	// n := len(a)
	for i := range pos {
		pos[i]--
	}
	n := len(a)
	r1, c1 := pos[0], pos[1]
	r2, c2 := pos[2], pos[3]
	// dp[x][y]表示到达某个地方的最短距离
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, a[i]+1)
		for j := range a[i] + 1 {
			dp[i][j] = inf
		}
	}
	dp[r1][c1] = 0
	var que []pair
	que = append(que, pair{r1, c1})

	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		r, c := cur.first, cur.second
		if r == r2 && c == c2 {
			break
		}
		if r > 0 {
			// move up
			nc := min(c, a[r-1])
			if dp[r-1][nc] > dp[r][c]+1 {
				dp[r-1][nc] = dp[r][c] + 1
				que = append(que, pair{r - 1, nc})
			}
		}
		if r+1 < n {
			nc := min(c, a[r+1])
			if dp[r+1][nc] > dp[r][c]+1 {
				dp[r+1][nc] = dp[r][c] + 1
				que = append(que, pair{r + 1, nc})
			}
		}
		if c > 0 && dp[r][c-1] > dp[r][c]+1 {
			dp[r][c-1] = dp[r][c] + 1
			que = append(que, pair{r, c - 1})
		}
		if c+1 <= a[r] && dp[r][c+1] > dp[r][c]+1 {
			dp[r][c+1] = dp[r][c] + 1
			que = append(que, pair{r, c + 1})
		}
	}

	return dp[r2][c2]
}

func abs(a int) int {
	return max(a, -a)
}
