package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		buf.WriteString(fmt.Sprintf("%d\n", process(reader)))
	}
	fmt.Print(buf.String())
}

func process(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	d := make([]int, n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &d[i], &a[i])
	}
	return solve(m, k, d, a)
}

func solve(n int, k int, d []int, a []int) int {
	type data struct {
		day int
		vol int
	}

	d = append(d, inf)
	a = append(a, 0)

	var que []data
	var curd int
	var res int
	var got int

	for i := range len(d) {
		cur := data{day: d[i], vol: a[i]}

		for len(que) > 0 && curd < cur.day {
			d, x := que[len(que)-1].day, que[len(que)-1].vol
			// que = que[1:]
			if d+k-1 < curd {
				// spoiled
				que = que[:len(que)-1]
				continue
			}

			if d > curd {
				got = 0
				curd = d
			}
			if n-got > x {
				got += x
				que = que[:len(que)-1]
			} else {
				sat := min(curd+(x-n+got)/n+1, min(d+k, cur.day))
				newx := x - (sat-curd)*n + got
				if newx <= 0 {
					que = que[:len(que)-1]
				} else {
					que[len(que)-1].vol = newx
				}
				res += sat - curd
				got = 0
				curd = sat
			}
		}

		que = append(que, cur)
	}

	return res
}

const inf = 1e18

func solve1(m int, k int, d []int, a []int) int {
	n := len(d)
	tr := Build(d[n-1] + k)
	var i int
	var res int
	for day := 1; day < d[n-1]+k; day++ {
		if i < n && d[i] == day {
			tr.Update(day, a[i])
			i++
		}
		j, ok := tr.Take(m)
		if ok && day-j < k {
			res++
		}
	}
	return res
}

type Tree struct {
	val  []int
	lazy []bool
}

func Build(n int) *Tree {
	val := make([]int, 4*n)
	lazy := make([]bool, 4*n)
	return &Tree{val: val, lazy: lazy}
}

func (tr *Tree) push(i int) {
	if tr.lazy[i] {
		tr.val[i*2+1] = 0
		tr.lazy[i*2+1] = true
		tr.val[i*2+2] = 0
		tr.lazy[i*2+2] = true
		tr.lazy[i] = false
	}
}

func (tr *Tree) Update(pos int, v int) {
	var loop func(i int, l int, r int)
	loop = func(i int, l int, r int) {
		if l == r {
			tr.val[i] += v
			return
		}
		tr.push(i)
		mid := (l + r) >> 1
		if pos <= mid {
			loop(i*2+1, l, mid)
		} else {
			loop(i*2+2, mid+1, r)
		}
		tr.val[i] = tr.val[i*2+1] + tr.val[i*2+2]
	}
	loop(0, 0, len(tr.val)/4-1)
}

func (tr *Tree) Take(x int) (int, bool) {
	var loop func(i int, l int, r int, x int) (res int, ok bool)
	loop = func(i int, l int, r int, x int) (res int, ok bool) {
		if l == r {
			ok = tr.val[i] >= x
			tr.val[i] = max(0, tr.val[i]-x)
			return l, ok
		}
		tr.push(i)
		mid := (l + r) >> 1

		if tr.val[i*2+2] >= x {
			res, ok = loop(i*2+2, mid+1, r, x)
		} else {
			x -= tr.val[i*2+2]
			tr.val[i*2+2] = 0
			tr.lazy[i*2+2] = true
			res, ok = loop(i*2+1, l, mid, x)
		}
		tr.val[i] = tr.val[i*2+1] + tr.val[i*2+2]
		return
	}

	return loop(0, 0, len(tr.val)/4-1, x)
}
