package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, _, _ := process(reader)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	for _, e := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", e[0], e[1]))
	}
	buf.WriteTo(os.Stdout)
}

func process(reader *bufio.Reader) ([][]int, []int, int, int) {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	a := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, m, a), a, n, m
}

func solve(n int, m int, a []int) [][]int {
	k := len(a)
	if n == k {
		return nil
	}
	c := n - 1

	// 选中一个被标记的节点u，它只能连接到n-k个被标记的节点上
	// c * (c - 1) / 2 + n - k >= m 必须成立
	if c*(c-1)/2+n-k < m {
		return nil
	}

	id := make([]int, n)
	for i := range n {
		id[i] = i
	}

	marked := make([]bool, n)
	for _, v := range a {
		marked[v-1] = true
	}

	swapAt := func(i int, j int) {
		id[i], id[j] = id[j], id[i]
	}

	if !marked[id[n-1]] {
		for i := 0; i < n-1; i++ {
			if marked[id[i]] {
				swapAt(i, n-1)
				break
			}
		}
	}

	var res [][]int
	for i := 0; i < n-2; i++ {
		res = append(res, []int{id[i] + 1, id[i+1] + 1})
	}
	m -= n - 2
	var pw int
	for i := 0; i < n-2; i++ {
		if !marked[id[i]] {
			res = append(res, []int{id[i] + 1, id[n-1] + 1})
			m--
			pw = i
			break
		}
	}
	// 现在是一棵树
	for i := 0; i < n-2 && m > 0; i++ {
		for j := i + 2; j < n-1 && m > 0; j++ {
			res = append(res, []int{id[i] + 1, id[j] + 1})
			m--
		}
	}

	for i := 0; i < n-1 && m > 0; i++ {
		if i == pw || marked[id[i]] {
			continue
		}
		res = append(res, []int{id[i] + 1, id[n-1] + 1})
		m--
	}

	return res
}
