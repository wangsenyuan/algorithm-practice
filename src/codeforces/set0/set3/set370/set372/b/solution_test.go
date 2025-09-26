package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	drive(reader, writer)
	writer.Flush()

	res := buf.String()

	if expect != res {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 5 5
00101
00000
00001
01000
00001
1 2 2 4
4 5 4 5
1 2 5 2
2 2 4 5
4 2 5 3
`
	expect := `10
1
7
34
5
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 7 5
0000100
0000010
0011000
0000000
1 7 2 7
3 1 3 1
2 3 4 5
1 2 2 7
2 2 4 7
`
	expect := `3
1
16
27
52
`
	runSample(t, s, expect)
}
