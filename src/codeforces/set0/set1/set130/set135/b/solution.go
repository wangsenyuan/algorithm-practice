package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	_, res := drive(reader)
	if len(res) == 0 {
		fmt.Fprintln(writer, "NO")
	} else {
		fmt.Fprintln(writer, "YES")
		for _, cur := range res {
			s := fmt.Sprintf("%v", cur)
			fmt.Fprintln(writer, s[1:len(s)-1])
		}
	}
}

func drive(reader *bufio.Reader) (points []Point, res [][]int) {
	points = make([]Point, 8)
	for i := range 8 {
		fmt.Fscan(reader, &points[i].x, &points[i].y)
	}
	res = solve(points)
	return
}

type Point struct {
	x int
	y int
}

func (a Point) sub(b Point) vector {
	return vector{a.x - b.x, a.y - b.y}
}

type vector struct {
	x int
	y int
}

func (a vector) cross(b vector) int {
	return a.x*b.y - a.y*b.x
}

func (a vector) dot(b vector) int {
	return a.x*b.x + a.y*b.y
}

func (a vector) len2() int {
	return a.x*a.x + a.y*a.y
}

func (p Point) add(v vector) Point {
	return Point{p.x + v.x, p.y + v.y}
}

func (v vector) rotate90() vector {
	return vector{-v.y, v.x}
}

func solve(points []Point) [][]int {

	cache := make(map[Point]int)
	for i := range 8 {
		cache[points[i]] = i
	}

	checkRestRect := func(used int) []int {
		var arr []int
		for i := range 8 {
			if used&(1<<i) == 0 {
				arr = append(arr, i)
			}
		}

		return orderRectangle(points, arr)
	}

	for i := range 8 {
		for j := range 8 {
			if i != j {
				v1 := points[j].sub(points[i])
				// d1 := v1.len2()
				p2 := points[i].add(v1.rotate90())
				if u, ok := cache[p2]; ok {
					// k != i & k != j holds
					p3 := points[j].add(v1.rotate90())
					if v, ok := cache[p3]; ok {
						// i, j, v, u are square corners
						flag := 1<<i | 1<<j | 1<<v | 1<<u
						rect := checkRestRect(flag)
						if len(rect) != 0 {
							return [][]int{
								{i + 1, j + 1, v + 1, u + 1},
								rect,
							}
						}
					}
				}
			}
		}
	}

	return nil
}

func orderRectangle(points []Point, arr []int) []int {
	a := append([]int(nil), arr...)
	for {
		if isRectangleOrder(points, a) {
			return []int{a[0] + 1, a[1] + 1, a[2] + 1, a[3] + 1}
		}
		if !nextPermutation(a) {
			break
		}
	}
	return nil
}

func isRectangleOrder(points []Point, arr []int) bool {
	for i := range 4 {
		a := points[arr[(i+3)%4]]
		b := points[arr[i]]
		c := points[arr[(i+1)%4]]
		v1 := a.sub(b)
		v2 := c.sub(b)
		if v1.len2() == 0 || v2.len2() == 0 || v1.dot(v2) != 0 {
			return false
		}
	}
	return true
}

func nextPermutation(a []int) bool {
	i := len(a) - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := len(a) - 1
	for a[j] <= a[i] {
		j--
	}
	a[i], a[j] = a[j], a[i]
	for l, r := i+1, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	return true
}
