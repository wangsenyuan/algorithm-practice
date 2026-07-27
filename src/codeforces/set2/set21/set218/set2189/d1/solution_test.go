package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if len(res) != 1 || res[0] != expect {
		t.Errorf("Sample expect %v, but got %#v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
3 3
001
`, -1)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
3 1
111
`, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, `1
4 100
1001
`, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
6 100
111001
`, 96)
}

func TestSample5(t *testing.T) {
	runSample(t, `1
6 100
111101
`, 64)
}

func TestSample6(t *testing.T) {
	runSample(t, `1
5 8
10001
`, 12)
}

func TestSample7(t *testing.T) {
	runSample(t, `1
4 100
1110
`, -1)
}

func TestSample8(t *testing.T) {
	runSample(t, `1
21 123456789
111000111000111000111
`, 336892528)
}

func TestSample9(t *testing.T) {
	runSample(t, `1
3 4
101
`, 2)
}
