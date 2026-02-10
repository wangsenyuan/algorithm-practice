package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 1
day
may
sun
fun
`
	expect := "aabb"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 1
day
may
gray
way
`
	expect := "aaaa"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 1
a
a
a
a
a
a
e
e
`
	expect := "aabb"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 1
day
may
sun
fun
test
hill
fest
thrill
`
	expect := "NO"
	runSample(t, s, expect)
}
