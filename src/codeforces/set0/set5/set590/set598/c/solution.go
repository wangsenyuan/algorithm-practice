package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, ans := drive(reader)
	fmt.Println(ans[0], ans[1])
}

func drive(reader *bufio.Reader) (points [][]int, ans []int) {
	var n int
	fmt.Fscan(reader, &n)
	points = make([][]int, n)
	for i := range n {
		points[i] = make([]int, 2)
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}
	ans = solve(n, points)
	return
}

type pt struct {
	id int
	x  int
	y  int
}

func (a pt) cross(b pt) int {
	return a.x*b.y - a.y*b.x
}

func (a pt) dot(b pt) int {
	return a.x*b.x + a.y*b.y
}

func (a pt) top() bool {
	return a.y < 0 || (a.y == 0 && a.x < 0)
}

func polarLess(a, b pt) bool {
	if a.top() != b.top() {
		return a.top()
	}
	return a.cross(b) > 0
}

func angleLess(a1, b1, a2, b2 pt) bool {
	p1 := pt{-1, a1.dot(b1), abs(a1.cross(b1))}
	p2 := pt{-1, a2.dot(b2), abs(a2.cross(b2))}
	return p1.cross(p2) > 0
}

func abs(a int) int {
	return max(a, -a)
}

func solve(n int, points [][]int) []int {
	arr := make([]pt, n)
	for i := range n {
		arr[i] = pt{id: i + 1, x: points[i][0], y: points[i][1]}
	}
	sort.Slice(arr, func(i, j int) bool {
		return polarLess(arr[i], arr[j])
	})

	// 算的不是距离，算的是角度的大小，所以平行的两个vector
	// 肯定是最近的两个
	ans := []int{0, 1}
	for i := 1; i < n; i++ {
		if angleLess(arr[i], arr[(i+1)%n], arr[ans[0]], arr[ans[1]]) {
			ans[0], ans[1] = i, (i+1)%n
		}
	}
	ans[0] = arr[ans[0]].id
	ans[1] = arr[ans[1]].id
	return ans
}

// 计算向量相对于原点的极角
func polarAngle(x, y int) float64 {
	return math.Atan2(float64(y), float64(x))
}

// 计算两个向量的非定向夹角
func nonOrientedAngle(x1, y1, x2, y2 int) float64 {
	angle1 := polarAngle(x1, y1)
	angle2 := polarAngle(x2, y2)

	diff := math.Abs(angle1 - angle2)

	// 返回较小的角度（非定向角总是 <= π）
	return min(diff, 2*math.Pi-diff)
}
