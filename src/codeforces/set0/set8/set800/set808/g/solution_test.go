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
	s := `winlose???winl???w??
win
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `glo?yto?e??an?
or
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `??c?????
abcab
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `emnd?t??m?gd?t?p?s??m?dp??t???m?????m?d?ydo????????i??u?d??dp??h??d?tdp???cj?dm?dpxf?hsf??rdmt?pu?tw
dmtdp
`
	expect := 11
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `irsdljdahusytoclelxidaaiaiaicaiaiaiaiiaiaiyyexmohdwmeyycaiaiaitclluaiaiaiznxweleaiaiaiixdwehyruhizbc
aiaiai
`
	expect := 6
	runSample(t, s, expect)
}
