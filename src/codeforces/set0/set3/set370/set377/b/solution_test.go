package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	assign, a, b, c := process(reader)
	expect := readString(reader)

	if len(assign) > 0 != (expect == "YES") {
		t.Fatalf("Sample expect %s, but got %v", expect, assign)
	}

	if len(assign) == 0 {
		return
	}

	expect_assign := readNNums(reader, len(assign))

	cnt := make([]int, len(b))
	check := func(arr []int) (days int, sum int) {
		clear(cnt)
		for i, st := range arr {
			st--
			if b[st] < a[i] {
				t.Fatalf("Student %d cannot fix bug %d", st+1, i+1)
			}
			cnt[st]++
			if cnt[st] == 1 {
				sum += c[st]
			}
		}
		days = slices.Max(cnt)
		return
	}

	d1, s1 := check(expect_assign)
	d2, s2 := check(assign)

	if d1 != d2 || s1 < s2 {
		t.Fatalf("Sample expect %v, but got %v", expect_assign, assign)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 4 9
1 3 1 2
2 1 3
4 3 6
YES
2 3 2 3
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 4 10
2 3 1 2
2 1 3
4 3 6
YES
1 3 1 3
	`)
}

func TestSample3(t *testing.T) {
	runSample(t, `3 4 9
2 3 1 2
2 1 3
4 3 6
YES
3 3 2 3
	`)
}
