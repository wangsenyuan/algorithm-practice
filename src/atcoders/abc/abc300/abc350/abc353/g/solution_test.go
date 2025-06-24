package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readNum(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6 3
4
5 30
2 10
4 25
2 15
49`)
}

func TestSample2(t *testing.T) {
	runSample(t, `6 1000000000
4
5 30
2 10
4 25
2 15
0`)
}

func TestSample3(t *testing.T) {
	runSample(t, `50 10
15
37 261
28 404
49 582
19 573
18 633
3 332
31 213
30 377
50 783
17 798
4 561
41 871
15 525
16 444
26 453
5000`)
}

func TestSample4(t *testing.T) {
	runSample(t, `50 1000000000
15
30 60541209756
48 49238708511
1 73787345006
24 47221018887
9 20218773368
34 40025202486
14 28286410866
24 82115648680
37 62913240066
14 92020110916
24 20965327730
32 67598565422
39 79828753874
40 52778306283
40 67894622518
606214471001`)
}
