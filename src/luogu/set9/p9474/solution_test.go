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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 3
1 2 3 4 5`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 3
1 7 8 3 4 6`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `100 43
11451 28255 11021 1888 13765 12592 30989 18758 7833 21591 15085 13547 11805 31668 23385 18266 30204 6101 22525 22939 13550 20258 21998 29574 11834 1879 21829 16600 6777 9016 18445 23687 5532 18560 4191 26195 11824 16922 10699 5790 31201 21139 506 17533 21309 2768 17554 4623 9403 30972 7770 13070 25852 17349 15263 419 10320 1480 1494 21887 27516 32073 25730 20775 31675 25640 1368 6543 8194 1040 28577 11243 26504 23282 26337 17858 16774 25522 24758 4932 16955 22304 4610 6109 14644 31017 22348 1290 31946 17999 31844 29548 3820 5462 10054 12827 32543 320 27201 8252`
	expect := 23219
	runSample(t, s, expect)
}
