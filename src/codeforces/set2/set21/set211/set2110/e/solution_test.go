package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, songs := process(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	for i := range len(res) {
		res[i]--
	}
	for i := 1; i < len(res); i++ {
		a := songs[res[i-1]]
		b := songs[res[i]]
		if a[0] != b[0] && a[1] != b[1] || a[0] == b[0] && a[1] == b[1] {
			t.Fatalf("Sample result %v, not correct", res)
		}

		if i >= 2 {
			c := songs[res[i-2]]
			if a[0] == b[0] && a[0] == c[0] || a[1] == b[1] && a[1] == c[1] {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSamaple1(t *testing.T) {
	s := `4
179 239
179 179
239 179
239 239`
	expect := true
	runSample(t, s, expect)
}

func TestSamaple2(t *testing.T) {
	s := `3
1 1
2 1
3 1`
	expect := false
	runSample(t, s, expect)
}


func TestSamaple3(t *testing.T) {
	s := `1
5 7`
	expect := true
	runSample(t, s, expect)
}

func TestSamaple4(t *testing.T) {
	s := `5
1 1
1 2
2 1
2 2
99 99`
	expect := false
	runSample(t, s, expect)
}

func TestSamaple5(t *testing.T) {
	s := `7
1 1
1 3
2 1
2 2
3 1
3 2
3 3`
	expect := true
	runSample(t, s, expect)
}