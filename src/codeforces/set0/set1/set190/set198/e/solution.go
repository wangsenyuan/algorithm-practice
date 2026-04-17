package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := drive(reader)

	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var x, y, p, r, n int
	fmt.Fscan(reader, &x, &y, &p, &r, &n)
	gripers := make([][]int, n)
	for i := range n {
		gripers[i] = make([]int, 5)
		fmt.Fscan(reader, &gripers[i][0], &gripers[i][1], &gripers[i][2], &gripers[i][3], &gripers[i][4])
	}
	return solve(x, y, p, r, gripers)
}

type gripper struct {
	id int
	x  int
	y  int
	m  int
	p  int
	r  int
}

func (g gripper) distFrom(x int, y int) int {
	dx := g.x - x
	dy := g.y - y
	return dx*dx + dy*dy
}

func solve(x int, y int, p int, r int, gripers [][]int) int {
	n := len(gripers)
	arr := make([]gripper, n)
	for i, cur := range gripers {
		arr[i] = gripper{i, cur[0], cur[1], cur[2], cur[3], cur[4]}
	}

	slices.SortFunc(arr, func(a, b gripper) int {
		return a.m - b.m
	})

	dists := make([]int, n)
	for i := range n {
		dists[i] = arr[i].distFrom(x, y)
	}

	slices.Sort(dists)
	dists = slices.Compact(dists)

	tr := NewTree(len(dists))
	for _, cur := range arr {
		d := cur.distFrom(x, y)
		i := sort.SearchInts(dists, d)
		tr.Add(i, cur)
	}

	var que [][]int
	que = append(que, []int{p, r})

	var ans int
	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		w, r := cur[0], cur[1]
		i := sort.SearchInts(dists, r*r)
		if i == len(dists) || dists[i] > r*r {
			i--
		}
		tr.Play(i, w, func(g gripper) {
			ans++
			que = append(que, []int{g.p, g.r})
		})
	}

	return ans
}

const inf = 1 << 60

type Tree struct {
	arr     [][]gripper
	minMass []int
	sz      int
}

func NewTree(n int) *Tree {
	arr := make([][]gripper, 4*n)
	minMass := make([]int, 4*n)
	for i := range 4 * n {
		minMass[i] = inf
	}
	return &Tree{arr, minMass, n}
}

func (t *Tree) Add(p int, g gripper) {
	var f func(i int, l int, r int)
	f = func(i int, l int, r int) {
		if l == r {
			t.arr[i] = append(t.arr[i], g)
			t.minMass[i] = min(t.minMass[i], g.m)
			return
		}
		mid := (l + r) / 2
		if p <= mid {
			f(2*i+1, l, mid)
		} else {
			f(2*i+2, mid+1, r)
		}
		t.minMass[i] = min(t.minMass[2*i+1], t.minMass[2*i+2])
	}
	f(0, 0, t.sz-1)
}

func (t *Tree) Play(p int, w int, f func(g gripper)) {
	var dfs func(i int, l int, r int)
	dfs = func(i int, l int, r int) {
		if t.minMass[i] > w || p < l {
			return
		}
		if l == r {
			for len(t.arr[i]) > 0 && t.arr[i][0].m <= w {
				f(t.arr[i][0])
				t.arr[i] = t.arr[i][1:]
			}
			if len(t.arr[i]) == 0 {
				t.minMass[i] = inf
			} else {
				t.minMass[i] = t.arr[i][0].m
			}
			return
		}
		mid := (l + r) / 2
		dfs(2*i+1, l, mid)
		dfs(2*i+2, mid+1, r)
		t.minMass[i] = min(t.minMass[2*i+1], t.minMass[2*i+2])
	}
	dfs(0, 0, t.sz-1)
}
