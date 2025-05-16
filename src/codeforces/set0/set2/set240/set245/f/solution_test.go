package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}
func TestSample1(t *testing.T) {
	s := `60 3
2012-03-16 16:15:25: Disk size is
2012-03-16 16:15:25: Network failute
2012-03-16 16:16:29: Cant write varlog
2012-03-16 16:16:42: Unable to start process
2012-03-16 16:16:43: Disk size is too small
2012-03-16 16:16:53: Timeout detected
`
	expect := "2012-03-16 16:16:43"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 2
2012-03-16 23:59:59:Disk size
2012-03-17 00:00:00: Network
2012-03-17 00:00:01:Cant write varlog
`
	expect := "-1"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 2
2012-03-16 23:59:59:Disk size is too sm
2012-03-17 00:00:00:Network failute dete
2012-03-17 00:00:01:Cant write varlogmysq
`
	expect := "2012-03-17 00:00:00"
	runSample(t, s, expect)
}
