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
1 5
5 2
3 7
7 3`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 10
2 10
3 10
4 10
5 5
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `30
22 37
12 37
37 58
29 57
43 57
57 58
58 53
45 4
1 4
4 51
35 31
21 31
31 51
51 53
53 48
60 55
52 55
55 33
36 9
10 9
9 33
33 19
5 23
47 23
23 32
50 44
26 44
44 32
32 19
19 48
`
	expect := 31
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `50
73 1
65 73
16 65
57 65
33 16
34 57
98 16
84 98
55 34
64 84
80 55
75 64
28 75
20 75
42 75
88 42
50 20
48 28
32 48
58 88
92 76
76 53
53 15
15 1
1 10
10 71
71 37
37 95
95 63
63 92
45 97
97 51
51 96
96 12
12 62
62 31
31 5
5 29
29 19
19 49
49 6
6 40
40 18
18 22
22 17
17 46
46 72
72 82
82 14
14 14
`
	expect := 2
	runSample(t, s, expect)
}
