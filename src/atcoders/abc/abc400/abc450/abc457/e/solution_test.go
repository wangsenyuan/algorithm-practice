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
	s := `4 3
1 3
1 1
2 4
4
1 4
2 4
1 3
1 1
`
	expect := []string{"Yes", "No", "Yes", "No"}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 10
2 6
2 5
3 6
1 6
1 2
5 6
2 3
3 7
2 3
1 2
10
1 2
3 5
1 4
1 5
1 5
5 7
1 6
2 3
5 7
2 4
`
	expect := []string{
		"Yes", "No", "No", "Yes", "Yes",
		"No", "Yes", "Yes", "No", "No",
	}
	runSample(t, s, expect)
}

func TestDoesNotReuseQueryAsCloth(t *testing.T) {
	s := `3 2
2 2
1 3
3
1 1
1 2
1 3
`
	expect := []string{"No", "No", "Yes"}
	runSample(t, s, expect)
}
