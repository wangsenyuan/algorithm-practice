package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect string) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 42
3 8 4
`, "Yes")
}

func TestSample2(t *testing.T) {
	runSample(t, `3 1
3 8 4
`, "No")
}

func TestSample3(t *testing.T) {
	runSample(t, `20 83298426
748 169 586 329 972 529 432 519 408 587 138 249 656 114 632 299 984 755 404 772
`, "Yes")
}

func TestSample4(t *testing.T) {
	runSample(t, `20 85415869
748 169 586 329 972 529 432 519 408 587 138 249 656 114 632 299 984 755 404 772
`, "No")
}
