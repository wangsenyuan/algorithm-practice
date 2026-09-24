package main

import (
	"bufio"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func readInts(t *testing.T, path string) []int {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	fields := strings.Fields(string(data))
	res := make([]int, len(fields))
	for i, f := range fields {
		res[i], err = strconv.Atoi(f)
		if err != nil {
			t.Fatal(err)
		}
	}
	return res
}

func TestSample1(t *testing.T) {
	runSample(t, `2 5
1 1 2 1
1 1 2 2
2 1 1 2
2 1 1 1
2 1 2 3
`, []int{1, 2, 1})
}

func TestSample2(t *testing.T) {
	in, err := os.ReadFile("P3332_2.in")
	if err != nil {
		t.Fatal(err)
	}
	runSample(t, string(in), readInts(t, "P3332_2.out"))
}

func TestKthLargestOnFullRange(t *testing.T) {
	// union of {1,2} on three positions: 2,2,2,1,1,1; 3rd largest is 2
	runSample(t, `3 3
1 1 3 1
1 1 3 2
2 1 3 3
`, []int{2})
}

func TestQueryRightmostPosition(t *testing.T) {
	runSample(t, `3 3
1 3 3 1
1 3 3 2
2 3 3 1
`, []int{2})
}
