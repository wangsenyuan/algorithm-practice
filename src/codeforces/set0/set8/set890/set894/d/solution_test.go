package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	var buf strings.Builder

	writer := bufio.NewWriter(&buf)

	drive(reader, writer)

	writer.Flush()

	res := buf.String()
	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2 2
5
1 8
2 4`, `11
4
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `6 4
2
1
1
3
2
2 4
1 3
3 2
1 7`, `11
6
3
28
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `8 1
21725
80273
97276
78838
78474
1896
6570
7 5267977`, `41283845
`)
}
