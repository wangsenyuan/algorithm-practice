package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.10f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, l, v1, v2, k int
	fmt.Fscan(reader, &n, &l, &v1, &v2, &k)
	return solve(n, l, v1, v2, k)
}

func solve(n int, l int, v1 int, v2 int, k int) float64 {

	calc := func(y float64, d float64) float64 {
		// x / v2 + (d - x) / v1 = y
		// x / v2 - x / v1 + d / v1 = y
		y = d/float64(v1) - y
		x := y * float64(v1*v2) / float64(v2-v1)
		return min(x, d)
	}

	check := func(expect float64) bool {
		if expect >= float64(l)/float64(v1) {
			// 步行
			return true
		}

		var t float64
		var pos float64
		// 车辆所在的位置
		var bus float64
		m := (n + k - 1) / k

		for range m {
			// 这批学生，目前的位置在pos，车在bus处
			dt := (bus - pos) / float64(v1+v2)
			t += dt
			pos += dt * float64(v1)
			bus = pos
			if t > expect {
				return false
			}
			// 这批学生还有t1 = expect - t的时间到达目的地，假设在车上x，那么剩余 l - pos - x 要走路
			// x / v2 + (l - pos - x) / v1 = t1
			x := calc(expect-t, float64(l)-pos)
			// 车辆必须运行这么远，才能将学生准时送达
			bus += x
			t += x / float64(v2)
			pos += x / float64(v2) * float64(v1)
			if t > expect {
				return false
			}
		}

		return true
	}

	var lo float64
	hi := float64(l) / float64(v1)

	for range 100 {
		mid := (lo + hi) / 2
		if check(mid) {
			hi = mid
		} else {
			lo = mid
		}
	}

	return (lo + hi) / 2
}
