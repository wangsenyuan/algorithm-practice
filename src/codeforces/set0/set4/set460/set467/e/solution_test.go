package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	a, res := drive(reader)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	for i, j := 0, 0; i < len(res); i++ {
		for j < len(a) && res[i] != a[j] {
			j++
		}
		if j == len(a) {
			t.Fatalf("Sample result %v, not a sub sequece of %v", res, a)
		}
		j++
	}
	for i := 0; i < len(res); i += 4 {
		if res[i] != res[i+2] || res[i+1] != res[i+3] {
			t.Fatalf("Sample result %v, not valid", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4
3 5 3 5
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10
35 1 2 1 2 35 100 200 100 200
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
1 2 2 2 2
`
	expect := 4
	runSample(t, s, expect)
}
