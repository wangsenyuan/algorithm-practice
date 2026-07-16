package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `snow affects sports such as skiing, snowboarding, and snowmachine travel.
snowboarding is a recreational activity and olympic and paralympic sport.
`, 141)
}

func TestSample2(t *testing.T) {
	runSample(t, `'co-co-co, codeforces?!'
`, 25)
}

func TestSample3(t *testing.T) {
	runSample(t, `thun-thun-thunder, thunder, thunder
thunder, thun-, thunder
thun-thun-thunder, thunder
thunder, feel the thunder
lightning then the thunder
thunder, feel the thunder
lightning then the thunder
thunder, thunder
`, 183)
}

func TestRepeatedWordStillUnique(t *testing.T) {
	runSample(t, "abc abc abc\n", 10)
}

func TestHyphenSeparatesWords(t *testing.T) {
	runSample(t, "abc-abc\n", 7)
}

func TestApostropheSeparatesWords(t *testing.T) {
	runSample(t, "abc'abc\n", 7)
}
