package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Fatalf("expect %d, got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 2
1 2 1
2 3 1
1 1 1 1
`
	runSample(t, s, 3)
}

func TestSample2(t *testing.T) {
	s := `4 5
1 2 2
2 3 1
3 4 2
4 1 2
1 3 3
1 2 1 1
`
	runSample(t, s, 19)
}

func TestSample3(t *testing.T) {
	s := `3 3
1 2 2
2 3 1
3 1 2
1 1 1 1
`
	runSample(t, s, 4)
}

func TestSample4(t *testing.T) {
	s := `45 44
1 2 849
1 3 447
3 4 977
1 5 731
3 6 164
1 7 2
5 8 838
3 9 21
4 10 923
4 11 889
2 12 913
8 13 387
2 14 953
11 15 210
13 16 744
11 17 16
4 18 620
1 19 462
19 20 38
7 21 68
17 22 419
22 23 538
20 24 165
8 25 253
12 26 668
3 27 409
26 28 717
9 29 232
3 30 861
25 31 788
26 32 613
13 33 551
31 34 953
14 35 508
26 36 550
26 37 617
4 38 744
34 39 204
39 40 573
33 41 671
23 42 994
2 43 513
18 44 384
8 45 346
26 29 9 12
`
	runSample(t, s, 1087954848132212060)
}