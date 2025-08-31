package p3664

import "testing"

func runSample(t *testing.T, cards []string, x byte, expect int) {
	res := score(cards, x)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	cards := []string{"ba", "ba"}
	expect := 0
	runSample(t, cards, 'b', expect)
}

func TestSample2(t *testing.T) {
	cards := []string{"aa", "ab", "ba", "ac"}
	expect := 2
	runSample(t, cards, 'a', expect)
}

func TestSample3(t *testing.T) {
	cards := []string{"ab", "aa", "ab", "bc", "cc", "bc", "bb", "ac", "bc", "bc", "aa", "aa", "ba", "bc", "cb", "ba", "ac", "bb", "cb", "ac", "cb", "cb", "ba", "bc", "ca", "ba", "bb", "cc", "cc", "ca", "ab", "bb", "bc", "ba", "ac", "bc", "ac", "ac", "bc", "bb", "bc", "ac", "bc", "aa", "ba", "cc", "ac", "bb", "ba", "bb"}
	expect := 16
	runSample(t, cards, 'b', expect)
}
