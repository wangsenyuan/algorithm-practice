package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
A
B
C
D
A-B 1:1
A-C 2:2
A-D 1:0
B-C 1:0
B-D 0:3
C-D 0:3
`, []string{"A", "D"})
}

func TestSample2(t *testing.T) {
	runSample(t, `2
a
A
a-A 2:1
`, []string{"a"})
}

func TestSample3(t *testing.T) {
	runSample(t, `4
TeMnHVvWKpwlpubwyhzqvc
AWJwc
bhbxErlydiwtoxy
EVASMeLpfqwjkke
AWJwc-TeMnHVvWKpwlpubwyhzqvc 37:34
bhbxErlydiwtoxy-TeMnHVvWKpwlpubwyhzqvc 38:99
bhbxErlydiwtoxy-AWJwc 33:84
EVASMeLpfqwjkke-TeMnHVvWKpwlpubwyhzqvc 79:34
EVASMeLpfqwjkke-AWJwc 24:37
EVASMeLpfqwjkke-bhbxErlydiwtoxy 3:6
`, []string{"AWJwc",
		"EVASMeLpfqwjkke"})
}
