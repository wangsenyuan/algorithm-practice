package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if math.IsNaN(res) || math.Abs(res-expect)/max(1, expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `0 0 4
6 0 4
`
	expect := 7.25298806364175601379
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `0 0 5
11 0 5
`
	expect := 0.0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6008 8591 6693
5310 8351 7192
`
	expect := 138921450.46886559338599909097
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `-9 8 7
-9 8 5
`
	expect := 78.53981633974482789995
	runSample(t, s, expect)
}

// func TestSample5(t *testing.T) {
// 	s := `44721 999999999 400000000
// 0 0 600000000
// `
// 	expect := 0.00188343226909637451
// 	runSample(t, s, expect)
// }
