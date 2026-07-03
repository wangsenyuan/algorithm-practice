package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	queries := make([][]int, k)
	for i := 0; i < k; i++ {
		s, _ := reader.ReadBytes('\n')
		var j int
		if s[0] == '1' {
			j = 4
		} else {
			j = 3
		}
		queries[i] = make([]int, j)
		var pos int
		for x := 0; x < j; x++ {
			pos = readInt(s, pos, &queries[i][x]) + 1
		}
	}

	res := solve(n, m, queries)
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
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

func solve(n, m int, queries [][]int) []int {
	type record struct {
		id  int
		col int
	}

	todo := make([][]record, n)

	tr := make(fenwick, m+2)

	k := len(queries)

	ans := make([]int, len(queries))

	for id := k - 1; id >= 0; id-- {
		cur := queries[id]
		if cur[0] == 1 {
			l, r, x := cur[1], cur[2], cur[3]
			l--
			r--
			tr.add(l, x)
			tr.add(r+1, -x)
		} else if cur[0] == 2 {
			i, x := cur[1], cur[2]
			i--
			for _, v := range todo[i] {
				ans[v.id] += tr.query(v.col) + x
			}
			todo[i] = todo[i][:0]
		} else {
			i, j := cur[1], cur[2]
			i--
			j--
			ans[id] = -tr.query(j)
			todo[i] = append(todo[i], record{id, j})
		}
	}

	for _, td := range todo {
		for _, v := range td {
			ans[v.id] += tr.query(v.col)
		}
	}

	var res []int
	for i := 0; i < k; i++ {
		if queries[i][0] == 3 {
			res = append(res, ans[i])
		}
	}
	return res
}

type fenwick []int

func (tr fenwick) add(p int, v int) {
	p++
	for p < len(tr) {
		tr[p] += v
		p += p & -p
	}
}

func (tr fenwick) query(p int) int {
	p++
	var res int
	for p > 0 {
		res += tr[p]
		p -= p & -p
	}
	return res
}

func solve1(n, m int, queries [][]int) []int {
	k := len(queries)

	type pair struct {
		first  int
		second int
	}

	row := make([]pair, n)

	nodes := make([]*node, k+1)
	nodes[0] = new(node)
	var it int

	var ans []int

	for _, cur := range queries {
		if cur[0] == 1 {
			l, r, x := cur[1], cur[2], cur[3]
			nodes[it+1] = nodes[it].Add(l-1, r-1, 0, m-1, x)
			it++
		} else if cur[0] == 2 {
			i := cur[1] - 1
			x := cur[2]
			row[i] = pair{it, x}
		} else {
			i, j := cur[1], cur[2]
			tmp := nodes[it].Query(j-1, 0, m-1, 0)
			tmp -= nodes[row[i-1].first].Query(j-1, 0, m-1, 0)
			tmp += row[i-1].second
			ans = append(ans, tmp)
		}
	}

	return ans
}

type node struct {
	left, right *node
	v           int
}

func (n *node) getValue() int {
	if n == nil {
		return 0
	}
	return n.v
}

func (n *node) getLeft() *node {
	if n == nil {
		return nil
	}
	return n.left
}

func (n *node) getRight() *node {
	if n == nil {
		return nil
	}
	return n.right
}

func (n *node) Add(L int, R int, l int, r int, v int) *node {
	// 要在区间L...R增加v
	tmp := new(node)
	tmp.left = n.getLeft()
	tmp.right = n.getRight()
	tmp.v = n.getValue()
	if L == l && r == R {
		tmp.v += v
		return tmp
	}
	mid := (l + r) / 2
	if L <= mid {
		tmp.left = tmp.left.Add(max(L, l), min(mid, R), l, mid, v)
	}
	if mid < R {
		tmp.right = tmp.right.Add(max(mid+1, L), min(r, R), mid+1, r, v)
	}
	return tmp
}

func (n *node) Query(p int, l int, r int, val int) int {
	if n == nil {
		return val
	}
	val += n.v
	if l == r {
		return val
	}
	mid := (l + r) / 2
	if p <= mid {
		return n.left.Query(p, l, mid, val)
	}
	return n.right.Query(p, mid+1, r, val)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
