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
		t.Fatalf("expect %q, got %q for input:\n%s", expect, res, s)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
3 3 3 3 3`, Alice)
}

func TestStatementSamples(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect string
	}{
		{"one_and_seven", `2
1 7`, Bob},
		{"seven_distinct", `7
1 3 9 7 4 2 100`, Alice},
		{"one_two_three", `3
1 2 3`, Alice},
		{"six_with_dupes", `6
2 1 3 4 2 4`, Bob},
		{"eight_mixed", `8
5 7 2 9 6 3 3 2`, Alice},
		{"single_large", `1
1000000000`, Alice},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			runSample(t, tc.input, tc.expect)
		})
	}
}

func TestEdgeCases(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect string
	}{
		{"single_one", `1
1`, Alice},
		{"two_consecutive", `2
1 2`, Bob},
		{"two_equal", `2
5 5`, Alice},
		{"two_gap_two", `2
2 4`, Alice},
		{"three_all_same", `3
10 10 10`, Alice},
		{"two_one_one", `2
1 1`, Alice},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			runSample(t, tc.input, tc.expect)
		})
	}
}
