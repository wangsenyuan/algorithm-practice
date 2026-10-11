package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, sample string) {
	t.Helper()
	s, k, res := drive(bufio.NewReader(strings.NewReader(s)))
	if !validMask(res, len(s), k) {
		t.Fatalf("invalid mask %q for k=%d s=%s", res, k, s)
	}
	got, want := cost(apply(s, res)), cost(apply(s, sample))
	if got != want {
		t.Fatalf("mask %q has cost %d, sample %q has cost %d", res, got, sample, want)
	}
}

func parseCase(s string) (int, string) {
	reader := bufio.NewReader(strings.NewReader(s))
	var n, k int
	var bracket string
	fmt.Fscan(reader, &n, &k, &bracket)
	return k, bracket
}

func validMask(mask string, n, k int) bool {
	if len(mask) != n {
		return false
	}
	ones := 0
	for _, ch := range mask {
		if ch != '0' && ch != '1' {
			return false
		}
		if ch == '1' {
			ones++
		}
	}
	return ones <= k
}

func apply(s, mask string) string {
	buf := make([]byte, 0, len(s))
	for i := range s {
		if mask[i] == '0' {
			buf = append(buf, s[i])
		}
	}
	return string(buf)
}

func cost(s string) int {
	balance, pairs := 0, 0
	for _, ch := range s {
		if ch == '(' {
			balance++
			continue
		}
		if balance > 0 {
			balance--
			pairs++
		}
	}
	return pairs * 2
}

func TestSample1(t *testing.T) {
	runSample(t, `2 1
)(
`, "00")
}

func TestSample2(t *testing.T) {
	runSample(t, `2 0
()
`, "00")
}

func TestSample3(t *testing.T) {
	runSample(t, `4 1
(())
`, "1000")
}

func TestSample4(t *testing.T) {
	runSample(t, `4 1
())(
`, "1000")
}

func TestSample5(t *testing.T) {
	runSample(t, `5 1
((())
`, "00010")
}

func TestSample6(t *testing.T) {
	runSample(t, `6 2
()()()
`, "101000")
}

func TestSample7(t *testing.T) {
	runSample(t, `6 2
(()())
`, "001001")
}

func TestSample8(t *testing.T) {
	runSample(t, `6 2
())(()
`, "100001")
}

func TestSample9(t *testing.T) {
	runSample(t, `7 3
(()((()
`, "1100001")
}

func TestSample10(t *testing.T) {
	runSample(t, `10 3
(()())())(
`, "0101001000")
}
