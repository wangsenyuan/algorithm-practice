package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

const mod = 1_000_000_007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

type fenwick []int

func (t fenwick) add(p int, v int) {
	for p < len(t) {
		t[p] = add(t[p], v)
		p += p & -p
	}
}

func (t fenwick) get(p int) int {
	var res int
	for p > 0 {
		res = add(res, t[p])
		p -= p & -p
	}
	return res
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	row := make([]fenwick, n+1)
	col := make([]fenwick, n+1)
	for i := range n + 1 {
		row[i] = make(fenwick, n+1)
		col[i] = make(fenwick, n+1)
	}

	update := func(x int, y int, v int) {
		row[x].add(y, v)
		col[y].add(x, v)
	}

	update(1, 1, 1)
	// 定义 f[x][y] 表示表示第一个递增子序列以 x 结尾、第二个递增子序列以 y 结尾时，好子序列的数目
	// 为避免重复统计，规定 x >= y
	for _, v := range a {
		// 把 v 添加到第一个子序列的后面，必须满足 x <= v 且 y <= v
		// 枚举 y，f[v][y] += sum_{x <= v} f[x][y]
		for y := 1; y <= v; y++ {
			tmp := col[y].get(v)
			update(v, y, tmp)
		}
		// 把 v 添加到第二个子序列的后面，必须满足 y <= v < x，这里 v != x 从而避免重复统计
		// 枚举 x，f[x][v] += sum_{y <= v} f[x][y]
		for x := v + 1; x <= n; x++ {
			tmp := row[x].get(v)
			update(x, v, tmp)
		}
	}

	var ans int
	for _, f := range row {
		ans = add(ans, f.get(n))
	}
	return ans
}
