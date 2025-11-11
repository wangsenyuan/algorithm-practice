package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	if res == nil {
		fmt.Println("NO")
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, "YES")
	for _, p := range res {
		fmt.Fprintln(writer, p[0], p[1])
	}
}

func drive(reader *bufio.Reader) (ws []int, res [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	points := make([][]int, n)
	for i := range n {
		points[i] = make([]int, 2)
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}
	ws = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &ws[i])
	}
	res = solve(n, points, ws)
	return
}

func solve(n int, points [][]int, ws []int) [][]int {
	if ws[0] != 0 {
		return nil
	}
	// 找到每行最右的点，和每列最高的点
	var mx, my int
	for _, p := range points {
		mx = max(mx, p[0])
		my = max(my, p[1])
	}

	rows := make([]int, my+1)
	cols := make([]int, mx+1)

	for _, p := range points {
		x, y := p[0], p[1]
		rows[y] = max(rows[y], x)
		cols[x] = max(cols[x], y)
	}

	d := make([][]int, mx+1)
	for i := range mx + 1 {
		d[i] = make([]int, cols[i]+1)
		for j := range cols[i] + 1 {
			d[i][j] = 2
			if i == 0 || j == 0 {
				d[i][j] = 1
			}
		}
	}

	d[0][0] = 0

	offset := max(mx, my) + 2

	res := make([][]int, n)

	res[0] = []int{0, 0}

	lines := make([][]int, 2*offset+1)

	it := 1

	waitAtLines := func(x int, y int) {
		d := y - x
		lines[d+offset] = append(lines[d+offset], x)
	}

	var que [][]int
	que = append(que, []int{0, 0})

	for len(que) > 0 {
		for _, p := range que {
			x, y := p[0], p[1]
			z := y - x
			if z < -offset || z > offset {
				return nil
			}
			if x < rows[y] {
				d[x+1][y]--
				if d[x+1][y] == 0 {
					waitAtLines(x+1, y)
				}
			}
			if y < cols[x] {
				d[x][y+1]--
				if d[x][y+1] == 0 {
					waitAtLines(x, y+1)
				}
			}
		}

		var next [][]int
		for it < n {
			w := ws[it]
			if len(lines[offset+w]) == 0 {
				break
			}
			x := lines[offset+w][0]
			lines[offset+w] = lines[offset+w][1:]
			y := x + w
			res[it] = []int{x, y}
			it++
			next = append(next, []int{x, y})
		}
		que = next
	}

	if it < n {
		return nil
	}

	return res
}
