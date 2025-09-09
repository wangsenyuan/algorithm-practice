package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 3
4 5
2 4
`
	expect := []int{1, 2, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `9
10 20
50 60
30 40
70 80
90 100
60 70
10 40
40 50
80 90
`
	expect := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `30
78 80
37 39
66 68
47 55
39 43
48 56
44 52
7 9
27 34
34 36
90 96
13 22
76 79
9 16
75 83
1 11
74 76
56 62
61 65
77 83
48 53
48 50
22 29
39 46
70 72
24 26
1 9
40 45
13 22
46 49
`
	expect := []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 6, 7, 8, 8, 7, 7, 7, 7, 7, 7, 7, 7, 7, 6, 5, 6, 6, 6, 6, 6, 6}
	runSample(t, s, expect)
}
