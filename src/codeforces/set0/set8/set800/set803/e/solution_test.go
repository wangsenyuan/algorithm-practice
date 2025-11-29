package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, s, res := drive(reader)

	if res == expect {
		return
	}
	if expect == "NO" || res == "NO" {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	var score int
	for i := range s {
		if s[i] != '?' && s[i] != res[i] {
			t.Fatalf("Sample result %s, not correct", res)
		}

		switch res[i] {
		case 'W':
			score++
		case 'L':
			score--
		}

		if abs(score) >= k && i < len(s)-1 {
			t.Fatalf("Sample result %s, not correct, as it reachs %d, at position %d", res, k, i)
		}
	}

	if score != k && score != -k {
		t.Fatalf("Sample result %s, not getting the correct score %d, got %d", res, k, score)
	}
}

func TestSample1(t *testing.T) {
	s := `3 2
L??
`
	expect := "LDL"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1
W??
`
	expect := "NO"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `20 5
?LLLLLWWWWW?????????
`
	expect := "WLLLLLWWWWWWWWLWLWDW"
	runSample(t, s, expect)
}
