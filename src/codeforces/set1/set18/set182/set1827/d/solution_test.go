package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	for _, x := range res {
		var y int
		fmt.Fscan(reader, &y)
		if x != y {
			t.Fatalf("Sample expect %v, but got %v", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2
1
0`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3
1 1
0 1`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `4
1 2 3
0 1 0`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `7
1 2 3 2 5 2
0 1 0 1 2 3`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `10
1 2 2 4 5 5 7 8 9
0 1 2 1 0 1 0 1 2`
	runSample(t, s)
}

func TestSample6(t *testing.T) {
	s := `4
1 2 1
0 1 0`
	runSample(t, s)
}
