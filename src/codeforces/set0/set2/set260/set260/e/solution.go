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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	_, _, ok, vertical, horizontal := drive(reader)
	if !ok {
		fmt.Fprintln(writer, "-1")
		return
	}
	fmt.Fprintf(writer, "%.1f %.1f\n", vertical[0], vertical[1])
	fmt.Fprintf(writer, "%.1f %.1f\n", horizontal[0], horizontal[1])
}

func drive(reader *bufio.Reader) (a []int, points [][]int, ok bool, vertical []float64, horizontal []float64) {
	var n int
	fmt.Fscan(reader, &n)
	points = make([][]int, n)
	for i := range n {
		points[i] = make([]int, 2)
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}
	a = make([]int, 9)
	for i := range 9 {
		fmt.Fscan(reader, &a[i])
	}
	ok, vertical, horizontal = solve(a, points)
	return
}

func solve(a []int, points [][]int) (ok bool, vertical []float64, horizontal []float64) {
	xs := slices.Clone(points)
	ys := slices.Clone(points)

	slices.SortFunc(xs, func(first []int, second []int) int {
		return cmp.Or(first[0]-second[0], first[1]-second[1])
	})

	slices.SortFunc(ys, func(first []int, second []int) int {
		return cmp.Or(first[1]-second[1], first[0]-second[0])
	})

	n := len(xs)
	rank := make(map[[2]int]int, n)
	for i, p := range ys {
		rank[[2]int{p[0], p[1]}] = i + 1
	}
	yr := make([]int, n)
	for i, p := range xs {
		yr[i] = rank[[2]int{p[0], p[1]}]
	}

	sumMask := func(flag int) int {
		var sum int
		for i := range 9 {
			if (flag>>i)&1 == 1 {
				sum += a[i]
			}
		}
		return sum
	}

	checkVertical := func(sum int) (bool, float64) {
		if xs[sum-1][0] == xs[sum][0] {
			return false, 0
		}
		return true, float64(xs[sum-1][0]+xs[sum][0]) / 2
	}

	checkHorizontal := func(sum int) (bool, float64) {
		if ys[sum-1][1] == ys[sum][1] {
			return false, 0
		}
		return true, float64(ys[sum-1][1]+ys[sum][1]) / 2
	}

	perms := permute(3, 0, nil)

	type candidate struct {
		k1, k2 int
		r1, r2 int
		e1, e2 int
		e3, e4 int
		x1, x2 float64
		y1, y2 float64
	}
	var cand []candidate

	var f1 func(flag int, s int)
	var f2 func(flag1 int, flag int, s int, x1 float64)
	var f3 func(flag1 int, flag2 int, x1 float64, x2 float64)

	f1 = func(flag int, s int) {
		if bits.OnesCount(uint(flag)) == 3 {
			sum1 := sumMask(flag)
			if ok, x1 := checkVertical(sum1); ok {
				f2(flag, 0, 0, x1)
			}
			return
		}
		for i := s; i < 9; i++ {
			if (flag>>i)&1 == 0 {
				f1(flag|(1<<i), i+1)
			}
		}
	}
	f3 = func(flag1 int, flag2 int, x1 float64, x2 float64) {
		var arr1, arr2, arr3 []int
		for i := range 9 {
			if (flag1>>i)&1 == 1 {
				arr1 = append(arr1, i)
			} else if (flag2>>i)&1 == 1 {
				arr2 = append(arr2, i)
			} else {
				arr3 = append(arr3, i)
			}
		}

		k1 := sumMask(flag1)
		k2 := sumMask(flag1 | flag2)

		for _, u := range perms {
			for _, v := range perms {
				for _, w := range perms {
					r1 := a[arr1[u[0]]] + a[arr2[v[0]]] + a[arr3[w[0]]]
					ok1, y1 := checkHorizontal(r1)
					if !ok1 {
						continue
					}
					r2 := r1 + a[arr1[u[1]]] + a[arr2[v[1]]] + a[arr3[w[1]]]
					ok2, y2 := checkHorizontal(r2)
					if !ok2 {
						continue
					}
					cand = append(cand, candidate{
						k1: k1, k2: k2,
						r1: r1, r2: r2,
						e1: a[arr1[u[0]]],
						e2: a[arr1[u[0]]] + a[arr1[u[1]]],
						e3: a[arr1[u[0]]] + a[arr2[v[0]]],
						e4: a[arr1[u[0]]] + a[arr1[u[1]]] + a[arr2[v[0]]] + a[arr2[v[1]]],
						x1: x1, x2: x2,
						y1: y1, y2: y2,
					})
				}
			}
		}
	}

	f2 = func(flag1 int, flag int, s int, x1 float64) {
		if bits.OnesCount(uint(flag)) == 3 {
			sum2 := sumMask(flag1 | flag)
			if ok, x2 := checkVertical(sum2); ok {
				f3(flag1, flag, x1, x2)
			}
			return
		}
		for i := s; i < 9; i++ {
			if flag1&(1<<i) == 0 && (flag>>i)&1 == 0 {
				f2(flag1, flag|(1<<i), i+1, x1)
			}
		}
	}

	f1(0, 0)

	type query struct {
		k, r int
		id   int
		pos  int
	}
	qs := make([]query, 0, len(cand)*4)
	for i, c := range cand {
		qs = append(qs,
			query{k: c.k1, r: c.r1, id: i, pos: 0},
			query{k: c.k1, r: c.r2, id: i, pos: 1},
			query{k: c.k2, r: c.r1, id: i, pos: 2},
			query{k: c.k2, r: c.r2, id: i, pos: 3},
		)
	}

	slices.SortFunc(qs, func(a, b query) int {
		return cmp.Or(a.k-b.k, a.r-b.r)
	})

	got := make([][4]int, len(cand))
	bit := newFenwick(n)
	ptr := 0
	for _, q := range qs {
		for ptr < q.k {
			bit.Add(yr[ptr], 1)
			ptr++
		}
		got[q.id][q.pos] = bit.Sum(q.r)
	}

	for i, c := range cand {
		if got[i][0] == c.e1 &&
			got[i][1] == c.e2 &&
			got[i][2] == c.e3 &&
			got[i][3] == c.e4 {
			return true, []float64{c.x1, c.x2}, []float64{c.y1, c.y2}
		}
	}

	return false, nil, nil
}

