package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `ahmed
3
ahmed posted on fatma's wall
fatma commented on ahmed's post
mona likes ahmed's post
`
	expect := []string{
		"fatma",
		"mona",
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `aba
1
likes likes posted's post
`
	expect := []string{
		"likes",
		"posted",
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `szfwtzfp
5
zqx posted on szfwtzfp's wall
r commented on scguem's post
r posted on civ's wall
r likes scguem's post
r likes scguem's post
`
	expect := []string{
		"zqx",
		"civ",
		"r",
		"scguem",
	}
	runSample(t, s, expect)
}
