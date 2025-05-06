package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	targets := make([][]int, n)
	for i := range n {
		targets[i] = readNNums(reader, 5)
	}
	m := readNum(reader)
	shots := make([][]int, m)
	for i := range m {
		shots[i] = readNNums(reader, 2)
	}
	return solve(targets, shots)
}

type target struct {
	x1 int
	x2 int
	y1 int
	y2 int
	z  int
	id int
}

func nthElement[T any](arr []T, l int, mid int, r int, cmp func(i int, j int) bool) {
	// Use quickselect algorithm
	quickSelect(arr, l, r, mid, cmp)
}

func quickSelect[T any](arr []T, left, right, n int, cmp func(i int, j int) bool) {
	if left == right {
		return
	}

	// Get the pivot position
	pivotIdx := partition(arr, left, right, cmp)

	if n == pivotIdx {
		return
	}
	if n < pivotIdx {
		quickSelect(arr, left, pivotIdx-1, n, cmp)
	} else {
		quickSelect(arr, pivotIdx+1, right, n, cmp)
	}
}

func partition[T any](arr []T, left, right int, cmp func(i int, j int) bool) int {
	// Choose rightmost element as pivot
	i := left - 1

	for j := left; j < right; j++ {
		if cmp(j, right) {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func solve(targets [][]int, shots [][]int) []int {
	n := len(targets)
	m := len(shots)

	type point struct {
		x, y int
		id   int
	}

	points := make([]point, m)
	for i, cur := range shots {
		points[i] = point{cur[0], cur[1], i}
	}

	cmp0 := func(i int, j int) bool {
		return points[i].x < points[j].x
	}
	cmp1 := func(i int, j int) bool {
		return points[i].y < points[j].y
	}

	cmp := []func(i int, j int) bool{cmp0, cmp1}

	type node struct {
		left, right    *node
		x1, y1, x2, y2 int
		val            int
	}

	pull := func(n *node) {
		n.val = min(n.left.val, n.right.val)
		n.x1 = min(n.left.x1, n.right.x1)
		n.y1 = min(n.left.y1, n.right.y1)
		n.x2 = max(n.left.x2, n.right.x2)
		n.y2 = max(n.left.y2, n.right.y2)
	}

	var build func(l int, r int, d int) *node

	build = func(l int, r int, d int) *node {
		if l == r {
			pt := points[l]
			return &node{x1: pt.x, y1: pt.y, x2: pt.x, y2: pt.y, val: pt.id}
		}
		mid := (l + r) / 2
		nthElement(points, l, mid, r, cmp[d])
		res := new(node)
		res.left = build(l, mid, d^1)
		res.right = build(mid+1, r, d^1)
		pull(res)
		return res
	}

	tr := build(0, m-1, 0)

	var query func(n *node, l int, r int, x1 int, x2 int, y1 int, y2 int, res *int)
	query = func(n *node, l int, r int, x1 int, x2 int, y1 int, y2 int, res *int) {
		if n.x2 < x1 || x2 < n.x1 || n.y2 < y1 || y2 < n.y1 {
			return
		}
		if n.val >= *res {
			return
		}
		if x1 <= n.x1 && n.x2 <= x2 && y1 <= n.y1 && n.y2 <= y2 {
			*res = n.val
			return
		}
		mid := (l + r) / 2
		if n.left.val < n.right.val {
			// left first
			query(n.left, l, mid, x1, x2, y1, y2, res)
			query(n.right, mid+1, r, x1, x2, y1, y2, res)
		} else {
			query(n.right, mid+1, r, x1, x2, y1, y2, res)
			query(n.left, l, mid, x1, x2, y1, y2, res)
		}
	}

	var update func(n *node, l int, r int, pos int)
	update = func(n *node, l int, r int, pos int) {
		if l == r {
			n.val = inf
			return
		}
		mid := (l + r) / 2
		if pos <= mid {
			update(n.left, l, mid, pos)
		} else {
			update(n.right, mid+1, r, pos)
		}
		pull(n)
	}

	pos := make([]int, m)
	for i := range points {
		pos[points[i].id] = i
	}

	ans := make([]int, m)

	ts := make([]target, n)
	for i := range n {
		ts[i] = target{
			x1: targets[i][0],
			x2: targets[i][1],
			y1: targets[i][2],
			y2: targets[i][3],
			z:  targets[i][4],
			id: i + 1,
		}
	}

	slices.SortFunc(ts, func(a, b target) int {
		return a.z - b.z
	})

	for _, t := range ts {
		i := inf
		query(tr, 0, m-1, t.x1, t.x2, t.y1, t.y2, &i)
		if i == inf {
			continue
		}
		ans[i] = t.id
		update(tr, 0, m-1, pos[i])
	}

	return ans
}

func solve1(targets [][]int, shots [][]int) []int {
	n := len(targets)
	ts := make([]target, n)
	box := []int{inf, -inf, inf, -inf}

	updateBox := func(x int, y int) {
		box[0] = min(box[0], min(x, y))
		box[1] = max(box[1], max(x, y))
		box[2] = min(box[2], min(x, y))
		box[3] = max(box[3], max(x, y))
	}
	for i := range n {
		ts[i] = target{
			x1: targets[i][0],
			x2: targets[i][1] + 1,
			y1: targets[i][2],
			y2: targets[i][3] + 1,
			z:  targets[i][4],
			id: i + 1,
		}
		updateBox(ts[i].x1, ts[i].y1)
		updateBox(ts[i].x2, ts[i].y2)
	}

	for _, s := range shots {
		updateBox(s[0], s[1])
	}

	box[1]++
	box[3]++
	var tr *Tree

	for i, s := range shots {
		tr = tr.Add(box[0], box[1], box[2], box[3], s[0], s[1], i)
	}

	slices.SortFunc(ts, func(a, b target) int {
		return a.z - b.z
	})

	m := len(shots)
	ans := make([]int, m)

	for _, t := range ts {
		i := tr.Query(t.x1, t.x2, t.y1, t.y2)
		if i == inf {
			continue
		}
		ans[i] = t.id
		s := shots[i]
		tr = tr.Remove(s[0], s[1])
	}

	return ans
}

const inf = 1 << 60

type Tree struct {
	children []*Tree
	val      int
	arr      []int
	box      []int
}

func (t *Tree) GetValue() int {
	if t == nil {
		return inf
	}
	return t.val
}

func (t *Tree) pull() *Tree {
	t.val = min(t.children[0].GetValue(), t.children[1].GetValue(), t.children[2].GetValue(), t.children[3].GetValue())
	if t.val == inf {
		return nil
	}
	return t
}

func (t *Tree) Add(x1, x2, y1, y2 int, x, y int, v int) *Tree {
	if t == nil {
		t = new(Tree)
		t.box = []int{x1, x2, y1, y2}
		t.val = inf
		t.children = make([]*Tree, 4)
	}
	if x1+1 == x2 && y1+1 == y2 {
		// leaf node
		t.val = min(t.val, v)
		t.arr = append(t.arr, v)
		return t
	}
	mx := (x1 + x2) / 2
	my := (y1 + y2) / 2
	if x < mx && y < my {
		t.children[0] = t.children[0].Add(x1, mx, y1, my, x, y, v)
	} else if x < mx && y >= my {
		t.children[1] = t.children[1].Add(x1, mx, my, y2, x, y, v)
	} else if x >= mx && y >= my {
		t.children[2] = t.children[2].Add(mx, x2, my, y2, x, y, v)
	} else {
		t.children[3] = t.children[3].Add(mx, x2, y1, my, x, y, v)
	}
	return t.pull()
}

func (t *Tree) Remove(x, y int) *Tree {
	if t == nil {
		return nil
	}
	if t.box[0]+1 == t.box[1] && t.box[2]+1 == t.box[3] {
		t.arr = t.arr[1:]
		if len(t.arr) == 0 {
			return nil
		}
		t.val = t.arr[0]
		return t
	}
	mx := (t.box[0] + t.box[1]) / 2
	my := (t.box[2] + t.box[3]) / 2
	if x < mx && y < my {
		t.children[0] = t.children[0].Remove(x, y)
	} else if x < mx && y >= my {
		t.children[1] = t.children[1].Remove(x, y)
	} else if x >= mx && y >= my {
		t.children[2] = t.children[2].Remove(x, y)
	} else {
		t.children[3] = t.children[3].Remove(x, y)
	}
	return t.pull()
}

func (t *Tree) Query(x1, x2, y1, y2 int) int {
	if t == nil || x2 <= t.box[0] || t.box[1] <= x1 || y2 <= t.box[2] || t.box[3] <= y1 {
		return inf
	}

	if x1 <= t.box[0] && t.box[1] <= x2 && y1 <= t.box[2] && t.box[3] <= y2 {
		return t.val
	}

	res := t.children[0].Query(x1, x2, y1, y2)
	res = min(res, t.children[1].Query(x1, x2, y1, y2))
	res = min(res, t.children[2].Query(x1, x2, y1, y2))
	res = min(res, t.children[3].Query(x1, x2, y1, y2))
	return res
}
