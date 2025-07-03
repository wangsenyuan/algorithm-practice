package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readNum(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	runSample(t, `9555
4`)
}

func TestSample2(t *testing.T) {
	runSample(t, `10000000005
2`)
}

func TestSample3(t *testing.T) {
	runSample(t, `800101
3`)
}

func TestSample4(t *testing.T) {
	runSample(t, `45
1`)
}

func TestSample5(t *testing.T) {
	runSample(t, `1000000000000001223300003342220044555
17`)
}

func TestSample6(t *testing.T) {
	runSample(t, `19992000
1`)
}

func TestSample8(t *testing.T) {
	runSample(t, `310200
2`)
}