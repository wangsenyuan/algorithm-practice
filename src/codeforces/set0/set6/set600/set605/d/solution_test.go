package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans, cards := process(reader)
	expect_len := readNum(reader)
	if len(ans) == 0 && expect_len == -1 {
		return
	}
	if len(ans) == 0 || expect_len == -1 || len(ans) != expect_len {
		t.Fatalf("Sample expect %d, but got %v", expect_len, ans)
	}

	var x, y int

	for _, i := range ans {
		i--
		card := cards[i]
		if card[0] > x || card[1] > y {
			t.Fatalf("Sample result not correct, at %d", i)
		}
		x, y = card[2], card[3]
	}

	if ans[len(ans)-1] != len(cards) {
		t.Fatalf("Sample result not correct, at %d", len(cards))
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
0 0 3 4
2 2 5 3
4 1 1 7
5 3 8 8
3
1 2 4
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
0 0 4 6
5 1 1000000000 1000000000
-1
	`)
}
