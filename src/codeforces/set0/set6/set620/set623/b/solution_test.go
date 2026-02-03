package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 1 4
4 2 3
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 3 2
5 17 13 5 6
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8 3 4
3 7 5 4 3 12 9 4
`
	expect := 13
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 4 3
9 9 4
`
	expect := 3
	runSample(t, s, expect)
}

// too slow to test
// func TestSample5(t *testing.T) {
// 	s := `45 393296667 817784089
// 513035443 513035443 513035445 513035445 513035443 555648978 463549879 830693049 524420119 541365334 562791911 736173182 644390537 639730339 913702156 807631127 679411095 457747249 554394051 803981524 901463184 651788488 792766018 587562656 564169971 645381787 497940709 886010956 577261234 513035444 513035443 513035443 513035445 513035445 513035444 513035444 513035443 513035443 513035445 513035445 513035443 513035443 513035443 513035443 513035445
// `
// 	expect := 16911756681
// 	runSample(t, s, expect)
// }
