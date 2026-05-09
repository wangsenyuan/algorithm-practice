package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.12f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	x := make([]float64, n)
	y := make([]float64, n)
	var sx, sy float64
	for i := 0; i < n; i++ {
		cur := readNNums(reader, 2)
		x[i] = float64(cur[0])
		y[i] = float64(cur[1])
		sx += x[i]
		sy += y[i]
	}
	for i := 0; i < n; i++ {
		x[i] /= sx
		y[i] /= sy
	}
	return solve(n, edges, x, y)
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

func solve(n int, edges [][]int, enter []float64, exit []float64) float64 {
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	parent := make([]int, n)
	for i := range n {
		parent[i] = -1
	}
	size := make([]int, n)
	sumEnter := make([]float64, n)
	sumExit := make([]float64, n)

	var dfs func(int, int)
	dfs = func(u int, p int) {
		parent[u] = p
		size[u] = 1
		sumEnter[u] = enter[u]
		sumExit[u] = exit[u]
		for _, v := range g[u] {
			if v == p {
				continue
			}
			dfs(v, u)
			size[u] += size[v]
			sumEnter[u] += sumEnter[v]
			sumExit[u] += sumExit[v]
		}
	}
	dfs(0, 0)

	var ans float64
	for u := range n {
		var totalExit, totalExitSize float64
		for _, v := range g[u] {
			s, ex, ey := component(u, v, parent, size, sumEnter, sumExit)
			ans += ex * (1 - ey)
			ans += enter[u] * ey * float64(n-1-s)
			totalExit += ey
			totalExitSize += ey * float64(s)
		}

		for _, p := range g[u] {
			s, ex, ey := component(u, p, parent, size, sumEnter, sumExit)
			restExit := totalExit - ey
			restExitSize := totalExitSize - ey*float64(s)
			ans += ex * (float64(n-1-s)*restExit - restExitSize)
		}
	}

	return ans
}

func component(u int, v int, parent []int, size []int, enter []float64, exit []float64) (int, float64, float64) {
	n := len(parent)
	if parent[v] == u {
		return size[v], enter[v], exit[v]
	}
	return n - size[u], 1 - enter[u], 1 - exit[u]
}
