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
	nums := readNNums(reader, 4)
	n, m, s, t := nums[0], nums[1], nums[2], nums[3]
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, m, s, t, edges)
}

func solve(n int, m int, s int, t int, edges [][]int) int {
	g := make([]map[int]struct{}, n)
	for i := range n {
		g[i] = make(map[int]struct{})
	}
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g[u][v] = struct{}{}
		g[v][u] = struct{}{}
	}

	que := make([]int, n)

	bfs := func(x int) []int {
		dist := make([]int, n)
		for i := range n {
			dist[i] = -1
		}
		dist[x] = 0
		var head, tail int
		que[head] = x
		head++
		for tail < head {
			u := que[tail]
			tail++
			for v := range g[u] {
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					que[head] = v
					head++
				}
			}
		}
		return dist
	}
	s--
	t--
	d1 := bfs(s)
	d2 := bfs(t)

	check := func(u int, v int) bool {
		if _, ok := g[u][v]; ok {
			return false
		}
		if d1[u] > d1[v] {
			u, v = v, u
		}

		return d1[u]+1+d2[v] >= d1[t]
	}

	var res int
	for u := 0; u < n; u++ {
		for v := u + 1; v < n; v++ {
			// 如果在u,v中间加一条线
			if check(u, v) {
				res++
			}
		}
	}

	return res
}
