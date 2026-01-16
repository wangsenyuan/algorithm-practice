package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1233
0213
2020
0303
`
	expect := 108
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
00300
00300
33333
00300
00300
`
	expect := 19683
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
00003
02030
00300
03020
30000
`
	expect := 108
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
21312
10003
10002
10003
23231
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5
12131
12111
12112
21311
21212
`
	expect := 24
	runSample(t, s, expect)
}
