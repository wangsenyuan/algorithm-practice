package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `abrakadabra
aba
`, 51)
}

func TestSample2(t *testing.T) {
	runSample(t, `aaaaa
a
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `rdddrdtdcdrrdcredctdordoeecrotet
dcre
`, 263)
}
