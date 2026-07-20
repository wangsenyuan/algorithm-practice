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
	if res == nil {
		res = []string{}
	}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `400 500 600 700 800
`, []string{
		"ABCDE", "BCDE", "ACDE", "ABDE", "ABCE", "ABCD",
		"CDE", "BDE", "ADE", "BCE", "ACE", "BCD", "ABE", "ACD", "ABD", "ABC",
		"DE", "CE", "BE", "CD", "AE", "BD", "AD", "BC", "AC", "AB",
		"E", "D", "C", "B", "A",
	})
}

func TestSample2(t *testing.T) {
	runSample(t, `800 800 900 900 1000
`, []string{
		"ABCDE", "ACDE", "BCDE", "ABCE", "ABDE", "ABCD",
		"CDE", "ACE", "ADE", "BCE", "BDE", "ABE", "ACD", "BCD", "ABC", "ABD",
		"CE", "DE", "AE", "BE", "CD", "AC", "AD", "BC", "BD", "AB",
		"E", "C", "D", "A", "B",
	})
}

func TestSample3(t *testing.T) {
	runSample(t, `128 256 512 1024 2048
`, []string{
		"ABCDE", "BCDE", "ACDE", "CDE", "ABDE", "BDE", "ADE", "DE",
		"ABCE", "BCE", "ACE", "CE", "ABE", "BE", "AE", "E",
		"ABCD", "BCD", "ACD", "CD", "ABD", "BD", "AD", "D",
		"ABC", "BC", "AC", "C", "AB", "B", "A",
	})
}
