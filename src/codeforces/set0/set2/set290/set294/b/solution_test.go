package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 12
1 3
2 15
2 5
2 1`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 10
2 1
2 4
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `50
1 24
1 16
1 33
2 34
1 26
2 35
1 39
2 44
2 29
2 28
1 44
2 48
2 50
2 41
2 9
1 22
2 11
2 27
1 12
1 50
2 49
1 17
2 43
2 6
1 39
2 28
1 47
1 45
2 32
1 43
2 40
1 10
1 44
2 31
2 26
2 15
2 20
1 49
1 36
2 43
2 8
1 46
2 43
2 26
1 30
1 23
2 26
1 32
2 25
2 42
`
	expect := 67
	runSample(t, s, expect)
}
