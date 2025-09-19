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
	_, _, res := drive(reader)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}

	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for _, row := range res {
		for _, v := range row {
			buf.WriteString(fmt.Sprintf("%d ", v))
		}
		buf.WriteString("\n")
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) (k int, a [][]int, res [][]int) {
	var n, m int
	fmt.Fscan(reader, &n, &m, &k)
	a = make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
	}
	for i := range n {
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	res = solve(k, a)
	return
}

type data struct {
	i int
	j int
	v int
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(k int, a [][]int) [][]int {
	n := len(a)
	m := len(a[0])

	// 得反过来处理
	arr := make([]data, n*m)
	for i := range n {
		for j := range m {
			arr[i*m+j] = data{i, j, a[i][j]}
		}
	}

	slices.SortFunc(arr, func(a, b data) int {
		return b.v - a.v
	})

	bfs := func(r int, c int, v int, w int) [][]int {
		ans := make([][]int, n)
		marked := make([][]bool, n)
		for i := range n {
			ans[i] = make([]int, m)
			marked[i] = make([]bool, m)
		}

		que := make([]int, n*m)
		var head, tail int
		que[head] = r*m + c
		head++
		marked[r][c] = true
		for tail < head {
			x, y := que[tail]/m, que[tail]%m
			tail++
			ans[x][y] = v
			for i := range 4 {
				p, q := x+dd[i], y+dd[i+1]
				if head < w && p >= 0 && p < n && q >= 0 && q < m && a[p][q] >= v && !marked[p][q] {
					que[head] = p*m + q
					head++
					marked[p][q] = true
				}
			}
		}

		return ans
	}

	set := NewDSU(n * m)

	for _, cur := range arr {
		r, c, v := cur.i, cur.j, cur.v
		for i := range 4 {
			x, y := r+dd[i], c+dd[i+1]
			if x >= 0 && x < n && y >= 0 && y < m && a[x][y] >= v {
				set.Union(r*m+c, x*m+y)
			}
		}

		if k%v == 0 && k/v <= set.cnt[set.Find(r*m+c)] {
			// found answer
			return bfs(r, c, v, k/v)
		}
	}

	return nil
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
