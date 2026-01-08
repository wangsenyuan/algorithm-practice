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
	fmt.Printf("%.8f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, w, v, u int
	fmt.Fscan(reader, &n, &w, &v, &u)
	points := make([][]int, n)
	for i := range n {
		points[i] = make([]int, 2)
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}
	return solve(w, v, u, points)
}

const inf = 1 << 60

func solve(w int, v int, u int, points [][]int) float64 {
	ok := true
	for _, p := range points {
		if p[0]*u < p[1]*v {
			ok = false
			break
		}
	}
	if ok {
		return float64(w) / float64(u)
	}

	slices.SortFunc(points, func(a []int, b []int) int {
		return cmp.Or(a[1]-b[1], a[0]-b[0])
	})

	var x, y int
	var t float64

	n := len(points)

	for i := range n {
		if points[i][0] < x {
			continue
		}
		x = points[i][0]
		t = max(float64(x)/float64(v), t+float64(points[i][1]-y)/float64(u))
		y = points[i][1]
	}
	t += float64(w-y) / float64(u)
	return t
}

func solve1(w int, v int, u int, points [][]int) float64 {
	lx, rx := inf, -inf
	var ty int
	for _, p := range points {
		lx = min(lx, p[0])
		rx = max(rx, p[0])
		ty = max(ty, p[1])
	}

	for i, p := range points {
		if p[0] == lx {
			shift(points, i)
			break
		}
	}
	n := len(points)

	var lower [][]int
	var pos int
	for pos < n && points[pos][0] < rx {
		lower = append(lower, points[pos])
		pos++
	}
	lower = append(lower, points[pos])
	var upper [][]int
	for pos < n {
		upper = append(upper, points[pos])
		pos++
	}
	upper = append(upper, points[0])
	slices.Reverse(upper)

	// 然后来计算，每个点到达 x = 0 的时刻， 有可能是负值
	a := calcArriveTime(upper, v)
	b := calcArriveTime(lower, v)

	// 当前时间和当前位置
	var t float64
	var y float64

	playUpper := func(i int, j int) {
		dt := a[i] - t

		if i == 0 && y == 0 || i > 0 && float64(upper[i-1][1]) <= y {
			y1 := y + float64(u)*dt
			if y1 >= float64(upper[i][1]) {
				y = y1
				return
			}
		}
		if j == 0 || j == len(lower) {
			// 即使跑到车里面也没有关系，下次会被调整回来
			y += dt * float64(u)
			return
		}

		y0 := calc(lower[j-1], lower[j], float64(upper[i][0])-float64(lower[j-1][0]))
		y = min(y0, y+dt*float64(u))
	}

	playerLower := func(i int, j int) {
		// i == 0 或者 i == len(upper) 不会发生
		// j > 0 也一定会成立
		dt := b[j] - t
		if y >= float64(upper[i-1][1]) {
			y1 := calc(upper[i-1], upper[i], float64(lower[j][0])-float64(upper[i-1][0]))
			if y1 > y+float64(u)*dt {
				// 要被撞到了, 必须退回到安全位置
				y = min(float64(lower[j][1]), y+float64(u)*dt)
			} else {
				y += float64(u) * dt
			}
			return
		}
		y = min(float64(lower[j][1]), y+float64(u)*dt)
	}

	for i, j := 0, 0; y < float64(ty); {
		for i < len(upper) && a[i] <= t {
			i++
		}
		if i == len(upper) {
			break
		}
		for j < len(lower) && b[j] <= t {
			j++
		}
		// 现在看一下哪个在前面
		if j == len(lower) || i < len(upper) && upper[i][0] <= lower[j][0] {
			// x := upper[i][0]
			playUpper(i, j)
			t = a[i]
			i++
		} else {
			playerLower(i, j)
			t = b[j]
			j++
		}

	}

	return t + (float64(w)-y)/float64(u)
}

func calc(first []int, second []int, dt float64) float64 {
	if first[0] == second[0] || first[1] == second[1] {
		// 垂直线, 或者水平线
		return float64(second[1])
	}
	slop := float64(second[1]-first[1]) / float64(second[0]-first[0])
	// 这个不大对
	return float64(first[1]) + slop*dt
}

func calcArriveTime(arr [][]int, v int) []float64 {
	n := len(arr)
	res := make([]float64, n)
	for i, cur := range arr {
		x := cur[0]
		res[i] = float64(x) / float64(v)
	}
	return res
}

func shift[T any](arr []T, k int) {
	slices.Reverse(arr[:k])
	slices.Reverse(arr[k:])
	slices.Reverse(arr)
}
