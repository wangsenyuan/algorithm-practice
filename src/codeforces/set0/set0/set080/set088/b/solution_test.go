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
	s := `2 2 1
ab
cd
1
A`
	runSample(t, s, -1)
}

func TestSample2(t *testing.T) {
	s := `2 2 1
ab
cd
1
e
`
	runSample(t, s, -1)
}

func TestSample3(t *testing.T) {
	s := `2 2 1
ab
cS
5
abcBA`
	runSample(t, s, 1)
}

func TestSample4(t *testing.T) {
	s := `3 9 4
qwertyuio
asdfghjkl
SzxcvbnmS
35
TheQuIcKbRoWnFOXjummsovertHeLazYDOG`
	runSample(t, s, 2)
}

func TestSample5(t *testing.T) {
	s := `8 6 4
efvmov
keofnw
pwajpe
knptky
Sibruu
rgdukk
bsxosd
hovgSe
10
ECreruXmsC`
	runSample(t, s, -1)
}
