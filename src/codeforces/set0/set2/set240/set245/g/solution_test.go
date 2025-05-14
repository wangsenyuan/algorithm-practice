package main

import (
	"bufio"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)

	sort.Strings(expect)
	sort.Strings(res)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
Mike Gerald
Kate Mike
Kate Tank
Gerald Tank
Gerald David
`
	expect := []string{
		"Mike 1",
		"Gerald 1",
		"Kate 1",
		"Tank 1",
		"David 2",
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
valera vanya
valera edik
pasha valera
igor valera
`
	expect := []string{
		"valera 0",
		"vanya 3",
		"edik 3",
		"pasha 3",
		"igor 3",
	}
	runSample(t, s, expect)
}
