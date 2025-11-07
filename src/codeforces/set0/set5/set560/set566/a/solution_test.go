package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect_score int) {
	reader := bufio.NewReader(strings.NewReader(s))
	names, pseudonym, score, res := drive(reader)
	if score != int32(expect_score) {
		t.Fatalf("Sample expect %d, but got %d", expect_score, score)
	}

	lcp := func(a, b string) int {
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] != b[i] {
				return i
			}
		}
		return min(len(a), len(b))
	}

	var sum int
	for _, cur := range res {
		i, j := cur[0]-1, cur[1]-1
		sum += lcp(names[i], pseudonym[j])
	}

	if sum != expect_score {
		t.Fatalf("Sample expect %d, but got %d", expect_score, sum)
	}
}

func TestSample1(t *testing.T) {
	s := `5
gennady
galya
boris
bill
toshik
bilbo
torin
gendalf
smaug
galadriel
`
	expect := 11
	runSample(t, s, expect)
}
