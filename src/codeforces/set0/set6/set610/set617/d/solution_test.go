package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `-931665727 768789996
234859675 808326671
-931665727 879145023`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `-494824697 -964138793
-494824697 671151995
-24543485 877798954`
	expect := 2
	runSample(t, s, expect)
}
