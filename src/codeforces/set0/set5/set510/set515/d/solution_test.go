package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if expect == "Not unique" {
		if len(res) > 0 {
			t.Fatalf("Sample expect %s, but got %v", expect, res)
		}
		return
	}
	if len(res) == 0 {
		t.Fatalf("Sample expect %s, but got %v", expect, res)
	}

	ans := strings.Join(res, "\n")

	if ans != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
...
.*.
...
`
	expect := "Not unique"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 4
..**
*...
*.**
....
`
	expect := `<>**
*^<>
*v**
<><>`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 4
*..*
....
`
	expect := `*<>*
<><>`
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `1 1
.
`
	expect := "Not unique"

	runSample(t, s, expect)
}
