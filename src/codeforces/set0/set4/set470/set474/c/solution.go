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
	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	moles := make([][]int, 4*n)
	for i := range n * 4 {
		moles[i] = make([]int, 4)
		fmt.Fscan(reader, &moles[i][0], &moles[i][1], &moles[i][2], &moles[i][3])
	}
	return solve(n, moles)
}

type Point struct {
	x, y int
}

func (this Point) Sub(that Point) Point {
	return Point{this.x - that.x, this.y - that.y}
}

func (this Point) Cross(that Point) int {
	a := this.x * that.y
	b := this.y * that.x
	return a - b
}

func (this Point) Dot(that Point) int {
	return this.x*that.x + this.y*that.y
}

func cross(a, b, c Point) int {
	return b.Sub(a).Cross(c.Sub(a))
}

func dot(a, b, c Point) int {
	return b.Sub(a).Dot(c.Sub(a))
}

func (this Point) distance2(that Point) int {
	dx := this.x - that.x
	dy := this.y - that.y
	return dx*dx + dy*dy
}

func solve(n int, moles [][]int) []int {

	check := func(a, b, c, d Point) bool {
		arr := []Point{a, b, c, d}
		o := arr[0]
		for i := 1; i < 4; i++ {
			if arr[i].x < o.x || arr[i].x == o.x && arr[i].y < o.y {
				o = arr[i]
			}
		}

		// 按照和o的角度进行排序
		slices.SortFunc(arr, func(p1, p2 Point) int {
			deg := cross(o, p1, p2)
			if deg != 0 {
				return -deg
			}
			return p1.distance2(o) - p2.distance2(o)
		})
		// 0, 1, 2, 3, 是4个数序的角度，
		for i := range 4 {
			l := (i + 3) % 4
			r := (i + 1) % 4
			if dot(arr[i], arr[l], arr[r]) != 0 {
				return false
			}
		}
		// 它们只是一个矩形，还需要知道距离
		w := arr[0].distance2(arr[1])
		if w == 0 {
			return false
		}
		for i := 1; i < 4; i++ {
			if arr[i].distance2(arr[(i+1)%4]) != w {
				return false
			}
		}
		return true
	}
	find := func(s int) int {
		pos := make([][]Point, 4)
		for j := 4 * s; j < 4*(s+1); j++ {
			cur := moles[j]
			x, y, a, b := cur[0], cur[1], cur[2], cur[3]
			// 逆时针旋转90度: (x,y) 绕 (a,b) 旋转
			// 公式: nx = a - (y - b) = a - y + b
			//      ny = b + (x - a) = b + x - a
			for range 4 {
				pos[j-4*s] = append(pos[j-4*s], Point{x, y})
				nx := a - y + b
				ny := b + x - a
				x, y = nx, ny
			}
			// slices.Reverse(pos[j-4*s])
		}
		best := 100
		for i, a := range pos[0] {
			for j, b := range pos[1] {
				for k, c := range pos[2] {
					for l, d := range pos[3] {
						if check(a, b, c, d) {
							best = min(best, i+j+k+l)
						}
					}
				}
			}
		}
		if best == 100 {
			return -1
		}
		return best
	}

	ans := make([]int, n)
	for i := range n {
		ans[i] = find(i)
	}

	return ans
}

func getValue(b bool) int {
	if b {
		return 1
	}
	return 0
}
