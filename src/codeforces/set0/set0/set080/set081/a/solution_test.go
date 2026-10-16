package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `hhoowaaaareyyoouu
`, "wre")
}

func TestSample2(t *testing.T) {
	runSample(t, `reallazy
`, "rezy")
}

func TestSample3(t *testing.T) {
	runSample(t, `abacabaabacabaa
`, "a")
}
