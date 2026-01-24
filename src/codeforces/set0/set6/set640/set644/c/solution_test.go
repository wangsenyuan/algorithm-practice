package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if len(res) != len(expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}

	normalize := func(arr []string) []string {
		res := slices.Clone(arr)
		for i, s := range res {
			ss := strings.Split(s, " ")
			slices.Sort(ss)
			res[i] = strings.Join(ss, " ")
		}
		slices.Sort(res)
		return res
	}

	a := normalize(res)
	b := normalize(expect)

	if !slices.Equal(a, b) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `10
http://abacaba.ru/test
http://abacaba.ru/
http://abacaba.com
http://abacaba.com/test
http://abacaba.de/
http://abacaba.ru/test
http://abacaba.de/test
http://abacaba.com/
http://abacaba.com/t
http://abacaba.com/test
`
	expect := []string{
		"http://abacaba.de http://abacaba.ru",
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `14
http://c
http://ccc.bbbb/aba..b
http://cba.com
http://a.c/aba..b/a
http://abc/
http://a.c/
http://ccc.bbbb
http://ab.ac.bc.aa/
http://a.a.a/
http://ccc.bbbb/
http://cba.com/
http://cba.com/aba..b
http://a.a.a/aba..b/a
http://abc/aba..b/a
`
	expect := []string{
		"http://cba.com http://ccc.bbbb",
		"http://a.a.a http://a.c http://abc",
	}
	runSample(t, s, expect)
}
