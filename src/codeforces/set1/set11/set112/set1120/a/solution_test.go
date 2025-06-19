package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, k, a, b, ok, res := process(reader)
	expect := readNum(reader)
	if expect < 0 {
		if ok {
			t.Fatalf("Sample expect %d, but got %t", expect, ok)
		}
		return
	}
	if !ok {
		t.Fatalf("Sample expect %d, but got %t", expect, ok)
	}
	var arr []int
	for i, j := 0, 0; i < len(a); i++ {
		if j < len(res) && res[j] == i+1 {
			// remove i
			j++
			continue
		}
		arr = append(arr, a[i])
	}
	if len(arr) < n*k {
		t.Fatalf("Sample result %v, not correct", res)
	}
	// len(arr) >= n * k
	f1 := make([]int, X)
	for _, v := range b {
		f1[v]++
	}

	var cnt1 int
	for i := range X {
		if f1[i] > 0 {
			cnt1++
		}
	}

	f2 := make([]int, X)
	var cnt2 int

	add := func(x int) {
		f2[x]++
		if f2[x] == f1[x] {
			cnt2++
		}
	}

	rem := func(x int) {
		if f1[x] > 0 && f2[x] == f1[x] {
			cnt2--
		}
		f2[x]--
	}

	m := len(arr)

	check := func(l int, r int) bool {
		// cnt2 == cnt1
		u := l / k
		u++
		u += (m - r - 1) / k
		return u >= n
	}

	for l, r := 0, 0; r < m; r++ {
		add(arr[r])
		// 至少需要k个数 => r - l + 1 >= k => r >= l + k - 1
		for l+k-1 <= r && cnt2 == cnt1 {
			if check(l, r) {
				return
			}
			rem(arr[l])
			if cnt2 < cnt1 {
				add(arr[l])
				break
			}
			l++
		}
	}
	t.Fatalf("not found valid schema in the result %v", arr)
}

func TestSample1(t *testing.T) {
	runSample(t, `7 3 2 2
1 2 3 3 2 1 2
2 2
1
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `13 4 3 3
3 2 6 4 1 4 4 7 1 3 3 2 4
4 3 4
-1
	`)
}

func TestSample3(t *testing.T) {
	runSample(t, `13 4 1 3
3 2 6 4 1 4 4 7 1 3 3 2 4
4 3 4
9
	`)
}
