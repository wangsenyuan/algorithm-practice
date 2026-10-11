package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "PRuvetSTAaYA\n", 5)
}

func TestSample2(t *testing.T) {
	runSample(t, "OYPROSTIYAOPECHATALSYAPRIVETSTASYA\n", 0)
}

func TestSample3(t *testing.T) {
	runSample(t, "helloworld\n", 0)
}
