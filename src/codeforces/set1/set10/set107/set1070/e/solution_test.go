package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	m, x, c, res := drive(reader)

	if res[0] != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	d := res[1]
	var took int
	var cnt int
	var cur int
	for i := range c {
		if c[i] <= d {
			cnt++
			cur += c[i]
			took += c[i]
		}
		if cnt > 0 && cnt%m == 0 && cnt < expect {
			// 不是最后一个任务
			// take a break
			took += cur
			cur = 0
		}
		if cnt == expect {
			break
		}
	}

	if took > x {
		t.Fatalf("Sample result %v, took %d, but expected %d", res, took, x)
	}
}

func TestSample1(t *testing.T) {
	s := `5 2 16
5 6 1 4 7
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 3 30
5 6 1 4 7
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 4 15
12 5 15 7 20 17
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 1 50
100
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `11 1 29
6 4 3 7 5 3 4 7 3 5 3
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `7 1 5
1 1 1 1 1 1 1
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `5 2 18
2 3 3 7 5
`
	expect := 4
	runSample(t, s, expect)
}
