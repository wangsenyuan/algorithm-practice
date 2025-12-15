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
	s := "2 2 4"
	expect := 19
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "7 2444902 2613424"
	expect := 619309304
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "242954372 1561852 4674408"
	expect := 364019214
	runSample(t, s, expect)
}
