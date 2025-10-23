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
	s := `5 2
1 2
2 3
	`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `100 3
1 2
2 1
3 1
	`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 2
1 1
2 100
	`
	expect := 100
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `73 19
21018 52113
53170 12041
44686 99498
73991 59354
66652 2045
56336 99193
85265 20504
51776 85293
21550 17562
70468 38130
7814 88602
84216 64214
69825 55393
90671 24028
98076 67499
46288 36605
17222 21707
25011 99490
92165 51620
	`
	expect := 860399
	runSample(t, s, expect)
}
