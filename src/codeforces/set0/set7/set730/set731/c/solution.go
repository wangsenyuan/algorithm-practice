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
	c := readNNums(reader, n)
	days := make([][]int, m)
	for i := 0; i < m; i++ {
		days[i] = readNNums(reader, 2)
	}
	return solve(m, k, c, days)
}

func solve(m int, k int, c []int, days [][]int) int {
	n := len(c)

	g := make([][]int, n)

	for _, cur := range days {
		u, v := cur[0]-1, cur[1]-1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	freq := make([]int, k+1)

	// n := len(c)

	marked := make([]int, n)

	var mx int
	var dfs func(u int, d int, x int) int
	dfs = func(u int, d int, x int) int {
		marked[u] = x
		freq[c[u]] += d
		if freq[c[u]] > mx {
			mx = freq[c[u]]
		}
		var res int

		for _, v := range g[u] {
			if marked[v] != x {
				res += dfs(v, d, x)
			}
		}
		return res + 1
	}

	var res int
	for u := 0; u < n; u++ {
		if marked[u] == 0 {
			mx = 0
			tmp := dfs(u, 1, 2*(u+1))
			res += tmp - mx
			dfs(u, -1, 2*(u+1)+1)
		}
	}

	return res
}
