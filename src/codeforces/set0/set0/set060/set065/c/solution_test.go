package main

import (
	"bufio"
	"fmt"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ok, ans, p0 := process(reader)
	flag := readString(reader)
	if flag == "YES" != ok {
		t.Fatalf("Sample expect %s but got %t", flag, ok)
	}
	if !ok {
		return
	}
	var expect_ans, x, y, z float64
	fmt.Fscanf(reader, "%f\n%f %f %f", &expect_ans, &x, &y, &z)

	if math.Abs(expect_ans-ans) > 1e-6 {
		t.Fatalf("Sample expect %f but got %f", expect_ans, ans)
	}

	p1 := []float64{x, y, z}
	for i := range 3 {
		if math.Abs(p1[i]-p0[i]) > 1e-6 {
			t.Fatalf("Sample expect %v but got %v", p1, p0)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4
0 0 0
0 10 0
10 10 0
10 0 0
0 0 0
1 1
5 5 25
YES
25.5000000000
10.0000000000 4.5000000000 0.0000000000
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4
0 0 0
0 10 0
10 10 0
10 0 0
0 0 0
1 1
5 5 50
NO
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `1
1 2 3
4 5 6
20 10
1 2 3
YES
0.0000000000
1.0000000000 2.0000000000 3.0000000000
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `1
1 2 3
4 5 6
20 10
1 2 3
YES
0.0000000000
1.0000000000 2.0000000000 3.0000000000
`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `1
0 0 0
0 0 1
10000 10000
0 0 1
YES
0.0000500000
0.0000000000 0.0000000000 0.5000000000
`
	runSample(t, s)
}
