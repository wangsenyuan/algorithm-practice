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
	s := `11
1 100
3 70
4
2
3 120
5
3 120
6
1 150
4
3 300
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 100
3 200
2
4
5
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7
1 20
2
6
4
6
6
2
`
	expect := 2
	runSample(t, s, expect)
}
