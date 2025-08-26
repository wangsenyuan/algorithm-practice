package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	drive(reader, writer)
}

func drive(reader *bufio.Reader, writer *bufio.Writer) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	L := make([]int, n-1)
	for i := range L {
		fmt.Fscan(reader, &L[i])
	}

	query := func() []int {
		var u, d int
		fmt.Fscan(reader, &u, &d)
		return []int{u, d}
	}

	solve(n, L, m, query, writer)
}

type data struct {
	dist int32
	cnt  int32
}

func merge(a, b []data) []data {
	var res []data
	for i, j := 0, 0; i < len(a) || j < len(b); {
		if i < len(a) && j < len(b) && a[i].dist == b[j].dist {
			res = append(res, data{dist: a[i].dist, cnt: a[i].cnt + b[j].cnt})
			i++
			j++
		} else if j == len(b) || (i < len(a) && a[i].dist < b[j].dist) {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}

	return res
}

func shift(arr []data, d int) []data {
	if len(arr) == 0 {
		return nil
	}
	res := make([]data, len(arr))
	copy(res, arr)
	for i := range res {
		res[i].dist += int32(d)
	}
	return res
}

func solve(n int, D []int, q int, query func() []int, writer *bufio.Writer) {
	dp := make([][]data, n+1)
	fp := make([][]int, n+1)

	deg := make([]byte, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = append(dp[i], data{dist: 0, cnt: 1})
		deg[i/2]++
	}
	que := make([]int32, n)
	var head, tail int
	for i := 1; i <= n; i++ {
		if deg[i] == 0 {
			que[head] = int32(i)
			head++
		}
	}

	for tail < head {
		v := que[tail]
		tail++
		if v == 1 {
			continue
		}
		u := v / 2
		dp[u] = merge(dp[u], shift(dp[v], D[v-2]))
		deg[u]--
		if deg[u] == 0 {
			que[head] = int32(u)
			head++
		}
	}

	for u := 1; u <= n; u++ {
		// 把second变成sum
		fp[u] = make([]int, len(dp[u]))
		fp[u][0] = int(dp[u][0].cnt) * int(dp[u][0].dist)
		for i := 1; i < len(dp[u]); i++ {
			fp[u][i] = fp[u][i-1] + int(dp[u][i].cnt)*int(dp[u][i].dist)
			dp[u][i].cnt += dp[u][i-1].cnt
		}
	}

	for range q {
		q := query()
		var v int
		var res int
		u, d := q[0], q[1]
		for d > 0 {
			j := sort.Search(len(dp[u]), func(i int) bool {
				return dp[u][i].dist > int32(d)
			})

			if j > 0 {
				res += d*int(dp[u][j-1].cnt) - fp[u][j-1]
			}

			if v > 0 {
				j := sort.Search(len(dp[v]), func(i int) bool {
					return dp[v][i].dist > int32(d-D[v-2])
				})
				if j > 0 {
					res -= (d-D[v-2])*int(dp[v][j-1].cnt) - fp[v][j-1]
				}
			}

			if u == 1 {
				break
			}
			d -= D[u-2]
			v = u
			u /= 2
		}
		fmt.Fprintln(writer, res)
	}
}
