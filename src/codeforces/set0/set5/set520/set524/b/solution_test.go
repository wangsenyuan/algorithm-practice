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
	s := `3 
10 1 
20 2 
30 3
`
	expect := 180
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 
3 1 
2 2 
4 3
`
	expect := 21
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 
5 10
`
	expect := 50
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := `10
168 538
836 439
190 873
206 47
891 591
939 481
399 898
859 466
701 777
629 222
`
	expect := 3478056
	runSample(t, s, expect)
}
