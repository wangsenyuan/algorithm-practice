package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	m, a, cnt, res := drive(reader)
	if cnt != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, cnt)
	}
	if expect < 0 {
		return
	}

	var diff, level int
	n := len(a)
	for i := range n {
		if a[i] != res[i] {
			if res[i] > m || res[i] < 1 {
				t.Fatalf("Sample result %v is invalid", res)
			}
			diff++
		}
		if res[i]&1 == 0 {
			level++
		} else {
			level--
		}
	}
	if level != 0 {
		t.Fatalf("Sample result %v, odd number not equal to even number", res)
	}

	if diff != cnt {
		t.Fatalf("Sample result %v, diff %d not equal to cnt %d", res, diff, cnt)
	}
}

func TestSample1(t *testing.T) {
	s := `6 2
5 6 7 9 4 5`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8 6
7 7 7 7 8 8 8 8`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 1
4 2 1 10`
	expect := -1
	runSample(t, s, expect)
}
