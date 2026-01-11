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
	s := `1 3
.x.
`
	runSample(t, s, 3)
}

func TestSample2(t *testing.T) {
	s := `2 2
xx
xx
`
	runSample(t, s, 1)
}

func TestSample3(t *testing.T) {
	s := `2 2
..
..
`
	runSample(t, s, 10)
}
func TestSample4(t *testing.T) {
	s := `3 1
x
.
x
`
	runSample(t, s, 2)
}

func TestSample5(t *testing.T) {
	s := `50 5
..xxx
.xx..
x.x..
..xx.
xxx..
x..x.
.x.x.
x..xx
.x..x
...x.
..xxx
x.x..
x.xxx
x.x..
.xx.x
xxx.x
x..xx
x.x..
x.xxx
.xxx.
xxxxx
x..xx
.x.x.
...x.
...x.
x.x.x
.xx..
xxx..
..xxx
.x.xx
...x.
.x...
xxxxx
x.x..
x.x.x
..x.x
x...x
x....
xxx..
.xx.x
..x.x
x.xx.
.x.xx
x..x.
x....
....x
.x...
.xxxx
xxxxx
.x..x
`
	runSample(t, s, 464550945)
}