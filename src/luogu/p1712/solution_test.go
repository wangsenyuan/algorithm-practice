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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6 3
3 5
1 2
3 4
2 2
1 5
1 4
2`)
}

func TestSample2(t *testing.T) {
	runSample(t, `20 9
0 100
43 72
24 32
79 80
52 82
57 74
31 58
3 12
29 50
54 57
9 87
36 38
2 15
49 85
65 81
59 97
81 98
47 98
4 75
66 89
55`)
}