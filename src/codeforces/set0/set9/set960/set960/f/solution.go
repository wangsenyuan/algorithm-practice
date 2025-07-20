package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 3)
	}
	return solve(n, edges)
}

type pair struct {
	first  int
	second int
}

func solve(n int, edges [][]int) int {
	cnt := make([]int, n)
	arr := make([][]pair, n)
	for i, cur := range edges {
		v, w := cur[1]-1, cur[2]
		arr[v] = append(arr[v], pair{w, i})
		cnt[v]++
	}

	m := len(edges)
	pos := make([]int, m)
	for v := range n {
		// 在weight相同的情况下，把后面的放在前面，这样在计算的时候，可以不考虑weight相同的情况，
		slices.SortFunc(arr[v], func(a, b pair) int {
			return cmp.Or(a.first-b.first, b.second-a.second)
		})
		for j, cur := range arr[v] {
			pos[cur.second] = j
		}
	}

	trs := make([]SegTree, n)
	for i := range n {
		trs[i] = make(SegTree, 2*cnt[i])
	}

	var ans int
	for i, cur := range edges {
		u, v, w := cur[0]-1, cur[1]-1, cur[2]
		// 要对u序列进行处理
		j := sort.Search(len(arr[u]), func(j int) bool {
			return arr[u][j].first >= w
		})

		best := trs[u].Get(0, j)
		best++
		ans = max(ans, best)
		tmp := trs[v].Get(pos[i], pos[i]+1)
		if best > tmp {
			trs[v].Update(pos[i], best)
		}
	}

	return ans
}

func last(arr []pair) pair {
	return arr[len(arr)-1]
}

type SegTree []int

func (tr SegTree) Update(i int, v int) {
	n := len(tr) / 2
	i += n
	tr[i] = v
	for i > 1 {
		tr[i>>1] = max(tr[i], tr[i^1])
		i >>= 1
	}
}

func (tr SegTree) Get(l int, r int) int {
	n := len(tr) / 2
	l += n
	r += n
	var res int
	for l < r {
		if l&1 == 1 {
			res = max(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
