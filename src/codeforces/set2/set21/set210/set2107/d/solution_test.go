package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if len(ans) == 0 {
		t.Fatalf("Sample expect answer, but got nothing")
	}
	x := fmt.Sprintf("%v", ans)
	x = x[1 : len(x)-1]
	x = strings.ReplaceAll(x, ",", " ")
	y := readString(reader)
	if x != y {
		t.Fatalf("Sample expect %s, but got %s", y, x)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 2
1 3
1 4
3 4 3 1 2 2
	`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4
2 1
2 4
2 3
3 4 3 1 1 1
	`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `5
1 2
2 3
3 4
4 5
5 5 1
	`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `1
1 1 1
	`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `8
6 3
3 5
5 4
4 2
5 1
1 8
3 7
5 8 7 2 4 2 1 6 6
	`
	runSample(t, s)
}

func TestSample6(t *testing.T) {
	s := `6
3 2
2 6
2 5
5 4
4 1
5 6 1 1 3 3
	`
	runSample(t, s)
}
