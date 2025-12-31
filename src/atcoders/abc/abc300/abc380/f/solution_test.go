package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1 1 2
2
4
1 3
`
	expect := "Aoki"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 4 4
98 98765 987654 987654321
987 9876 9876543 98765432
123 12345 1234567 123456789
`
	expect := "Takahashi"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 1 8
10
10
1 2 3 4 5 6 7 8
`
	expect := "Aoki"
	runSample(t, s, expect)
}
