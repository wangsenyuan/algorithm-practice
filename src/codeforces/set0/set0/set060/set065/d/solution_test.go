package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, expect []string) {
	res := solve(s)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "G????SS???H"
	expect := []string{"Gryffindor", "Ravenclaw"}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "H?"
	expect := []string{"Gryffindor", "Ravenclaw", "Slytherin"}
	runSample(t, s, expect)
}
