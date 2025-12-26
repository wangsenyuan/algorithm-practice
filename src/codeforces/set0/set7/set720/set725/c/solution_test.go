package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	x, res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	check := func(r int, c int) bool {
	top:
		for i := 0; i+1 < len(x); i++ {
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					if dr == 0 && dc == 0 {
						continue
					}
					nr, nc := r+dr, c+dc
					if nr >= 0 && nr <= 1 && nc >= 0 && nc < 13 && res[nr][nc] == x[i+1] {
						r, c = nr, nc
						continue top
					}
				}
			}
			return false
		}
		return true
	}

	for r, row := range res {
		for c := range len(row) {
			if row[c] == x[0] && check(r, c) {
				return
			}
		}
	}
	t.Fatalf("Sample result %v, not correct, check failed", res)
}

func TestSample1(t *testing.T) {
	s := "ABCDEFGHIJKLMNOPQRSGTUVWXYZ"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "BUVTYZFQSNRIWOXXGJLKACPEMDH"
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "UTEDBZRVWLOFUASHCYIPXGJMKNQ"
	expect := true
	runSample(t, s, expect)
}