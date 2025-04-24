package main

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	res, ps := process(bufio.NewReader(strings.NewReader(s)))
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	a := make(map[string]int)
	for _, p := range ps {
		i := strings.Index(p, " ")
		name := p[:i]
		cnt, _ := strconv.Atoi(p[i+1:])
		a[name] = cnt
	}
	n := len(res)
	h := make([]int, n)
	for j, cur := range res {
		i := strings.Index(cur, " ")
		name := cur[:i]
		tmp, _ := strconv.Atoi(cur[i+1:])
		// tmp is the height
		var cnt int
		for j1 := 0; j1 < j; j1++ {
			if h[j1] > tmp {
				cnt++
			}
		}
		if cnt != a[name] {
			t.Fatalf("Sample result %v, not correct", res)
		}
		h[j] = tmp
	}
}

func TestSample1(t *testing.T) {
	s := `4
a 0
b 2
c 0
d 0
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
vasya 0
petya 1
manya 3
dunay 3
`
	expect := false
	runSample(t, s, expect)
}

