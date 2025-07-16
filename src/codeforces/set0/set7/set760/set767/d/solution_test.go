package main

import (
	"bufio"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	k, a, b, ok, res := process(reader)

	expect := readNum(reader)

	if expect == -1 == ok {
		t.Fatalf("Sample expect %d, but got %t", expect, ok)
	}

	if !ok {
		return
	}

	if expect != len(res) {
		t.Fatalf("Sample expect %d, but got %d", expect, len(res))
	}

	// a的序号没有关系
	for _, i := range res {
		a = append(a, b[i-1])
	}

	sort.Ints(a)

	prev := -1
	var free int
	for i := 0; i < len(a); {
		free += (a[i] - prev) * k
		j := i
		for i < len(a) && a[i] == a[j] {
			i++
		}

		if i-j > free {
			t.Fatalf("Sample result %v, not correct", res)
		}
		free -= i - j
		prev = a[j]
	}
}

func TestSample1(t *testing.T) {
	s := `3 6 2
1 0 1
2 0 2 0 0 2
3`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 1 2
0 0 0
1
-1`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `2 1 2
0 1
0
1`
	runSample(t, s)
}
