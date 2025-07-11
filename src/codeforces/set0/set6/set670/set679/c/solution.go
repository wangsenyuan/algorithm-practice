package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		grid[i] = readString(reader)
	}
	return solve1(n, k, grid)
}

func solve(n int, k int, a []string) int {
	dir4 := []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	id := make([][]int, n)
	for i := range id {
		id[i] = make([]int, n)
	}
	size := []int{0}
	var dfs func(int, int)
	dfs = func(i, j int) {
		id[i][j] = len(size) - 1
		size[len(size)-1]++
		for _, d := range dir4 {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < n && 0 <= y && y < n && a[x][y] != 'X' && id[x][y] == 0 {
				dfs(x, y)
			}
		}
	}
	for i, row := range a {
		for j, v := range row {
			if v == 'X' || id[i][j] > 0 {
				continue
			}
			size = append(size, 0)
			dfs(i, j)
		}
	}

	var ts, ans int

	vis := make([]int, len(size))
	for top := range n - k + 1 {
		for r := range n {
			for _, row := range id[top : top+k] {
				size[row[r]]--
			}

			l := r + 1 - k
			if l < 0 {
				continue
			}

			s := 0
			ts++
			for i := max(top-1, 0); i <= min(top+k, n-1); i++ {
				j, step, c := l, 1, k
				if i != top-1 && i != top+k {
					j, step, c = l-1, k+1, 2
				}
				for range c {
					if 0 <= j && j < n {
						p := id[i][j]
						if p > 0 && vis[p] != ts {
							vis[p] = ts
							s += size[p]
						}
					}
					j += step
				}
			}
			ans = max(ans, s)

			for _, row := range id[top : top+k] {
				size[row[l]]++
			}
		}

		for l := n - k + 1; l < n; l++ {
			for _, row := range id[top : top+k] {
				size[row[l]]++
			}
		}
	}
	return ans + k*k
}

func solve1(n int, k int, grid []string) int {
	if k == n {
		return n * n
	}
	set := NewDSU(n * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '.' {
				if i > 0 && grid[i-1][j] == '.' {
					set.Union((i-1)*n+j, i*n+j)
				}
				if j > 0 && grid[i][j-1] == '.' {
					set.Union(i*n+j-1, i*n+j)
				}
			}
		}
	}

	var ans int
	cnt := make([]int, n*n)
	id := make([][]int, n*n)

	for i := range n {
		id[i] = make([]int, n)
		for j := range n {
			id[i][j] = set.Find(i*n + j)
			if grid[i][j] == '.' && set.Find(i*n+j) == i*n+j {
				cnt[i*n+j] = set.cnt[i*n+j]
				ans = max(ans, cnt[i*n+j])
			}
		}
	}

	freq := make([]int, n*n)

	var sum int

	add := func(i int, j int) {
		p := set.Find(i*n + j)
		freq[p]++

		if freq[p] == 1 {
			sum += cnt[p]
		}
	}

	rem := func(i int, j int) {
		p := set.Find(i*n + j)
		freq[p]--
		if freq[p] == 0 {
			sum -= cnt[p]
		}
	}

	for i := k - 1; i < n; i++ {
		// 把头部的区域给覆盖掉
		for u := 0; u < k; u++ {
			for v := 0; v < k; v++ {
				if grid[i-u][v] == '.' {
					cnt[id[i-u][v]]--
				}
			}
		}
		for j := k - 1; j < n; j++ {
			if j >= k {
				for u := 0; u < k; u++ {
					// 第一列的恢复
					if grid[i-u][j-k] == '.' {
						cnt[id[i-u][j-k]]++
					}
					// 新的一列去除掉
					if grid[i-u][j] == '.' {
						cnt[id[i-u][j]]--
					}
				}
			}
			// 把四周连接起来
			for u := 0; u < k; u++ {
				if i >= k {
					add(i-k, j-u)
				}
				if i+1 < n {
					add(i+1, j-u)
				}
				if j >= k {
					add(i-u, j-k)
				}
				if j+1 < n {
					add(i-u, j+1)
				}
			}
			ans = max(ans, k*k+sum)
			// restore
			for u := 0; u < k; u++ {
				if i >= k {
					rem(i-k, j-u)
				}
				if i+1 < n {
					rem(i+1, j-u)
				}
				if j >= k {
					rem(i-u, j-k)
				}
				if j+1 < n {
					rem(i-u, j+1)
				}
			}
		}
		for u := 0; u < k; u++ {
			for v := 0; v < k; v++ {
				if grid[i-u][n-1-v] == '.' {
					cnt[id[i-u][n-1-v]]++
				}
			}
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
