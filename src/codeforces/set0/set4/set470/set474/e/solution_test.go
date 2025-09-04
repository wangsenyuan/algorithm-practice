package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	d, h, res := drive(reader)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	for i := 0; i+1 < len(res); i++ {
		u := res[i] - 1
		v := res[i+1] - 1
		if u > v || abs(h[u]-h[v]) < d {
			t.Fatalf("Sample result %v is not correct", res)
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func TestSample1(t *testing.T) {
	s := `5 2
1 3 6 7 4`
	runSample(t, s, 4)
}

func TestSample2(t *testing.T) {
	s := `10 3
2 1 3 6 9 11 7 3 20 18`
	runSample(t, s, 6)
}

func TestSample3(t *testing.T) {
	s := `87 9251
44243 43803 31356 25085 37429 40103 16225 16828 11148 1052 39431 20858 7018 14692 28388 32588 26498 25148 29777 15299 41387 11500 11246 10537 23289 40449 7780 2521 34574 35306 42224 22506 3036 40157 16 3131 29509 8851 26481 13779 19050 40537 7494 27005 7880 45546 30797 7650 10785 21193 22252 43185 10451 44252 31223 3421 19196 33816 39279 19663 6546 1761 24753 1344 866 46085 13446 42731 44 45374 19905 16142 22155 19184 20464 31176 28168 19368 17906 38706 44447 29519 16841 27338 35425 9257 23284`
	runSample(t, s, 57)
}
