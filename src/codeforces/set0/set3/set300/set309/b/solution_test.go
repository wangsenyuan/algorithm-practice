package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, r, c, s := process(reader)

	if len(res) > r {
		t.Fatalf("sample result %v, not correct", res)
	}

	var ans bytes.Buffer
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		ans.WriteString(strings.TrimSpace(s))
		ans.WriteString(" ")
	}

	cnt0 := len(strings.Split(ans.String(), " "))

	var buf bytes.Buffer
	for _, x := range res {
		xx := strings.Split(x, " ")
		if len(xx) > c {
			t.Fatalf("sample result %v, not correct", res)
		}
		buf.WriteString(x)
		buf.WriteString(" ")
	}

	cnt1 := len(strings.Split(buf.String(), " "))

	if cnt0 != cnt1 {
		t.Fatalf("sample result %v, not correct", res)
	}
	w := buf.String()
	w = strings.TrimSpace(w)

	if !strings.Contains(s, w) {
		t.Fatalf("sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `9 4 12
this is a sample text for croc final round
this is a
sample text
for croc
final round
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `9 1 9
this is a sample text for croc final round
this is a
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `6 2 3
croc a a a croc a
a a
a
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `2 2 5
first second
first
`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `10 12 7
l ry ksgx bxeb t w szsw m bf eyfc
l ry
ksgx
bxeb t
w szsw
m bf
eyfc
`
	runSample(t, s)
}

func TestSample6(t *testing.T) {
	s := `1 1 1
abc
`
	runSample(t, s)
}
