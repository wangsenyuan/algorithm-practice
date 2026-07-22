package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `10 10
+ 6
+ 10
+ 5
- 10
- 5
- 6
+ 10
+ 3
+ 6
+ 3
`, []string{
		"Success",
		"Conflict with 6",
		"Success",
		"Already off",
		"Success",
		"Success",
		"Success",
		"Success",
		"Conflict with 10",
		"Already on",
	})
}
