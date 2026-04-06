package main

import "testing"

func check(t *testing.T, n int64, expect bool) {
	t.Helper()
	ok, ans := solve(n)
	if ok != expect {
		t.Fatalf("expect %v, got %v for n = %d", expect, ok, n)
	}
	if !ok {
		return
	}

	num := int64(0)
	den := n
	for _, cur := range ans {
		a, b := cur[0], cur[1]
		if !(1 <= a && a < b && 1 < b && b < n && n%b == 0) {
			t.Fatalf("bad fraction %v for n = %d", cur, n)
		}
		num += a * (n / b)
	}
	if num != n-1 {
		t.Fatalf("wrong sum %d/%d for n = %d", num, den, n)
	}
}

func TestSample1(t *testing.T) {
	check(t, 2, false)
}

func TestSample2(t *testing.T) {
	check(t, 6, true)
}

func TestPrimePowers(t *testing.T) {
	for _, n := range []int64{3, 4, 5, 8, 9, 16, 25} {
		check(t, n, false)
	}
}

func TestComposite(t *testing.T) {
	for _, n := range []int64{10, 12, 15, 18, 20, 30, 42} {
		check(t, n, true)
	}
}
