package main

import (
	"testing"
)

func abs(num int) int {
	return max(num, -num)
}

func runSample(t *testing.T, x, y int, expect []int) {
	res := solve(x, y)
	if len(res) != 2 {
		t.Fatalf("solve(%d, %d) expect 2 values, got %v", x, y, res)
	}
	p, q := res[0], res[1]
	if p&q != 0 {
		t.Fatalf("solve(%d, %d) = %v: p & q must be 0, got %d & %d = %d", x, y, res, p, q, p&q)
	}
	expectDist := abs(x-expect[0]) + abs(y-expect[1])
	gotDist := abs(x-p) + abs(y-q)
	if gotDist != expectDist {
		t.Fatalf("solve(%d, %d) = %v (dist %d); expect dist %d (e.g. %v)", x, y, res, gotDist, expectDist, expect)
	}
}

func TestSample1(t *testing.T) { runSample(t, 0, 0, []int{0, 0}) }
func TestSample2(t *testing.T) { runSample(t, 1, 1, []int{2, 1}) }
func TestSample3(t *testing.T) { runSample(t, 3, 6, []int{3, 8}) }
func TestSample4(t *testing.T) { runSample(t, 7, 11, []int{6, 9}) }
func TestSample5(t *testing.T) { runSample(t, 4, 4, []int{4, 3}) }
func TestSample6(t *testing.T) { runSample(t, 123, 321, []int{128, 321}) }
func TestSample7(t *testing.T) { runSample(t, 1073741823, 1073741822, []int{1073741824, 1073741822}) }

