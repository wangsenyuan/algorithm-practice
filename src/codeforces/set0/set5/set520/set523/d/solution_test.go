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
			t.Fatalf("Sample expect %d, but got %d", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 2
1 5
2 5
3 5
6
7
11`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `6 1
1 1000000000
2 1000000000
3 1000000000
4 1000000000
5 1000000000
6 3
1000000001
2000000001
3000000001
4000000001
5000000001
5000000004`
	runSample(t, s)
}
