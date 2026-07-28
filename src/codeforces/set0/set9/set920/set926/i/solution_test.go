package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect string) {
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	input := `1
05:43
`
	expect := `23:59`
	runSample(t, input, expect)
}

func TestSample2(t *testing.T) {
	input := `4
22:00
03:21
16:03
09:59
`
	expect := `06:37`
	runSample(t, input, expect)
}
