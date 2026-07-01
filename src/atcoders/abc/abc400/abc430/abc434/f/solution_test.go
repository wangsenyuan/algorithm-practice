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
		t.Fatalf("Sample expect %q, but got %q", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
abc
ac
ahc
`, "abcahcac")
}

func TestSample2(t *testing.T) {
	runSample(t, `4
aaa
a
aaaa
a
`, "aaaaaaaaa")
}

func TestSample3(t *testing.T) {
	runSample(t, `15
ks
sy
k
ysk
yks
ky
ksy
sk
syk
s
kys
sky
ys
yk
y
`, "kksksykykysskskyssyksyyksykyskyys")
}

func TestSecondCandidateCanSwapBeforeLastPair(t *testing.T) {
	runSample(t, `3
a
aab
b
`, "aabab")
}

func TestTwoStrings(t *testing.T) {
	runSample(t, `2
a
b
`, "ba")
}
