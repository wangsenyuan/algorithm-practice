package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	f, b, ans, a := drive(reader)
	if ans != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, ans)
	}
	if ans == "Possible" {
		n := len(f)
		for i := range n {
			if b[i] != f[a[i]-1] {
				t.Fatalf("Sample result %v is not correct, expect %v", a, b)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3
3 2 1
1 2 3`, "Possible")
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3
1 1 1
1 1 1`, "Ambiguity")
}

func TestSample3(t *testing.T) {
	runSample(t, `3 3
1 2 1
3 3 3`, "Impossible")
}

func TestSample4(t *testing.T) {
	runSample(t, `2 10
2 1
1 1 1 1 1 1 2 2 2 2 2`, "Possible")
}
