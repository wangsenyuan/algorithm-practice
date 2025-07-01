package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans, res, _, _ := process(reader)
	fmt.Println(ans)
	if ans < 0 {
		return
	}
	for _, cur := range res {
		fmt.Println(cur[0], cur[1])
	}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) (ans int, res [][]int, pos []int, grid []string) {
	n := readNum(reader)
	pos = readNNums(reader, 3)
	grid = make([]string, n)
	for i := range n {
		grid[i] = readString(reader)
	}
	ans, res = solve(n, pos, grid)
	return
}

type key struct {
	a int
	b int
	c int
}

func newKey(a, b, c int) key {
	tmp := []int{a, b, c}
	sort.Ints(tmp)
	return key{tmp[0], tmp[1], tmp[2]}
}

func solve(n int, pos []int, grid []string) (int, [][]int) {

	type data struct {
		val int
		key
	}

	dp := make([][][]data, n)
	for i := range n {
		dp[i] = make([][]data, n)
		for j := range n {
			dp[i][j] = make([]data, n)
			for k := range n {
				dp[i][j][k] = data{-1, key{-1, -1, -1}}
			}
		}
	}
	sort.Ints(pos)
	dp[pos[0]-1][pos[1]-1][pos[2]-1].val = 0

	que := make([]key, 0, n*n*n)
	que = append(que, key{pos[0] - 1, pos[1] - 1, pos[2] - 1})
	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		a, b, c := cur.a, cur.b, cur.c
		for d := range n {
			if grid[a][d] == grid[b][c] && d != b && d != c {
				nk := newKey(d, b, c)
				if dp[nk.a][nk.b][nk.c].val < 0 {
					dp[nk.a][nk.b][nk.c] = data{dp[a][b][c].val + 1, key{a, b, c}}
					que = append(que, nk)
				}
			}
			if grid[b][d] == grid[a][c] && d != a && d != c {
				nk := newKey(a, d, c)
				if dp[nk.a][nk.b][nk.c].val < 0 {
					dp[nk.a][nk.b][nk.c] = data{dp[a][b][c].val + 1, key{a, b, c}}
					que = append(que, nk)
				}
			}
			if grid[c][d] == grid[a][b] && d != a && d != b {
				nk := newKey(a, b, d)
				if dp[nk.a][nk.b][nk.c].val < 0 {
					dp[nk.a][nk.b][nk.c] = data{dp[a][b][c].val + 1, key{a, b, c}}
					que = append(que, nk)
				}
			}
		}
	}

	if dp[0][1][2].val < 0 {
		return -1, nil
	}
	res := make([][]int, dp[0][1][2].val)
	a, b, c := 0, 1, 2
	for i := len(res) - 1; i >= 0; i-- {
		x, y, z := dp[a][b][c].key.a, dp[a][b][c].key.b, dp[a][b][c].key.c
		u, v := findDiff(a, b, c, x, y, z)
		res[i] = []int{u + 1, v + 1}
		a, b, c = x, y, z
	}
	return dp[0][1][2].val, res
}

func findDiff(a, b, c, x, y, z int) (int, int) {
	freq := make(map[int]int)
	freq[a]++
	freq[b]++
	freq[c]++
	freq[x]--
	freq[y]--
	freq[z]--
	arr := make([]int, 2)
	for k, v := range freq {
		if v > 0 {
			arr[1] = k
		} else if v < 0 {
			arr[0] = k
		}
	}
	return arr[0], arr[1]
}