func permute(n int, flag int, buf []int) [][]int {
	if flag == (1<<n)-1 {
		return [][]int{slices.Clone(buf)}
	}
	var res [][]int
	for i := range n {
		if (flag>>i)&1 == 0 {
			buf = append(buf, i)
			sub := permute(n, flag|(1<<i), buf)
			res = append(res, sub...)
			buf = buf[:len(buf)-1]
		}
	}
	return res
}

type fenwick struct {
	arr []int
}

func newFenwick(n int) *fenwick {
	return &fenwick{arr: make([]int, n+1)}
}

func (f *fenwick) Add(pos int, val int) {
	for pos < len(f.arr) {
		f.arr[pos] += val
		pos += pos & -pos
	}
}

func (f *fenwick) Sum(pos int) int {
	var res int
	for pos > 0 {
		res += f.arr[pos]
		pos -= pos & -pos
	}
	return res
}

func solve1(a []int, points [][]int) (ok bool, vertical []float64, horizontal []float64) {

	xs := slices.Clone(points)
	ys := slices.Clone(points)

	slices.SortFunc(xs, func(first []int, second []int) int {
		return cmp.Or(first[0]-second[0], first[1]-second[1])
	})

	slices.SortFunc(ys, func(first []int, second []int) int {
		return cmp.Or(first[1]-second[1], first[0]-second[0])
	})

	n := len(xs)
	trs := make([]*node, n+1)

	trs[0] = new(node)

	offset := ys[0][1]
	my := ys[n-1][1]

	for i := range n {
		trs[i+1] = trs[i].add(xs[i][1]-offset, 0, my-offset)
	}

	// len(a) = 9
	// 从9个中选出3个放在最左边，3个放在中间
	// C(9, 3) * C(6, 3) = 9 * 8 * 7 / 6 * 6 * 5 * 4 / 6 = 1680

	var f1 func(flag int, s int) bool
	var f2 func(flag1 int, flag int, s int) bool
	var f3 func(flag1 int, flag2 int) bool

	vertical = make([]float64, 2)
	horizontal = make([]float64, 2)

	check1 := func(flag int) (bool, float64) {
		var sum int
		for i := range 9 {
			if (flag>>i)&1 == 1 {
				sum += a[i]
			}
		}
		if xs[sum-1][0] == xs[sum][0] {
			return false, 0
		}
		// 必须正好有sum个，在最左边
		return true, float64(xs[sum-1][0]+xs[sum][0]) / 2
	}

	f1 = func(flag int, s int) bool {
		// C(9, 3)
		if bits.OnesCount(uint(flag)) == 3 {
			if ok, x1 := check1(flag); ok && f2(flag, 0, 0) {
				vertical[0] = x1
				return true
			}
			return false
		}

		for i := s; i < 9; i++ {
			if (flag>>i)&1 == 0 {
				if f1(flag|(1<<i), i+1) {
					return true
				}
			}
		}

		return false
	}

	f2 = func(flag1 int, flag int, s int) bool {
		// C(6, 3)
		if bits.OnesCount(uint(flag)) == 3 {
			if ok, x2 := check1(flag1 | flag); ok && f3(flag1, flag) {
				vertical[1] = x2
				return true
			}
			return false
		}

		for i := s; i < 9; i++ {
			if flag1&(1<<i) == 0 && (flag>>i)&1 == 0 {
				if f2(flag1, (1<<i)|flag, i+1) {
					return true
				}
			}
		}

		return false
	}

	perms := permute(3, 0, nil)

	check2 := func(flag int) (bool, float64) {
		var sum int
		for i := range 9 {
			if (flag>>i)&1 == 1 {
				sum += a[i]
			}
		}
		if ys[sum-1][1] == ys[sum][1] {
			return false, 0
		}
		return true, float64(ys[sum-1][1]+ys[sum][1]) / 2
	}

	getBound := func(flag int) int {
		var sum int
		for i := range 9 {
			if (flag>>i)&1 == 1 {
				sum += a[i]
			}
		}
		return ys[sum-1][1]
	}

	check3 := func(arr1 []int, u []int, arr2 []int, v []int, y1 int, y2 int) bool {
		var sum int
		for _, i := range arr1 {
			sum += a[i]
		}
		// 这个是第一个平面的部分
		a0 := trs[sum].query(0, my-offset, 0, y1-offset)
		if a0 != a[arr1[u[0]]] {
			return false
		}
		a1 := trs[sum].query(0, my-offset, 0, y2-offset)
		if a1 != a[arr1[u[1]]]+a0 {
			return false
		}

		for _, i := range arr2 {
			sum += a[i]
		}
		// 两块区域的sum
		b0 := trs[sum].query(0, my-offset, 0, y1-offset)
		if b0 != a[arr2[v[0]]]+a0 {
			return false
		}
		b1 := trs[sum].query(0, my-offset, 0, y2-offset)
		if b1-b0 != a[arr2[v[1]]]+a[arr1[u[1]]] {
			return false
		}
		return true
	}

	f3 = func(flag1 int, flag2 int) bool {
		// 垂直方向已经处理好了，现在处理水平方向
		var arr1 []int
		var arr2 []int
		var arr3 []int
		for i := range 9 {
			if (flag1>>i)&1 == 1 {
				arr1 = append(arr1, i)
			} else if (flag2>>i)&1 == 1 {
				arr2 = append(arr2, i)
			} else {
				arr3 = append(arr3, i)
			}
		}

		// 6 * 6 * 6
		for _, u := range perms {
			for _, v := range perms {
				for _, w := range perms {
					vf1 := (1 << arr1[u[0]]) | (1 << arr2[v[0]]) | (1 << arr3[w[0]])
					if ok, y1 := check2(vf1); ok {
						vf2 := (1 << arr1[u[1]]) | (1 << arr2[v[1]]) | (1 << arr3[w[1]])
						if ok2, y2 := check2(vf2 | vf1); ok2 && check3(arr1, u, arr2, v, getBound(vf1), getBound(vf2|vf1)) {
							// 这里还必须保证，9个格子中的数字是相同的
							horizontal[0] = y1
							horizontal[1] = y2
							return true
						}
					}
				}
			}
		}

		return false
	}

	ok = f1(0, 0)

	return
}

type node struct {
	left, right *node
	cnt         int
}

func (n node) add(v int, l int, r int) *node {
	if l < r {
		mid := (l + r) / 2
		if v <= mid {
			if n.left == nil {
				n.left = new(node)
			}
			n.left = n.left.add(v, l, mid)
		} else {
			if n.right == nil {
				n.right = new(node)
			}
			n.right = n.right.add(v, mid+1, r)
		}
	}

	n.cnt++
	return &n
}

func (n *node) query(l int, r int, L int, R int) int {
	if n == nil || R < l || r < L {
		return 0
	}
	if L <= l && r <= R {
		return n.cnt
	}

	mid := (l + r) / 2

	return n.left.query(l, mid, L, R) + n.right.query(mid+1, r, L, R)
}
