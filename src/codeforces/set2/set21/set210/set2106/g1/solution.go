package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	ask := func(arr []int) int {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("? 1 %d", len(arr)))
		for _, x := range arr {
			buf.WriteString(fmt.Sprintf(" %d", x))
		}
		fmt.Println(buf.String())
		return readNum(reader)
	}

	toggle := func(x int) {
		fmt.Printf("? 2 %d\n", x)
	}

	for range tc {
		n := readNum(reader)
		edges := make([][]int, n-1)
		for i := range n - 1 {
			edges[i] = readNNums(reader, 2)
		}

		res := solve(n, edges, ask, toggle)

		var buf bytes.Buffer
		buf.WriteString("!")

		for _, x := range res {
			buf.WriteString(fmt.Sprintf(" %d", x))
		}
		fmt.Println(buf.String())
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

func solve(n int, edges [][]int, ask func(arr []int) int, toggle func(int)) []int {
	arr := make([]int, n)
	for i := range n {
		arr[i] = i + 1
	}

	s1 := ask(arr[1:])
	toggle(1)
	s2 := ask(arr[1:])
	value := make([]int, n)

	root := 1
	if abs(s1-s2) != 2*(n-1) {
		// root != 1
		l, r := 1, n-1
		for l < r {
			mid := (l + r) / 2
			s1 := ask(arr[l : mid+1])
			toggle(1)
			s2 := ask(arr[l : mid+1])
			if abs(s1-s2) == 2*(mid-l+1) {
				// root肯定在后半段
				l = mid + 1
			} else {
				r = mid
			}
		}
		root = arr[l]
	}

	value[root-1] = ask([]int{root})

	for i := 1; i <= n; i++ {
		if i != root {
			tmp := ask([]int{i})
			value[i-1] = tmp - value[0]
			if root != 1 {
				value[i-1] -= value[root-1]
			}
		}
	}

	return value
}

func abs(num int) int {
	return max(num, -num)
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
