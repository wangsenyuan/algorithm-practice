package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	drive(reader, writer)
	res := buf.String()
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}
func TestSample1(t *testing.T) {
	s := `8
set chemlabreport 1
set physicsexercise 2
set chinesemockexam 3
query physicsexercise
query chinesemockexam
remove physicsexercise
query physicsexercise
query chinesemockexam
`
	expect := "1\n2\n-1\n1\n"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8
set physicsexercise 2
set chinesemockexam 3
set physicsexercise 1
query physicsexercise
query chinesemockexam
undo 4
query physicsexercise
query chinesemockexam
`
	expect := "0\n1\n0\n-1\n"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
query economicsessay
remove economicsessay
query economicsessay
undo 2
query economicsessay
`
	expect := "-1\n-1\n-1\n"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
set economicsessay 1
remove economicsessay
undo 1
undo 1
query economicsessay
`
	expect := "-1\n"
	runSample(t, s, expect)
}
