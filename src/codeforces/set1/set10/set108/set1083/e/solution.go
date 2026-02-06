package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(in *bufio.Reader) int {
	buf := make([]byte, 4096)
	var _i, _n int
	rc := func() byte {
		if _i == _n {
			_n, _ = in.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	rd := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n := rd()
	rects := make([][]int, n)
	for i := range n {
		rects[i] = []int{rd(), rd(), rd()}
	}
	return solve(rects)
}

type vec struct{ x, y int }

func (a vec) sub(b vec) vec { return vec{a.x - b.x, a.y - b.y} }
func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y }
func (a vec) det(b vec) int { return a.x*b.y - a.y*b.x }

func (a vec) detCmp(b vec) int {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w)
}

func solve(rects [][]int) int {
	slices.SortFunc(rects, func(a []int, b []int) int {
		return a[0] - b[0]
	})
	// now desc by y
	// dp[i] = max(dp[j] + (x[i] - x[j]) * y[i] - a[i])
	//  = dp[j] + x[i] * y[i] - x[j] * y[i] - a[i]
	//  = x[i] * y[i] - a[i] + dp[j] - x[j] * y[i]
	// -x[j] 作为斜率
	que := []vec{{0, 0}}

	var res int

	for _, rect := range rects {
		x, y, a := rect[0], rect[1], rect[2]
		p := vec{-y, 1}

		for len(que) > 1 && p.dot(que[0]) <= p.dot(que[1]) {
			que = que[1:]
		}

		// dp[i] = C + dp[j] - y * x[j]
		// {-y, 1} . {x[j], dp[j]}
		// -y * x[j] + dp[j]]

		cur := p.dot(que[0]) + x*y - a

		res = max(res, cur)

		v := vec{x, cur}

		for len(que) > 1 && que[len(que)-1].sub(que[len(que)-2]).detCmp(v.sub(que[len(que)-1])) >= 0 {
			que = que[:len(que)-1]
		}
		que = append(que, v)
	}

	return res
}
