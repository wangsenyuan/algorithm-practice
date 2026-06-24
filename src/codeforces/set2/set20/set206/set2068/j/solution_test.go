package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	var tc int
	fmt.Fscan(reader, &tc)
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %q, but got %q", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
4
WRRWWWRR
`, "YES")
}

func TestSample2(t *testing.T) {
	runSample(t, `1
1
WR
`, "NO")
}

func TestSample3(t *testing.T) {
	runSample(t, `1
20
WWWWRRWRRRRRWRRWRWRRWRRWWWWWWWRWWRWWRRRR
`, "YES")
}

func TestSample4(t *testing.T) {
	runSample(t, `1
5
RWWRRWWRRW
`, "NO")
}
