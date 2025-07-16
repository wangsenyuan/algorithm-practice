package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	x, g, res := process(reader)
	expect := readString(reader)
	if expect == "YES" != (len(res) > 0) {
		t.Fatalf("Sample expect %s, but got %v", expect, res)
	}
	if expect == "NO" {
		return
	}
	var buf bytes.Buffer
	for _, i := range res {
		buf.WriteString(g[i-1])
	}
	y := buf.String()

	y += y

	if !strings.Contains(y, x) {
		t.Fatalf("Sample expect %s, but got %s", x, y)
	}
}

func TestSample1(t *testing.T) {
	s := `3 1
abc
4
b
a
c
d
YES`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4 2
aabbccdd
4
dd
ab
bc
cd
NO`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `10 10
lgfrjgityzwtmfyygmpsmokiwphewhpoelsvnctwxmpimqvblgrisozncsidqlqzovrzlgvneovgvkjoxvprlqhlgokaooflsiih
20
gvkjoxvprl
ihlgfrjgit
ygmpsmokiw
phewhpoels
qhlgokaoof
rzlgvneovg
gityzwtmfy
vnctwxmpim
mqvblgriso
ncsidqlqzo
vrzlgvneov
yzwtmfyygm
mpsmokiwph
ityzwtmfyy
ctwxmpimqv
zovrzlgvne
nctwxmpimq
qvblgrisoz
lsiihlgfrj
zncsidqlqz
YES`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `2 1
aa
2
a
b
NO`
	runSample(t, s)
}
