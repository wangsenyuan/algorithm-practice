package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, ans := range res {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 3)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	return solve(n, edges, a, queries)
}

const mod = 1_000_000_007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

const X = 1e7 + 10

var primes []int
var lpf [X]int
var pos [X]int

func init() {
	for i := 2; i < X; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, p := range primes {
			if p*i >= X {
				break
			}
			lpf[p*i] = p
			if i%p == 0 {
				break
			}
		}
	}
	for i, p := range primes {
		pos[p] = i
	}
}

type event struct {
	id   int
	x    int
	kind int
}

func solve(n int, edges [][]int, a []int, queries [][]int) []int {

	adj := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	m := len(primes)

	dep := make([]int, n+1)

	H := bits.Len(uint(n + 1))
	fa := make([][]int, n+1)
	fa[0] = make([]int, H)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		fa[u] = make([]int, H)
		fa[u][0] = p
		for i := 1; i < H; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}

		for _, v := range adj[u] {
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(0, 1)

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := H - 1; i >= 0; i-- {
			if dep[u]-(1<<i) >= dep[v] {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := H - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	ans := make([]int, len(queries))
	record := make([]map[int][]int, len(queries))
	factors := make([][]pair, len(queries))
	events := make([][]event, n+1)
	for i := range ans {
		u, v, x := queries[i][0], queries[i][1], queries[i][2]
		p := lca(u, v)
		events[p] = append(events[p], event{id: i, x: x, kind: 0})
		events[u] = append(events[u], event{id: i, x: x, kind: 1})
		events[v] = append(events[v], event{id: i, x: x, kind: 1})

		factors[i] = getFactors(x)
		ans[i] = gcd(a[p-1], x)
	}

	freq := make([][]int, m)
	for i := range m {
		freq[i] = make([]int, 25)
	}

	var dfs3 func(p int, u int)
	dfs3 = func(p int, u int) {

		arr := getFactors(a[u-1])
		for _, f := range arr {
			w := f.first
			c := f.second
			freq[pos[w]][c]++
		}

		for _, e := range events[u] {
			if e.kind == 0 {
				id := e.id
				record[id] = make(map[int][]int)
				for _, f := range factors[id] {
					w := f.first
					record[id][pos[w]] = slices.Clone(freq[pos[w]])
				}
			}
		}

		for _, v := range adj[u] {
			if p != v {
				dfs3(u, v)
			}
		}

		for _, e := range events[u] {
			if e.kind == 1 {
				id := e.id
				tmp := record[id]
				for _, f := range factors[id] {
					w := f.first
					c := f.second
					before := tmp[pos[w]]
					cur := freq[pos[w]]
					var sum int
					for i := range 25 {
						if i <= c {
							sum += (cur[i] - before[i]) * i
						} else {
							sum += (cur[i] - before[i]) * c
						}
					}
					ans[id] = mul(ans[id], pow(w, sum))
				}
			}
		}

		for _, f := range arr {
			w := f.first
			c := f.second
			freq[pos[w]][c]--
		}
	}

	dfs3(0, 1)

	return ans
}

type pair struct {
	first  int
	second int
}

func getFactors(x int) []pair {
	var arr []pair
	for num := x; num > 1; {
		w := lpf[num]
		var c int
		for num%w == 0 {
			c++
			num /= w
		}
		arr = append(arr, pair{first: w, second: c})
	}
	slices.SortFunc(arr, func(x pair, y pair) int {
		return cmp.Or(x.first-y.first, x.second-y.second)
	})
	return arr
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
