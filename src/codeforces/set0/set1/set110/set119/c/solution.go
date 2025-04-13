package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, _, res := process(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
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

func process(reader *bufio.Reader) (n int, m int, k int, tasks [][]int, res [][]int) {
	n, m, k = readThreeNums(reader)
	tasks = make([][]int, m)
	for i := range m {
		tasks[i] = readNNums(reader, 3)
	}
	res = solve(n, m, k, tasks)
	return
}

const inf = 1 << 60

type pair struct {
	first  int
	second int
}

func solve(n int, m int, k int, tasks [][]int) [][]int {
	type item struct {
		id int
		a  int
		b  int
		c  int
	}

	items := make([]item, m)
	for i := range m {
		items[i] = item{i, tasks[i][0], tasks[i][1], tasks[i][2]}
	}

	slices.SortFunc(items, func(x, y item) int {
		return x.c - y.c
	})

	dp := make([][][]int, m)
	from := make([][][]pair, m)
	for i := range m {
		dp[i] = make([][]int, 101)
		from[i] = make([][]pair, 101)
		for j := range 101 {
			dp[i][j] = make([]int, n+1)
			from[i][j] = make([]pair, n+1)
			for u := range n + 1 {
				dp[i][j][u] = -inf
				from[i][j][u] = pair{-1, -1}
			}
		}
	}

	for i := 0; i < m; {
		j := i
		for i < m && items[i].c == items[j].c {
			for u := range j {
				for v := 0; v <= items[u].b-items[u].a; v++ {
					w := v + items[u].a
					for d1 := 1; d1 < n; d1++ {
						if dp[u][v][d1] < 0 {
							break
						}
						if w+k >= items[i].a && w+k <= items[i].b && dp[u][v][d1]+w+k > dp[i][w+k-items[i].a][d1+1] {
							dp[i][w+k-items[i].a][d1+1] = dp[u][v][d1] + w + k
							from[i][w+k-items[i].a][d1+1] = pair{u, v}
						}
						if w*k >= items[i].a && w*k <= items[i].b && dp[u][v][d1]+w*k > dp[i][w*k-items[i].a][d1+1] {
							dp[i][w*k-items[i].a][d1+1] = dp[u][v][d1] + w*k
							from[i][w*k-items[i].a][d1+1] = pair{u, v}
						}
					}
				}
			}
			for v := 0; v <= items[i].b-items[i].a; v++ {
				dp[i][v][1] = v + items[i].a
			}
			i++
		}
	}

	ans := -inf
	pos := -1
	that_v := -1
	for i := range m {
		for v := range 101 {
			if dp[i][v][n] > ans {
				ans = dp[i][v][n]
				pos = i
				that_v = v
			}
		}
	}
	if ans < 0 {
		return nil
	}

	var res [][]int

	for len(res) < n {
		res = append(res, []int{items[pos].id + 1, that_v + items[pos].a})
		tmp := from[pos][that_v][n-len(res)+1]
		pos = tmp.first
		that_v = tmp.second
	}

	reverse(res)

	return res
}

func reverse(arr [][]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
