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
	s := `123123
1
2->00
`
	expect := 10031003
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `123123
1
3->
`
	expect := 1212
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `222
2
2->0
0->7
`
	expect := 777
	runSample(t, s, expect)
}
