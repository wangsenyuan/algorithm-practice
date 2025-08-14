package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}
	ans := buf.String()
	ans = strings.TrimSpace(ans)
	expect = strings.TrimSpace(expect)

	if ans != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `12
Widget me(50,40)
VBox grandpa
HBox father
grandpa.pack(father)
father.pack(me)
grandpa.set_border(10)
grandpa.set_spacing(20)
Widget brother(30,60)
father.pack(brother)
Widget friend(20,60)
Widget uncle(100,20)
grandpa.pack(uncle)
`
	expect := `brother 30 60
father 80 60
friend 20 60
grandpa 120 120
me 50 40
uncle 100 20
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `15
Widget pack(10,10)
HBox dummy
HBox x
VBox y
y.pack(dummy)
y.set_border(5)
y.set_spacing(55)
dummy.set_border(10)
dummy.set_spacing(20)
x.set_border(10)
x.set_spacing(10)
x.pack(pack)
x.pack(dummy)
x.pack(pack)
x.set_border(0)
`
	expect := `dummy 0 0
pack 10 10
x 40 10
y 10 10
`
	runSample(t, s, expect)
}
