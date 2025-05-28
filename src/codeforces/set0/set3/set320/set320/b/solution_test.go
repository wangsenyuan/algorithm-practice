package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	for _, x := range res {
		expect := readString(reader)
		if expect == "YES" != x {
			t.Errorf("Sample expect %s, but got %v", expect, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 1 5
1 5 11
2 1 2
1 2 9
2 1 2
NO
YES`
	runSample(t, s)
}


func TestSample2(t *testing.T) {
	s := `10
1 -311 -186
1 -1070 -341
1 -1506 -634
1 688 1698
2 2 4
1 70 1908
2 1 2
2 2 4
1 -1053 1327
2 5 4
NO
NO
NO
YES`
	runSample(t, s)
}
