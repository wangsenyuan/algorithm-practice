package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	words, res := drive(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	set := make(map[string]int)
	for i, w := range words {
		s := res[i]
		if len(s) == 0 || len(s) > 4 || !isAbbr(w, s) {
			t.Fatalf("Sample result %s, not an abbr for %s", s, w)
		}
		set[s]++
	}
	if len(set) != len(words) {
		t.Fatalf("Sample result %s, has duplicates", res)
	}
}

func isAbbr(w string, s string) bool {
	for i, j := 0, 0; i < len(s); i++ {
		for j < len(w) && s[i] != w[j] {
			j++
		}
		if j == len(w) {
			return false
		}
		j++
	}
	return true
}

func TestSample1(t *testing.T) {
	s := `6
privet
spasibo
codeforces
java
marmelad
normalno
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
aaa
aa
a
aaaa
aaaaa
`
	expect := false
	runSample(t, s, expect)
}
