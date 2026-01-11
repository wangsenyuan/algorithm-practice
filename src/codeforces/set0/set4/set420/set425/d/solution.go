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
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	points := make([][]int, n)
	for i := range n {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		points[i] = []int{x, y}
	}
	return solve(points)
}

type point struct {
	id int
	x  int
	y  int
}

func solve(points [][]int) int {
	var mx, my int

	var offset int
	for _, p := range points {
		mx = max(mx, p[0])
		my = max(my, p[1])
		offset = max(offset, p[0]-p[1], p[1]-p[0])

	}

	n := len(points)

	arr := make([]point, n)

	for i, p := range points {
		arr[i] = point{i, p[0], p[1]}
	}

	slices.SortFunc(arr, func(a, b point) int {
		return cmp.Or(a.x-b.x, a.y-b.y)
	})

	lines := make([][]point, 2*offset+1)

	hoz := make([][]int, my+1)

	for _, cur := range arr {
		w := cur.y - cur.x
		lines[w+offset] = append(lines[w+offset], cur)
		hoz[cur.y] = append(hoz[cur.y], cur.id)
	}

	ver := make([][]int, mx+1)

	// 从最底下的line开始处理
	var res int

	mem := make([]int, mx+1)

	for _, vs := range lines {
		// y = x + w
		for _, v := range vs {
			mem[v.x]++
		}

		for i := len(vs) - 1; i >= 0; i-- {
			x1, y1 := vs[i].x, vs[i].y
			// vs[i].id 肯定是x1的最后一个
			hoz[y1] = hoz[y1][:len(hoz[y1])-1]

			for j1, j2 := len(hoz[y1])-1, len(ver[x1])-1; j1 >= 0 && j2 >= 0; {
				id1 := hoz[y1][j1]
				id2 := ver[x1][j2]
				dx := x1 - points[id1][0]
				dy := y1 - points[id2][1]
				if dx == dy {
					res += mem[x1-dx]
					j1--
					j2--
				} else if dx < dy {
					j1--
				} else {
					j2--
				}
			}
			ver[x1] = append(ver[x1], vs[i].id)
		}

		for _, v := range vs {
			mem[v.x]--
		}
	}

	return res
}
