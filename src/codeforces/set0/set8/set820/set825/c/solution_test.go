package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
2 1 9`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 20
10 3 6 3
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `100 10
246 286 693 607 87 612 909 312 621 37 801 558 504 914 416 762 187 974 976 123 635 488 416 659 988 998 93 662 92 749 889 78 214 786 735 625 921 372 713 617 975 119 402 411 878 138 548 905 802 762 940 336 529 373 745 835 805 880 816 94 166 114 475 699 974 462 61 337 555 805 968 815 392 746 591 558 740 380 668 29 881 151 387 986 174 923 541 520 998 947 535 651 103 584 664 854 180 852 726 93
`
	expect := 1
	runSample(t, s, expect)
}
