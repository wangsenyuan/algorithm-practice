package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string) {
	t.Helper()

	s, queries, res := drive(bufio.NewReader(strings.NewReader(input)))

	pref := make([]int, len(s)+1)
	for i := range s {
		if s[i] == 'T' {
			pref[i+1] = pref[i] + 2
		} else {
			pref[i+1] = pref[i] + 1
		}
	}

	checkExists := func(x int) bool {
		for i := range len(s) {
			for j := i; j < len(s); j++ {
				if pref[j+1]-pref[i] == x {
					return true
				}
			}
		}
		return false
	}

	for i, x := range queries {
		if res[i] == nil {
			if checkExists(x) {
				t.Fatalf("query %d: has no solution", x)
			}
			continue
		}
		l, r := res[i][0], res[i][1]
		if pref[r]-pref[l-1] != x {
			t.Fatalf("query %d: %d != %d", x, pref[r]-pref[l-1], x)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 3
TWTWT
5
1
7
`)
}
