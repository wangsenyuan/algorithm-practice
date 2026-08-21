package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 10
x2 -10
`, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 2
+6 +7 /1 -13
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `8 1
+1 x2 x3 +4 +5 +6 -7 -8
`, 166666677)
}

func TestSample4(t *testing.T) {
	runSample(t, `9 864209753
-918273645 x564738291 /365107362 x734582911 -654321789 x998244353 +172519103 /482193765 /482091376
`, 601980218)
}

func TestAdditionBeforeOrAfterMultiplication(t *testing.T) {
	// The two orders yield 17 and 24, so the expected value is 41 / 2.
	runSample(t, `2 5
+7 x2
`, 500000024)
}
