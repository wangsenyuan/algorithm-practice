package main

import (
	"bufio"
	"strings"
	"testing"
)

func runDrive(t *testing.T, input string, expect int) {
	t.Helper()
	r := bufio.NewReader(strings.NewReader(input))
	res := drive(r)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	input := `4 7
1 4
6 9
4 7
3 5
`
	expect := 1
	runDrive(t, input, expect)
}

func TestSample2(t *testing.T) {
	input := `2 7
40 45
47 74
`
	expect := 2
	runDrive(t, input, expect)
}

func TestSample3(t *testing.T) {
	input := `74 77
22 23
21 23
11 20
15 19
22 26
26 27
10 14
18 23
22 22
16 18
3 23
6 6
2 22
13 24
13 24
8 12
25 26
3 26
5 17
7 7
8 21
2 6
19 22
6 16
2 27
10 18
15 16
10 12
26 27
9 12
5 5
21 22
24 27
26 26
12 16
26 27
1 6
17 18
8 14
4 11
26 27
5 14
25 25
13 17
22 27
5 9
19 21
8 15
19 19
6 8
11 22
8 12
11 15
10 24
22 24
8 10
25 27
7 7
7 13
15 17
24 25
26 26
27 27
26 27
25 26
2 4
19 25
4 8
25 27
19 23
16 27
15 21
14 19
2 15
`
	expect := 0
	runDrive(t, input, expect)
}

func TestShiftBeyondOriginalRightEndpoint(t *testing.T) {
	input := `1 43
1 31
`
	expect := 3
	runDrive(t, input, expect)
}

func TestLargeCostDoesNotOverflow(t *testing.T) {
	input := `10 0
1000000000000000000 1000000000000000000
1000000000000000000 1000000000000000000
1000000000000000000 1000000000000000000
1000000000000000000 1000000000000000000
1000000000000000000 1000000000000000000
1000000000000000000 1000000000000000000
1000000000000000000 1000000000000000000
1000000000000000000 1000000000000000000
1000000000000000000 1000000000000000000
1000000000000000000 1000000000000000000
`
	expect := 0
	runDrive(t, input, expect)
}
