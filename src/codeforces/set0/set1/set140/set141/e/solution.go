package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, ok, res := process(reader)
	if !ok {
		fmt.Println(-1)
		return
	}
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func process(reader *bufio.Reader) (n int, edges [][]int, ok bool, res []int) {
	n, m := readTwoNums(reader)
	edges = make([][]int, m)
	for i := range m {
		s, _ := reader.ReadBytes('\n')
		var u, v, w int
		pos := readInt(s, 0, &u) + 1
		pos = readInt(s, pos, &v) + 1
		if s[pos] == 'S' {
			w++
		}
		edges[i] = []int{u, v, w}
	}
	ok, res = solve(n, edges)
	return
}

type pair struct {
	first  int
	second int
}

func solve(n int, edges [][]int) (ok bool, res []int) {
	if n%2 == 0 {
		// 肯定没有答案
		return false, nil
	}
	g := make([]map[int]int, n)
	for i := range n {
		g[i] = make(map[int]int)
	}
	todo := make([][]int, 2)
	set := NewDSU(n)
	cnt := make([]int, 2)

	connect := func(u int, v int, w int) {
		g[u][v] = w
		g[v][u] = w
	}

	disconect := func(u int, v int) {
		delete(g[u], v)
		delete(g[v], u)
	}

	for i, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		if u == v {
			// self-loop no use
			continue
		}
		if set.Union(u, v) {
			connect(u, v, i)
			cnt[w]++
		} else {
			todo[w] = append(todo[w], i)
		}
	}

	if set.sz != 1 {
		return false, nil
	}

	fa := make([]int, n)

	dep := make([]int, n)
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = p
		for v := range g[u] {
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(-1, 0)

	exp := 1
	if cnt[0] < cnt[1] {
		// 那么要使用w = 0的边去置换w=1的边
		exp = 0
	}

	for cnt[0] != cnt[1] && len(todo[exp]) > 0 {
		eid := todo[exp][0]
		e := edges[eid]
		todo[exp] = todo[exp][1:]
		a, b := e[0]-1, e[1]-1

		found := -1
		u, v := a, b
		for found < 0 && u >= 0 && v >= 0 && u != v {
			if dep[u] < dep[v] {
				u, v = v, u
			}
			p := fa[u]
			// 只要u != v, p >= 0 肯定成立
			if edges[g[u][p]][2] == 1^exp {
				found = u
				break
			}

			u = p
		}

		if found < 0 {
			continue
		}

		disconect(u, fa[u])
		connect(a, b, eid)
		dep[0] = 0
		dfs(-1, 0)
		cnt[exp]++
		cnt[1^exp]--
	}

	if cnt[0] != cnt[1] {
		return false, nil
	}

	var dfs2 func(p int, u int)
	dfs2 = func(p int, u int) {
		for v, eid := range g[u] {
			if v == p {
				continue
			}
			res = append(res, eid+1)
			dfs2(u, v)
		}
	}

	dfs2(-1, 0)

	return true, res
}

type DSU struct {
	arr []int
	cnt []int
	sz  int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt, n}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	this.sz--
	return true
}
