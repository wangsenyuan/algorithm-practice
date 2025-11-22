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
	s := `3
add 1
remove
add 2
add 3
remove
remove
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
add 3
add 2
add 1
remove
add 4
remove
remove
remove
add 6
add 7
add 5
remove
remove
remove
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `11
add 10
add 9
add 11
add 1
add 5
add 6
remove
add 3
add 8
add 2
add 4
remove
remove
remove
remove
remove
add 7
remove
remove
remove
remove
remove
`
	expect := 2
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := `4
add 1
add 3
remove
add 4
add 2
remove
remove
remove
`
	expect := 2
	runSample(t, s, expect)
}
