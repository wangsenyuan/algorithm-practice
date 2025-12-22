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
		t.Errorf("Sample failed.\nExpect:\n%s\nGot:\n%s\n", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6
1 2 3 4 3 5
3
1 2
3 5
2 6
`
	expect := `2
2
4
`
	runSample(t, s, expect)
}
