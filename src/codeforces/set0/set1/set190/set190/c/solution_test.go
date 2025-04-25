package main

import (
	"bytes"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "pair pair int int int"
	expect := "pair<pair<int,int>,int>"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "pair int"
	expect := "Error occurred"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {

	var buf bytes.Buffer
	for range 10000 {
		buf.WriteString("pair<int,")
	}
	buf.WriteString("int")
	for range 10000 {
		buf.WriteByte('>')
	}

	expect := buf.String()
	s := expect
	s = strings.ReplaceAll(s, "<", " ")
	s = strings.ReplaceAll(s, ">", " ")
	s = strings.ReplaceAll(s, ",", " ")
	runSample(t, s, expect)
}
