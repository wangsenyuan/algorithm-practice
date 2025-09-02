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
	s := `2 1
2 1
`
	expect := []int{2, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 2
1 2
1 1
`
	expect := []int{2, 1, 3}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3
1 3
2 3
1 3
`
	runSample(t, s, nil)
}

func TestSample4(t *testing.T) {
	s := `100 50
11 28
11 1
98 58
38 27
24 27
67 37
90 48
91 14
43 29
3 64
24 6
53 19
97 65
13 27
75 53
37 82
69 75
94 99
1 26
95 60
45 27
100 82
71 49
86 99
74 58
88 68
39 63
38 23
22 39
29 58
62 83
62 1
61 58
2 30
41 48
83 90
1 17
73 81
23 53
71 16
43 29
27 78
54 48
6 89
75 27
16 93
81 81
97 31
53 32
15 96
`

	// 在这个例子中， 43的其实是在24和38的后面的，但是没有被体现出来

	expect := []int{2, 4, 5, 7, 8, 9, 10, 91, 12, 45, 53, 1, 14, 17, 18, 19, 20, 13, 22, 54, 21, 25, 43, 24, 38, 26, 28, 11, 41, 30, 31, 23, 32, 33, 34, 67, 35, 36, 71, 40, 42, 61, 44, 29, 46, 47, 90, 74, 48, 75, 49, 50, 39, 95, 51, 52, 55, 98, 56, 57, 88, 58, 59, 3, 97, 60, 63, 64, 65, 27, 81, 66, 68, 69, 73, 70, 72, 76, 62, 100, 77, 37, 78, 79, 80, 6, 82, 83, 84, 85, 16, 87, 89, 15, 92, 93, 96, 86, 94, 99}
	runSample(t, s, expect)
}
