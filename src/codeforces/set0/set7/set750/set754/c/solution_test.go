package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	for _, x := range res {
		y := readString(reader)
		if x != y {
			t.Fatalf("Sample expect %s, but got %s", x, y)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2
Vladik netman
2
?: Hello, Vladik!
?: Hi
netman: Hello, Vladik!
Vladik: Hi
`
	runSample(t, s, true)
}

func TestSample2(t *testing.T) {
	s := `2
netman vladik
3
netman:how are you?
?:wrong message
vladik:im fine
`
	runSample(t, s, false)
}

func TestSample3(t *testing.T) {
	s := `3
netman vladik Fedosik
2
?: users are netman, vladik, Fedosik
vladik: something wrong with this chat
`
	runSample(t, s, false)
}

func TestSample4(t *testing.T) {
	s := `4
netman tigerrrrr banany2001 klinchuh
4
?: tigerrrrr, banany2001, klinchuh, my favourite team ever, are you ready?
klinchuh: yes, coach!
?: yes, netman
banany2001: yes of course.
netman: tigerrrrr, banany2001, klinchuh, my favourite team ever, are you ready?
klinchuh: yes, coach!
tigerrrrr: yes, netman
banany2001: yes of course.
`
	runSample(t, s, true)
}

func TestSample5(t *testing.T) {
	s := `7
SBkyKniF rR5X K3ddkoeg6 auvBIAv4ZG Q5yW19Zp Hg CdZoe0Hg
1
?:Hg!Q5yW19Zp!Q5yW19Zp CdZoe0Hg,auvBIAv4ZG,Q5yW19Zp.K3ddkoeg6,Q5yW19Zp auvBIAv4ZG?SBkyKniF,auvBIAv4ZG,
rR5X:Hg!Q5yW19Zp!Q5yW19Zp CdZoe0Hg,auvBIAv4ZG,Q5yW19Zp.K3ddkoeg6,Q5yW19Zp auvBIAv4ZG?SBkyKniF,auvBIAv4ZG,
`
	runSample(t, s, true)
}
