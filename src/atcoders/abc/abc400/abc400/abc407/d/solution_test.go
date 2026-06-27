package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 4
1 2 3 8
4 0 7 10
5 2 4 2
`, 15)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 11
1 2 4 8 16 32 64 128 256 512 1024
`, 2047)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 5
74832 16944 58683 32965 97236
52995 43262 51959 40883 58715
13846 24919 65627 11492 63264
29966 98452 75577 40415 77202
`, 131067)
}
