package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	if len(res) == 0 {
		fmt.Println("-1")
	} else {
		s := fmt.Sprintf("%v", res)
		fmt.Println(s[1 : len(s)-1])
	}
}

func drive(reader *bufio.Reader) (a [][]int, res []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a = make([][]int, n)

	for i := 0; i < n; i++ {
		a[i] = make([]int, m)

		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	res = solve(a)
	return
}

func solve(a [][]int) []int {
	// n := len(a)
	m := len(a[0])

	type pair struct {
		first  int
		second int
	}

	// 相等的话，还是要搞个特殊的点才行

	arr := make([]pair, m)

	deg := make([]int, m)

	g := make([][]int, m)

	doIt := func(row []int) {
		var pos int
		for i := range m {
			if row[i] > 0 {
				arr[pos] = pair{row[i], i}
				pos++
			}
		}
		slices.SortFunc(arr[:pos], func(x, y pair) int {
			return x.first - y.first
		})

		prev := -1
		for i := 0; i < pos; {
			j := i
			for i < pos && arr[i].first == arr[j].first {
				if prev >= 0 {
					g[prev] = append(g[prev], arr[i].second)
					deg[arr[i].second]++
				}
				i++
			}

			if i-j > 1 {
				id := len(g)
				g = append(g, make([]int, 0, 1))
				deg = append(deg, 0)
				for k := j; k < i; k++ {
					g[arr[k].second] = append(g[arr[k].second], id)
					deg[id]++
				}
				prev = id
			} else {
				prev = arr[j].second
			}
		}
	}

	for _, row := range a {
		doIt(row)
	}
	// 然后必须不能形成环
	que := make([]int, len(g))
	marked := make([]bool, len(g))
	var head, tail int
	for i := range m {
		if deg[i] == 0 {
			que[head] = i
			head++
			marked[i] = true
		}
	}
	for tail < head {
		u := que[tail]
		tail++
		for _, v := range g[u] {
			deg[v]--
			if deg[v] == 0 {
				marked[v] = true
				que[head] = v
				head++
			}
		}
	}
	var ans []int
	for i := range head {
		if que[i] < m {
			ans = append(ans, que[i]+1)
		}
	}
	if len(ans) < m {
		return nil
	}
	return ans
}
