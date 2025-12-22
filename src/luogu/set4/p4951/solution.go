package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, h := readThreeNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 4)
	}
	res := solve(n, edges, h)

	fmt.Printf("%.4f\n", res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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
func solve(n int, edges [][]int, h int) float64 {
	// let h - sum(c + x * t) * s >= 0
	// edge = (u, v, c, t)
	uf := CreateUFSet(n)
	check := func(x float64) bool {
		sort.Slice(edges, func(i, j int) bool {
			return float64(edges[i][2]-edges[j][2]) <= x*float64(edges[j][3]-edges[i][3])
		})
		var a, b int
		for _, edge := range edges {
			u, v := edge[0]-1, edge[1]-1
			if uf.Union(u, v) {
				a += edge[2]
				b += edge[3]
			}
		}

		uf.Reset()

		return float64(h-a) >= float64(b)*x
	}

	var l float64
	var r float64 = 1 << 30
	for i := 0; i < 60; i++ {
		mid := (l + r) / 2
		if check(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}

type UFSet struct {
	set  []int
	cnt  []int
	size int
}

func CreateUFSet(n int) UFSet {
	set := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = 1
		set[i] = i
	}
	return UFSet{set, cnt, n}
}

func (uf *UFSet) Reset() {
	for i := 0; i < len(uf.set); i++ {
		uf.set[i] = i
		uf.cnt[i] = 1
	}
	uf.size = len(uf.set)
}

func (uf *UFSet) Union(a, b int) bool {
	var find func(u int) int
	find = func(u int) int {
		if uf.set[u] != u {
			uf.set[u] = find(uf.set[u])
		}
		return uf.set[u]
	}

	pa := find(a)
	pb := find(b)
	if pa == pb {
		return false
	}

	if uf.cnt[pa] < uf.cnt[pb] {
		pa, pb = pb, pa
	}

	uf.set[pb] = pa
	uf.cnt[pa] += uf.cnt[pb]
	uf.size--

	return true
}
