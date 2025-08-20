package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	buf := new(strings.Builder)
	writer := bufio.NewWriter(buf)
	drive(reader, writer)
	writer.Flush()
	res := buf.String()
	res = strings.TrimSpace(res)
	expect = strings.TrimSpace(expect)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `hello,2,ok,0,bye,0,test,0,one,1,two,2,a,0,b,0`
	expect := `3
hello test one
ok bye two
a b`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `a,5,A,0,a,0,A,0,a,0,A,0`
	expect := `2
a
A a A a A`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `A,3,B,2,C,0,D,1,E,0,F,1,G,0,H,1,I,1,J,0,K,1,L,0,M,2,N,0,O,1,P,0`
	expect := `4
A K M
B F H L N O
C D G I P
E J`
	runSample(t, s, expect)
}
