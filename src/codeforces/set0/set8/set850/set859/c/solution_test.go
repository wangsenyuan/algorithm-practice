package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
141 592 653
`
	epxect := []int{653, 733}
	runSample(t, s, epxect)
}

func TestSample2(t *testing.T) {
	s := `5
10 21 10 21 10
`
	epxect := []int{31, 41}
	runSample(t, s, epxect)
}
