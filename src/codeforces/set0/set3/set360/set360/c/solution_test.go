package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 2
yz
`, 26)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 3
yx
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 7
abcd
`, 21962)
}

func TestSample4(t *testing.T) {
	runSample(t, `1 1
a
`, 25)
}

func TestSample5(t *testing.T) {
	runSample(t, `1 0
z
`, 26)
}

func TestSample6(t *testing.T) {
	runSample(t, `2 0
yz
`, 25*26)
}

func TestSample7(t *testing.T) {
	runSample(t, `47 0
tmvdxhumdsbypyocuydhqevkgfijsbdoucgugqkruktmlem
`, 129331153)
}
