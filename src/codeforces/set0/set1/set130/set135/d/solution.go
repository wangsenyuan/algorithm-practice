package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(a []string) int {
	dsu := getComponents(a)
	n := len(a)
	m := len(a[0])

	ok := make([]bool, n*m)
	for i := range n * m {
		ok[i] = true
	}
	for i := range n {
		if a[i][0] == '0' {
			w := dsu.Find(i * m)
			ok[w] = false
		}
		if a[i][m-1] == '0' {
			w := dsu.Find(i*m + m - 1)
			ok[w] = false
		}
	}
	for j := range m {
		if a[0][j] == '0' {
			w := dsu.Find(j)
			ok[w] = false
		}
		if a[n-1][j] == '0' {
			w := dsu.Find((n-1)*m + j)
			ok[w] = false
		}
	}

	que := make([]int, n*m)
	// for 0
	marked := make([]bool, n*m)
	// for 1
	vis := make([]int, n*m)

	bfs := func(sr int, sc int) int {
		var head, tail int
		que[head] = sr*m + sc
		head++

		vis[sr*m+sc]++

		for tail < head {
			r, c := que[tail]/m, que[tail]%m
			tail++

			for i := range 4 {
				nr, nc := r+dd[i], c+dd[i+1]
				if nr >= 0 && nr < n && nc >= 0 && nc < m && vis[nr*m+nc] == 1 {
					vis[nr*m+nc]++
					que[head] = nr*m + nc
					head++
				}
			}
		}

		return head
	}

	play := func(no int) int {
		var head, tail int
		que[head] = no
		head++
		var todo [][]int
		for tail < head {
			r, c := que[tail]/m, que[tail]%m
			tail++
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					if dr == 0 && dc == 0 {
						continue
					}
					nr, nc := r+dr, c+dc
					if nr >= 0 && nr < n && nc >= 0 && nc < m {
						if a[nr][nc] == '0' {
							if !marked[nr*m+nc] {
								marked[nr*m+nc] = true
								que[head] = nr*m + nc
								head++
							}
						} else {
							// 这些是边
							if vis[nr*m+nc] == 0 {
								vis[nr*m+nc] = 1
								todo = append(todo, []int{nr, nc})
							}
						}
					}
				}
			}
		}
		// 这些边，必须是连在一起的
		ok := bfs(todo[0][0], todo[0][1]) == len(todo)
		if ok {
			for _, cur := range todo {
				r, c := cur[0], cur[1]
				var cnt int
				for i := range 4 {
					nr, nc := r+dd[i], c+dd[i+1]
					if nr >= 0 && nr < n && nc >= 0 && nc < m && vis[nr*m+nc] > 0 {
						cnt++
					}
				}
				if cnt > 2 {
					ok = false
					break
				}
			}
		}
		for _, cur := range todo {
			vis[cur[0]*m+cur[1]] = 0
		}
		if ok {
			return len(todo)
		}
		return 0
	}

	var ans int

	for i := range n {
		for j := range m {
			if a[i][j] == '1' {
				if i > 0 && j > 0 && a[i-1][j-1] == '1' && a[i-1][j] == '1' && a[i][j-1] == '1' {
					ans = max(ans, 4)
				}
			} else {
				no := i*m + j
				if dsu.Find(no) == no && ok[no] {
					ans = max(ans, play(no))
				}
			}
		}
	}

	return ans
}

func getComponents(a []string) *DSU {
	n := len(a)
	m := len(a[0])
	dsu := NewDSU(n * m)
	for i := range n {
		for j := range m {
			if a[i][j] == '0' {
				if i > 0 && a[i-1][j] == '0' {
					dsu.Union(i*m+j, (i-1)*m+j)
				}
				if j > 0 && a[i][j-1] == '0' {
					dsu.Union(i*m+j, i*m+j-1)
				}
				if i > 0 && j > 0 && a[i-1][j-1] == '0' {
					dsu.Union(i*m+j, (i-1)*m+j-1)
				}
				if i > 0 && j+1 < m && a[i-1][j+1] == '0' {
					dsu.Union(i*m+j, (i-1)*m+j+1)
				}
			}
		}
	}

	return dsu
}

