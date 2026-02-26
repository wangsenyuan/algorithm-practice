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

	writer.Flush()

	res := buf.String()
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `aaaaa
2
a
a
`
	expect := `1
1
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `abacaba
3
ac
ba
a
`
	expect := `1
2
4
`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `abcabccbacba
10
cab
cba
c
a
b
ab
bc
ac
z
zasdx
`
	expect := `1
1
3
4
4
4
3
2
0
0
`
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `ygkpypcmqcdyobghazieayxxaoiukrmqrbjepkhswwqhrybvdwwysisnaxjqcuqbqdnmbimwsvgactidnhogjfpfbgsawmynbjgw
10
pf
swy
pcy
ws
awm
cyp
ydo
imw
rhy
wqh
`
	expect := `1
1
1
2
1
1
1
1
1
1
`
	runSample(t, s, expect)
}
