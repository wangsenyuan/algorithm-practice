package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expectPainted int) {
	t.Helper()
	k, a, res := drive(bufio.NewReader(strings.NewReader(s)))

	colr := make(map[int]map[int]bool)

	var cnt int
	freq := make(map[int]int)
	for i, v := range res {
		if v > 0 {
			freq[v]++
			cnt++

			w := a[i]
			if _, ok := colr[w]; !ok {
				colr[w] = make(map[int]bool)
			}
			if colr[w][v] {
				t.Fatalf("value %d painted twice with color %d", w, v)
			}
			colr[w][v] = true
		}
		if v > k {
			t.Fatalf("can only use color <= %d, but use %d", k, v)
		}
	}
	if cnt != expectPainted {
		t.Fatalf("Sample result %v, not correct", res)
	}

	if len(freq) != k {
		t.Fatalf("Sample result %v, is not valid", res)
	}

}

func TestSample1(t *testing.T) {
	runSample(t, `10 3
3 1 1 1 1 10 3 10 10 2
`, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 4
1 1 1 1
`, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, `1 1
1
`, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, `13 1
3 1 4 1 5 9 2 6 5 3 5 8 9
`, 8)
}

func TestSample5(t *testing.T) {
	runSample(t, `13 2
3 1 4 1 5 9 2 6 5 3 5 8 9
`, 12)
}

func TestSample6(t *testing.T) {
	runSample(t, `13 3
3 1 4 1 5 9 2 6 5 3 5 8 9
`, 12)
}

func TestSample7(t *testing.T) {
	runSample(t, `4 3
2 1 2 1
`, 3)
}
