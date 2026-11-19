package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s + "\n")))
	if !optimal(s, res) {
		t.Fatalf("Sample %q got %q, which is not optimal", s, res)
	}
}

func optimal(s string, ans string) bool {
	if len(ans) != len(s) {
		return false
	}
	cnt := map[byte]int{'R': 0, 'S': 0, 'P': 0}
	for i := range s {
		cnt[s[i]]++
	}
	best := cnt['R']
	if cnt['S'] > best {
		best = cnt['S']
	}
	if cnt['P'] > best {
		best = cnt['P']
	}
	beats := map[byte]byte{'R': 'S', 'S': 'P', 'P': 'R'}
	total := 0
	for i := range ans {
		loser, ok := beats[ans[i]]
		if !ok {
			return false
		}
		total += cnt[loser]
	}
	return total == len(s)*best
}

func TestSample1(t *testing.T) {
	runSample(t, "RRRR")
}

func TestSample2(t *testing.T) {
	runSample(t, "RSP")
}

func TestSample3(t *testing.T) {
	runSample(t, "S")
}

func TestSample4(t *testing.T) {
	runSample(t, "RRRSS")
}
