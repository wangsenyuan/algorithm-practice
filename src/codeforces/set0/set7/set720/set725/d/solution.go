package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
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
	n := readNum(reader)
	teams := make([][]int, n)
	for i := 0; i < n; i++ {
		teams[i] = readNNums(reader, 2)
	}
	return solve(teams)
}

type team struct {
	id int
	t  int
	w  int
}

func solve(teams [][]int) int {
	n := len(teams)

	arr := make([]team, n-1)

	for i := 1; i < n; i++ {
		arr[i-1] = team{i, teams[i][0], teams[i][1]}
	}

	// 需要的越小越好
	slices.SortFunc(arr, func(a, b team) int {
		x := a.w + 1 - a.t
		y := b.w + 1 - b.t
		return cmp.Compare(x, y)
	})

	first := teams[0]

	t1 := first[0]
	pos := make([]int, n)
	pos[0] = -1
	best := 1
	for i := 0; i < n-1; i++ {
		pos[arr[i].id] = i
		if arr[i].t > t1 {
			best++
		}
	}

	slices.SortFunc(arr, func(a, b team) int {
		return cmp.Compare(b.t, a.t)
	})

	// 是溢出了么？
	tr := NewTree(n - 1)

	for i, cur := range arr {
		ti, wi := cur.t, cur.w

		tr.Update(pos[cur.id], wi+1-ti)

		// 修改到x, 全部给出去似乎也ok
		var x int
		if i+1 < len(arr) {
			x = arr[i+1].t
		}

		// 必须保证i后面的部分，都要比 ti-1 小
		if ti > x && x <= t1 {
			y := t1 - x
			cnt := tr.Count(y)
			tmp := i + 1 - cnt
			best = min(best, tmp+1)
		}
	}

	return best
}

const inf = 1 << 61

type Tree struct {
	val []int
	cnt []int
}

func NewTree(n int) *Tree {
	val := make([]int, 4*n)
	cnt := make([]int, 4*n)
	return &Tree{val: val, cnt: cnt}
}

func (t *Tree) Update(p int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			t.val[i] = v
			t.cnt[i] = 1
			return
		}
		mid := (l + r) >> 1
		if p <= mid {
			loop(2*i+1, l, mid)
		} else {
			loop(2*i+2, mid+1, r)
		}
		t.val[i] = min(inf, t.val[2*i+1]+t.val[2*i+2])
		t.cnt[i] = t.cnt[2*i+1] + t.cnt[2*i+2]
	}
	loop(0, 0, len(t.val)/4-1)
}

func (t *Tree) Count(w int) int {
	var loop func(i int, l int, r int, w int) int
	loop = func(i int, l int, r int, w int) int {
		if l == r {
			if w >= t.val[i] {
				return t.cnt[i]
			}
			return 0
		}
		mid := (l + r) >> 1
		if w <= t.val[2*i+1] {
			return loop(2*i+1, l, mid, w)
		}
		// w > t.val[2 * i + 2]
		return t.cnt[2*i+1] + loop(2*i+2, mid+1, r, w-t.val[2*i+1])
	}
	return loop(0, 0, len(t.val)/4-1, w)
}
