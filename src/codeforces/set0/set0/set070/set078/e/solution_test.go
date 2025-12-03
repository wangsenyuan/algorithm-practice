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
	s := `3 3
1YZ
1YY
100

0YZ
0YY
003
	`
	runSample(t, s, 2)
}

func TestSample2(t *testing.T) {
	s := `4 4
Y110
1Y1Z
1Y0Y
0100

Y001
0Y0Z
0Y0Y
0005
	`
	runSample(t, s, 3)
}

func TestSample3(t *testing.T) {
	s := `3 1
4Z1
908
146

3Z2
180
811
	`
	runSample(t, s, 24)
}
