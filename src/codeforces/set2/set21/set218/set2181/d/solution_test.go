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
	runSample(t, `2
2 2 11 3 2
3 4 12 1 1 2
`, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
2 0 7 2 4
1 4 9 4
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
5 104 293 12 36 27 97 5
5 51 485 24 18 20 73 288
4 574 753 27 8 11 119
10 280 657 21 12 12 13 22 21 89 9 14 147
9 40 997 564 13 29 203 38 13 12 45 27
7 476 727 39 19 69 32 66 6 16
`, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, `40
0 617 934
0 24 871
0 276 911
0 753 878
0 283 858
0 759 989
0 10 849
0 149 997
0 582 906
0 466 922
0 551 997
0 482 952
0 295 967
0 254 902
0 223 955
0 235 892
0 578 946
0 117 855
0 97 850
0 260 917
0 466 914
0 522 866
0 182 919
0 425 950
0 492 862
0 176 856
0 603 991
0 289 911
0 254 907
0 221 948
0 708 962
0 376 939
0 238 957
0 205 969
0 173 969
0 613 925
0 591 918
0 725 960
0 811 972
0 109 873
`, 38)
}
