package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res, _ := process(reader)
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for _, x := range res {
				buf.WriteString(fmt.Sprintf("%d ", x))
			}
			buf.WriteByte('\n')
		}
	}
	fmt.Println(buf.String())
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

func process(reader *bufio.Reader) (res []int, songs [][]int) {
	n := readNum(reader)
	songs = make([][]int, n)
	for i := range n {
		songs[i] = readNNums(reader, 2)
	}
	res = solve(songs)
	return
}

func solve(songes [][]int) []int {
	var xs []int
	var ys []int
	for _, cur := range songes {
		xs = append(xs, cur[0])
		ys = append(ys, cur[1])
	}
	xs = sortAndUnique(xs)
	ys = sortAndUnique(ys)

	n := len(xs) + len(ys)
	m := len(songes)
	g := make([]map[int]struct{}, n)
	for i := range n {
		g[i] = make(map[int]struct{})
	}

	type edge struct {
		u  int
		v  int
		id int
	}

	pos := make([]edge, m)

	for i, cur := range songes {
		u := sort.SearchInts(xs, cur[0])
		v := sort.SearchInts(ys, cur[1]) + len(xs)
		g[u][v] = struct{}{}
		g[v][u] = struct{}{}
		pos[i] = edge{u, v, i}
	}

	var odd int
	var start int
	for i := range n {
		if len(g[i])%2 == 1 {
			odd++
			start = i
		}
	}
	if odd > 2 {
		return nil
	}

	var arr []int
	var dfs func(u int)
	dfs = func(u int) {
		for v := range g[u] {
			delete(g[u], v)
			delete(g[v], u)
			dfs(v)
		}
		arr = append(arr, u)
	}

	dfs(start)

	if len(arr) != m+1 {
		return nil
	}

	slices.SortFunc(pos, func(a, b edge) int {
		return cmp.Or(a.u-b.u, a.v-b.v)
	})

	marked := make([]bool, m)

	res := make([]int, m)

	for i := range m {
		u := arr[i]
		v := arr[i+1]
		if u >= len(xs) {
			// v < len(xs)
			u, v = v, u
		}
		j := sort.Search(m, func(j int) bool {
			if pos[j].u > u || pos[j].u == u && pos[j].v >= v {
				return true
			}
			return false
		})

		if j == m || marked[j] || pos[j].u != u || pos[j].v != v {
			return nil
		}
		marked[j] = true
		res[i] = pos[j].id + 1
	}
	return res
}

func sortAndUnique(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	sort.Ints(res)
	var n int
	for i := 1; i <= len(res); i++ {
		if i == len(res) || res[i] != res[i-1] {
			res[n] = res[i-1]
			n++
		}
	}
	return res[:n]
}
