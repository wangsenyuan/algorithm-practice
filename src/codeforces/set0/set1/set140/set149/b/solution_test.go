package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, s string, expect int, expectRadixes []int) {
	cnt, radixes := solve(s)
	if cnt != expect {
		t.Errorf("Sample expect %d, but got %d", expect, cnt)
	}

	if expect <= 0 {
		return
	}

	if !slices.Equal(radixes, expectRadixes) {
		t.Errorf("Sample expect %v, but got %v", expectRadixes, radixes)
	}

}

func TestSample1(t *testing.T) {
	s := "11:20"
	radixes := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}
	runSample(t, s, len(radixes), radixes)
}

func TestSample2(t *testing.T) {
	s := "2A:13"
	expect := 0
	runSample(t, s, expect, nil)
}

func TestSample3(t *testing.T) {
	s := "000B:00001"
	expect := -1
	runSample(t, s, expect, nil)
}

func TestSample4(t *testing.T) {
	s := "00:21"
	expect := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}
	runSample(t, s, len(expect), expect)
}

func TestSample5(t *testing.T) {
	s := "0000P:0000E"
	expect := 0
	runSample(t, s, expect, nil)
}
