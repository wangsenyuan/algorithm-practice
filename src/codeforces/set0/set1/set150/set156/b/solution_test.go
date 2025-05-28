package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	for _, x := range res {
		y := readString(reader)
		if x != y {
			t.Fatalf("Sample expect %s, but got %s", x, y)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `1 1
+1
Truth`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 2
-1
-2
-3
Not defined
Not defined
Not defined
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `4 1
+2
-3
+4
-1
Lie
Not defined
Lie
Not defined
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `4 3
-4
-3
-1
-3
Not defined
Truth
Not defined
Truth
`
	runSample(t, s)
}
