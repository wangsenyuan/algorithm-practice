package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %q, but got %q", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1 10
10 1
9 9
9 9
`, "Team 1")
}

func TestSample2(t *testing.T) {
	runSample(t, `1 1
2 2
3 3
2 2
`, "Team 2")
}

func TestSample3(t *testing.T) {
	runSample(t, `3 3
2 2
1 1
2 2
`, "Draw")
}

func TestTeam2MustHaveResponseToBothTeam1Assignments(t *testing.T) {
	runSample(t, `1 1
2 3
3 2
2 2
`, "Draw")
}
