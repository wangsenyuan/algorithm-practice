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
		if y != x {
			t.Fatalf("Sample expect %s, but got %v", s, res)
		}
	}

}

func TestSample1(t *testing.T) {
	s := `3
8 4 1
2
2 3
1 2
5
12`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `6
1 2 4 8 16 32
4
1 6
2 5
3 4
1 2
60
30
12
3`
	runSample(t, s)
}
