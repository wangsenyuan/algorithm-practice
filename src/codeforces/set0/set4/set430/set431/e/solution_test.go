package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	for i, v := range expect {
		w := res[i]
		if math.Abs(v-w)/max(1, v) > 1e-4 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
1 2 0
2 2
1 2 1
2 3
`
	expect := []float64{1.5, 1.6666667}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 5
1 3 0 1
2 3
2 1
1 3 2
2 3
2 4
`
	expect := []float64{1.66667, 1.00000, 2.33333, 2.66667}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 5
47 90 84 21 59 25 58 47 57 76
1 6 25
2 3517
2 2432
2 2325
2 2647
`
	expect := []float64{408.10000, 299.60000, 288.90000, 321.10000}

	runSample(t, s, expect)
}
