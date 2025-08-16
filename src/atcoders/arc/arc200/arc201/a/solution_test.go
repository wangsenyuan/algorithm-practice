package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
3 1 4
1 5 3`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1
1 1 1`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
5 7 5
1 11 99
3 1 2`
	expect := 7
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
1000000000 1000000000 1000000000
1000000000 1000000000 1000000000
1000000000 1000000000 1000000000
1000000000 1000000000 1000000000
1000000000 1000000000 1000000000`
	expect := 2500000000
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `6
835549144 866512240 105679868
473233032 625162103 823002638
125467290 37501686 380787083
8043910 721085797 254272563
97327826 744196952 18713225
978152989 90127986 33086297`
	expect := 998830769
	runSample(t, s, expect)
}
