package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
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

func drive(reader *bufio.Reader) int {
	n, _, k := readThreeNums(reader)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = readString(reader)
	}
	nums := readNNums(reader, 4)
	x1, y1, x2, y2 := nums[0], nums[1], nums[2], nums[3]
	return solve(a, x1, y1, x2, y2, k)
}

const inf = 1 << 60

func solve(a []string, x1 int, y1 int, x2 int, y2 int, k int) int {
	x1--
	y1--
	x2--
	y2--

	n := len(a)
	m := len(a[0])

	dist := make([][]int, n)

	rows := make([]*Set, n)
	cols := make([]*Set, m)
	for j := range m {
		cols[j] = NewSet(n)
	}

	for i := range n {
		dist[i] = make([]int, m)
		rows[i] = NewSet(m)
		for j := range m {
			dist[i][j] = -1
			rows[i].Set(j)
			cols[j].Set(i)
		}
	}

	que := make([]int, n*m)
	var head, tail int
	que[head] = x1*m + y1
	head++
	dist[x1][y1] = 0

	add := func(r, c int, v int) {
		dist[r][c] = v
		que[head] = r*m + c
		head++
		rows[r].Clear(c)
		cols[c].Clear(r)
	}

	rows[x1].Clear(y1)
	cols[y1].Clear(x1)

	for tail < head {
		r, c := que[tail]/m, que[tail]%m
		tail++
		// 先向上移动

		for {
			i := cols[c].QueryMax(0, r)
			if i < 0 || r-i > k || a[i][c] == '#' {
				break
			}
			add(i, c, dist[r][c]+1)
		}
		for {
			i := cols[c].QueryMin(r, n-1)
			if i == n || i-r > k || a[i][c] == '#' {
				break
			}
			add(i, c, dist[r][c]+1)
		}
		for {
			j := rows[r].QueryMax(0, c)
			if j < 0 || c-j > k || a[r][j] == '#' {
				break
			}
			add(r, j, dist[r][c]+1)
		}
		for {
			j := rows[r].QueryMin(c, m-1)
			if j == m || j-c > k || a[r][j] == '#' {
				break
			}
			add(r, j, dist[r][c]+1)
		}
	}

	return dist[x2][y2]
}

type Set struct {
	tr1 *SegTree
	tr2 *SegTree
}

func NewSet(n int) *Set {
	tr1 := NewSegTree(n, -1, func(a, b int) int {
		return max(a, b)
	})
	tr2 := NewSegTree(n, n, func(a, b int) int {
		return min(a, b)
	})
	return &Set{tr1, tr2}
}

func (set *Set) QueryMin(L int, R int) int {
	return set.tr2.Find(L, R+1)
}

func (set *Set) QueryMax(L int, R int) int {
	return set.tr1.Find(L, R+1)
}

func (set *Set) Clear(pos int) {
	set.tr1.Update(pos, -1)
	set.tr2.Update(pos, set.tr2.initValue)
}

func (set *Set) Set(pos int) {
	set.tr1.Update(pos, pos)
	set.tr2.Update(pos, pos)
}

type SegTree struct {
	arr       []int
	sz        int
	initValue int
	fn        func(int, int) int
}

func NewSegTree(n int, iv int, fn func(int, int) int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = iv
	}
	return &SegTree{arr, n, iv, fn}
}

func (tree *SegTree) Update(pos int, v int) {
	pos += tree.sz
	tree.arr[pos] = v
	for pos > 0 {
		tree.arr[pos>>1] = tree.fn(tree.arr[pos], tree.arr[pos^1])
		pos >>= 1
	}
}

func (tree *SegTree) Find(l, r int) int {
	l += tree.sz
	r += tree.sz

	ans := tree.initValue

	for l < r {
		if l&1 == 1 {
			ans = tree.fn(ans, tree.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			ans = tree.fn(ans, tree.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return ans
}
