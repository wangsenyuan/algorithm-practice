package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a, b := readTwoNums(reader)
	res := solve(a, b)
	if len(res) == 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		for _, p := range res {
			fmt.Println(p[0], p[1])
		}
	}
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

type point struct {
	x int
	y int
}

func (this point) cross(that point) int {
	return this.x*that.y - this.y*that.x
}

// 判断两个向量是否垂直
func (v1 point) isPerpendicular(v2 point) bool {
	return v1.x*v2.x+v1.y*v2.y == 0 // 点积为0
}

func (this point) dist(that point) int {
	dx := this.x - that.x
	dy := this.y - that.y
	return dx*dx + dy*dy
}

func getPoints(a int) []point {
	var res []point
	for ax := 1; ax < a; ax++ {
		// ax * ax + ay * ay = a * a
		u := a*a - ax*ax
		v := int(math.Sqrt(float64(u)))
		if v*v == u {
			res = append(res, point{ax, v})
		}
		// v = ay
	}
	return res
}

func solve(a int, b int) [][]int {
	arr1 := getPoints(a)
	arr2 := getPoints(b)

	// a 有可能等于b

	for _, p1 := range arr1 {
		for _, dx1 := range []int{-1, 1} {
			dy1 := -dx1
			for _, p2 := range arr2 {
				if check(p1.x*dx1, p1.y*dy1, p2.x, p2.y) {
					return [][]int{
						{p1.x * dx1, p1.y * dy1},
						{p2.x, p2.y},
						{0, 0},
					}
				}
			}
		}
	}

	return nil
}

func check(x1 int, y1 int, x2 int, y2 int) bool {
	// 不能是一条直线
	a := point{x1, y1}
	b := point{x2, y2}
	c := a.cross(b)
	if c == 0 {
		return false
	}
	// 不是直线
	if y1 == y2 || x1 == x2 {
		// 第三条线平行于x/y
		return false
	}
	return a.isPerpendicular(b)
}
