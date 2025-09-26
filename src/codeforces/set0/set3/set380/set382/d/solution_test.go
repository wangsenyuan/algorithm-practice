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
	s := `1 1
#`
	runSample(t, s, 0)
}

func TestSample2(t *testing.T) {
	s := `3 4
####
#>^#
####`
	runSample(t, s, 3)
}

func TestSample3(t *testing.T) {
	s := `3 4
####
#><#
####`
	runSample(t, s, -1)
}

func TestSample4(t *testing.T) {
	s := `7 5
#####
##v##
##v##
#####
##^##
##^##
#####`
	runSample(t, s, 4)
}

func TestSample5(t *testing.T) {
	s := `7 5
#####
##v##
##v##
##<##
##^##
##^##
#####`
	runSample(t, s, 5)
}
