package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	got := drive(reader)
	if got != expect {
		t.Fatalf("Sample mismatch\nexpect:\n%q\ngot:\n%q", expect, got)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
3 1 2 5 1
`, `     /\     
  /\/  \    
 /      \   
/        \  
          \/
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 5 1
`, `/\     
  \    
   \   
    \  
     \/
`)
}
