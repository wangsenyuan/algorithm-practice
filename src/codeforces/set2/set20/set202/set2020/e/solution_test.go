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
	s := `2
1 2
5000 5000
`
	expect := 500000007
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
1 1
1000 2000
`
	expect := 820000006
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
343 624 675 451 902 820
6536 5326 7648 2165 9430 5428
`
	expect := 280120536
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := `1
1
10000
`
	expect := 1
	runSample(t, s, expect)
}
