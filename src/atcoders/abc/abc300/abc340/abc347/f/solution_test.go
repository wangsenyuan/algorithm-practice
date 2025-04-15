package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7 3
3 1 4 1 5 9 2
6 5 3 5 8 9 7
9 3 2 3 8 4 6
2 6 4 3 3 8 3
2 7 9 5 0 2 8
8 4 1 9 7 1 6
9 3 9 9 3 7 5
`
	expect := 154
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 1
3 1 4 1 5 9 2
6 5 3 5 8 9 7
9 3 2 3 8 4 6
2 6 4 3 3 8 3
2 7 9 5 0 2 8
8 4 1 9 7 1 6
9 3 9 9 3 7 5
`
	expect := 27
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `16 4
74 16 58 32 97 52 43 51 40 58 13 24 65 11 63 29
98 75 40 77 15 50 83 85 35 46 38 37 56 38 63 55
95 42 10 70 53 40 25 10 70 32 33 19 52 79 74 58
33 91 53 11 65 63 78 77 81 46 81 63 11 82 55 62
39 95 92 69 77 89 14 84 53 78 71 81 66 39 96 29
74 26 60 55 89 35 32 64 17 26 74 92 84 33 59 82
23 69 10 95 94 14 58 58 97 95 62 58 72 55 71 43
93 77 27 87 74 72 91 37 53 80 51 71 37 35 97 46
81 88 26 79 78 30 53 68 83 28 59 28 74 55 20 86
93 13 25 19 53 53 17 24 69 14 67 81 10 19 69 90
88 83 62 92 22 31 27 34 67 48 42 32 68 14 96 87
44 69 25 48 68 42 53 82 44 42 96 31 13 56 68 83
63 87 24 75 16 70 63 99 95 10 63 26 56 12 77 49
94 83 69 95 48 41 40 97 45 61 26 38 83 91 44 31
43 69 54 64 20 60 17 15 62 25 58 50 59 63 88 70
72 95 21 28 41 14 77 22 64 78 33 55 67 51 78 40
`
	// 1096  + 1098 + 1101
	expect := 3295
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4, 2
1 2 3 4
9 10 8 6
13 14 11 12
5 6 15 16
`
	// (9 + 16) * 8 / 2 = 100
	// 100 + 5 + 15
	expect := 121
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3, 1
1 2 3
4 5 6
7 8 9
`
	// (9 + 16) * 8 / 2 = 100
	// 100 + 5 + 15
	expect := 24
	runSample(t, s, expect)
}