func solve1(a []string) int {
	n := len(a)
	m := len(a[0])

	dsu := getComponents(a)

	next := make([][][2]int, n)
	for i := range n {
		next[i] = make([][2]int, m)
	}
	for i := range n {
		prev := -1
		for j := range m {
			if a[i][j] == '1' {
				next[i][j][1] = prev
				prev = j
			}
		}
		prev = m
		for j := m - 1; j >= 0; j-- {
			if a[i][j] == '1' {
				next[i][j][0] = prev
				prev = j
			}
		}
	}

	vis := make([][]bool, n)
	for i := range n {
		vis[i] = make([]bool, m)
	}

	checkSides := func(r int, c int) bool {
		var cnt int
		for i := range 4 {
			nr, nc := r+dd[i], c+dd[i+1]
			if nr >= 0 && nr < n && nc >= 0 && nc < m && vis[nr][nc] {
				cnt++
			}
		}
		return cnt <= 2
	}

	play := func(sr int, sc int, wc int) int {
		// 从sr, sc开始，沿着右手进行移动，看是否能够返回
		r, c := sr, sc
		// 目前是指向右边的 0/1/2/3 分别表示上/右/下/左
		d := 0
		// 这里tot计算不对
		var tot int
		var ans int
		// use todo to avoid a n check

		add := func(r int, c int, d int) {
			if d%2 == 0 {
				tot += abs(next[r][c][d/2]-c) - 1
			}
		}
		ok := true

		var todo [][]int
		for {
			todo = append(todo, []int{r, c})
			vis[r][c] = true
			ans++
			// 如果它的右边是1，必须移动到右边去
			nr, nc := r+dd[(d+1)%4], c+dd[(d+2)%4]
			if nr < 0 || nr == n || nc < 0 || nc == m {
				// 右边必须存在
				break
			}
			if a[nr][nc] == '1' {
				// 必须向右边移动
				d = (d + 1) % 4
				r, c = nr, nc
			} else {
				// 右边是0
				if dsu.Find(nr*m+nc) != wc {
					ok = false
					break
				}

				nr, nc = r+dd[d], c+dd[d+1]
				// 如果前方不存在
				if nr < 0 || nr == n || nc < 0 || nc == m {
					ok = false
					break
				}
				if a[nr][nc] == '1' {
					// 必须往前运动
					add(r, c, d)
					r, c = nr, nc
				} else {
					// 必须往左边运动
					nr, nc = r+dd[(d+3)%4], c+dd[(d+4)%4]

					if nr < 0 || nr == n || nc < 0 || nc == m || a[nr][nc] == '0' {
						ok = false
						break
					}
					add(r, c, d)
					add(r, c, (d+3)%4)
					d = (d + 3) % 4
					r, c = nr, nc
				}
			}

			if r == sr && c == sc {
				break
			}
		}

		if ok {
			for _, cur := range todo {
				if !checkSides(cur[0], cur[1]) {
					ok = false
					break
				}
			}
		}

		for _, cur := range todo {
			vis[cur[0]][cur[1]] = false
		}

		if !ok {
			return 0
		}

		tot /= 2

		if tot != dsu.sz[wc] {
			return 0
		}
		return ans
	}

	marked := make([]bool, n*m)
	var ans int
	for i := range n {
		for j := range m {
			if a[i][j] == '1' {
				// 有一个特殊形式要判定一下
				if i > 0 && j > 0 {
					if a[i-1][j-1] == '1' && a[i-1][j] == '1' && a[i][j-1] == '1' {
						ans = max(ans, 4)
					}
				}
				if j+1 < m && a[i][j+1] == '0' {
					wc := dsu.Find(i*m + j + 1)
					if !marked[wc] {
						ans = max(ans, play(i, j, wc))
						marked[wc] = true
					}
				}
			}
		}
	}

	return ans
}

func abs(num int) int {
	return max(num, -num)
}

type DSU struct {
	fa []int
	sz []int
}

func NewDSU(n int) *DSU {
	fa := make([]int, n)
	sz := make([]int, n)
	for i := range n {
		fa[i] = i
		sz[i] = 1
	}
	return &DSU{fa, sz}
}

func (d *DSU) Find(x int) int {
	if d.fa[x] != x {
		d.fa[x] = d.Find(d.fa[x])
	}
	return d.fa[x]
}

func (d *DSU) Union(x, y int) {
	x = d.Find(x)
	y = d.Find(y)
	if x == y {
		return
	}
	if d.sz[x] < d.sz[y] {
		x, y = y, x
	}
	d.sz[x] += d.sz[y]
	d.fa[y] = x
}
