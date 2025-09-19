package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	var buf bytes.Buffer
	for i := range res {
		for j := range res[i] {
			buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
		}
		buf.WriteString("\n")
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) (a [][]int, res [][]int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a = make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	res = solve(a)
	return
}

type pair struct {
	first  int
	second int
}

func cmp_pair(x pair, y pair) int {
	return cmp.Or(x.first-y.first, x.second-y.second)
}

func solve(a [][]int) [][]int {
	n, m := len(a), len(a[0])

	set := NewDSU(n * m)

	row := make([]pair, m)
	for i := range n {
		for j := range m {
			row[j] = pair{a[i][j], i*m + j}
		}
		slices.SortFunc(row, cmp_pair)
		for j := 0; j+1 < m; j++ {
			if row[j].first == row[j+1].first {
				set.Union(row[j].second, row[j+1].second)
			}
		}
	}

	col := make([]pair, n)
	for j := range m {
		for i := range n {
			col[i] = pair{a[i][j], i*m + j}
		}
		slices.SortFunc(col, cmp_pair)
		for i := 0; i+1 < n; i++ {
			if col[i].first == col[i+1].first {
				set.Union(col[i].second, col[i+1].second)
			}
		}
	}

	adj := make([][]int, n*m)
	deg := make([]int, n*m)

	connect := func(x int, y int) {
		adj[x] = append(adj[x], y)
		deg[y]++
	}

	for i := range n {
		for j := range m {
			row[j] = pair{a[i][j], i*m + j}
		}
		slices.SortFunc(row, cmp_pair)
		for j := 0; j+1 < m; j++ {
			if row[j].first < row[j+1].first {
				connect(set.Find(row[j].second), set.Find(row[j+1].second))
			}
		}
	}

	for j := range m {
		for i := range n {
			col[i] = pair{a[i][j], i*m + j}
		}
		slices.SortFunc(col, cmp_pair)
		for i := 0; i+1 < n; i++ {
			if col[i].first < col[i+1].first {
				connect(set.Find(col[i].second), set.Find(col[i+1].second))
			}
		}
	}

	que := make([]int, len(adj))
	var head, tail int
	val := make([]int, len(adj))

	for i := range len(adj) {
		j := set.Find(i)
		if i == j && deg[i] == 0 {
			val[i] = 1
			que[head] = i
			head++
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		for _, v := range adj[u] {
			val[v] = max(val[v], val[u]+1)
			deg[v]--
			if deg[v] == 0 {
				que[head] = v
				head++
			}
		}
	}

	ans := make([][]int, n)
	for i := range n {
		ans[i] = make([]int, m)
		for j := range m {
			ans[i][j] = val[set.Find(i*m+j)]
		}
	}

	return ans
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
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
	return true
}
