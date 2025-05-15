package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	ans := readNum(reader)
	if res != ans {
		t.Errorf("Sample expect %d, but got %d", ans, res)
	}
}

func TestSample1(t *testing.T) {
	s := `10 3
abab
bab
3	
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `841 7
qjqhrksmvedtqldrqgchhsofokfcovut
qhtmothoulodshrfejterjlguvooccsvqrrdfqfvkqhtecuhhuqhshthrkusrc
5
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `933 5
abaabdcbbabacbdddadbbb
babbadbaaadbbbbaabbaabccbbdbadbbbbbbdbcbdbaaadbdbdbbbbdcbbdcbdaadbd
15
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `901 8
kikjbkgkkjeeficiigjjidfhkfdckjdkkbkfhkhdcjidjbfdfkbhbfjeiffkfgcaigck
gcjifhjjfedbfbdhickjbkkghfhbigkeff
28
`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `6557487 3
aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
aaaaaaaaa
11172014
`
	runSample(t, s)
}

func TestSample6(t *testing.T) {
	s := `1263204 1
dcbabceabbaebddaaecbddaceaedacddadadcbbadbdccdecdacdcaadbceeddccbceaade
abddbccbdca
1894805
`
	runSample(t, s)
}

func TestSample7(t *testing.T) {
	s := `6387150 3
grlqvotdxnwkxqxvlcidkkgqxmluvqucmxeotdjdgooes
qkoeldlicvkdlxlqvqrnggl
425809
`
	runSample(t, s)
}

func TestSample8(t *testing.T) {
	s := `2308056 3
gabbfdehcegkmjicakkfhfahllalmkhhhjejfibdhgldjcbfmkfbmileigdmlhajfcciemfcbg
m
4616112
`
	runSample(t, s)
}
