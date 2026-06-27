package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 4
3 -1 -4 1
-5 9 -2 -6
-5 3 -5 8
`, 23)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 5
-70 11 -45 -54 -30
-99 39 -83 -69 -77
-48 -21 -43 -96 -24
-54 -65 21 -88 -44
-90 -33 -67 -29 -62
`, 39)
}

func TestSample3(t *testing.T) {
	runSample(t, `8 9
-74832 16944 58683 32965 97236 -52995 43262 -51959 40883
-58715 13846 24919 65627 -11492 -63264 29966 -98452 -75577
40415 77202 15542 -50602 83295 85415 -35304 46520 -38742
37482 56721 -38521 63127 55608 95115 42893 10484 70510
53019 40623 25885 -10246 70973 32528 -33423 19322 52097
79880 74931 -58277 -33783 91022 -53003 11085 -65924 -63548
78622 -77307 81181 46875 -81091 63881 11160 -82217 -55492
62770 39530 -95923 92440 -69899 77737 89392 -14281 84899
`, 2232232)
}

func TestSingleWideRow(t *testing.T) {
	runSample(t, `1 21
-1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1
`, -1)
}

func TestWideSquare(t *testing.T) {
	a := make([][]int64, 21)
	for i := range a {
		a[i] = make([]int64, 21)
		for j := range a[i] {
			a[i][j] = -1
		}
	}
	if res := solve(a); res != -1 {
		t.Fatalf("expect -1, but got %d", res)
	}
}
