package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	words, res := drive(reader)

	if res == expect {
		return
	}
	if expect == "-1" || res == "-1" {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	check := func(t string) bool {
		for _, s := range words {
			var cnt int
			for j := range len(s) {
				if t[j] != s[j] {
					cnt++
				}
			}
			if cnt == 1 || cnt > 2 {
				return false
			}
		}
		return true
	}

	if !check(res) {
		t.Fatalf("Sample result %s, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
abac
caab
acba
`
	expect := "acab"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4
kbbu
kbub
ubkb
`
	expect := "kbub"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 4
abcd
dcba
acbd
dbca
zzzz
`
	expect := "-1"
	runSample(t, s, expect)
}
