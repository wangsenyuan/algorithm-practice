package main

import (
	"bufio"
	"cmp"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if len(res) != len(expect) {
		t.Fatalf("length not match, expect %v, got %v", expect, res)
	}
	if len(expect) == 0 {
		return
	}

	type point struct {
		x int
		y int
	}

	checkAndCal := func(rect []int) int {
		var arr []point
		for i := 0; i < len(rect); i += 2 {
			arr = append(arr, point{x: rect[i], y: rect[i+1]})
		}
		slices.SortFunc(arr, func(a, b point) int {
			return cmp.Or(a.x-b.x, a.y-b.y)
		})
		// 面积为0时, 无法按照正常的rect判断
		if arr[2].x == arr[0].x || arr[1].y == arr[0].y {
			return 0
		}
		if arr[0].x != arr[1].x || arr[0].y != arr[2].y || arr[2].x != arr[3].x {
			t.Fatalf("invalid rect %v", rect)
		}
		dx := arr[2].x - arr[0].x
		dy := arr[1].y - arr[0].y
		return dx * dy
	}

	w := checkAndCal(expect)
	v := checkAndCal(res)

	if w != v {
		t.Fatalf("Sample xpect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `16
-5 1 1 2 2 3 3 4 4 5 5 6 6 7 7 10
`, []int{1, 2, 1, 7, 6, 2, 6, 7})
}

func TestSample2(t *testing.T) {
	runSample(t, `8
0 0 -1 2 2 1 1 3
`, nil)
}

func TestSample3(t *testing.T) {
	runSample(t, `8
0 0 0 0 0 5 0 5
`, []int{0, 0, 0, 5, 0, 0, 0, 5})
}
