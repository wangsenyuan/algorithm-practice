package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	for _, cur := range res {
		if len(cur) != 4 {
			t.Fatalf("Sample expect 4 teams, but got %d", len(cur))
		}
		// ignore
		readString(reader)
		for _, x := range cur {
			y := readString(reader)
			if x != y {
				t.Fatalf("Sample expect %s, but got %s", x, y)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `8
1 3 1 7
Barcelona 158
Milan 90
Spartak 46
Anderlecht 48
Celtic 32
Benfica 87
Zenit 79
Malaga 16
Group A:
Barcelona
Benfica
Spartak
Celtic
Group B:
Milan
Zenit
Anderlecht
Malaga
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `16
957 624 660 738
BHTXvn 379
BzAcZuVA 3
sadCjjoxTbLK 367
HPb 319
zMXuyAKOOwPASEBPgdHT 266
sOxtRopt 782
McusDFgpzhV 13
EaEQYJeFNxURmNCxj 570
wHlWHmGdIQTFhT 942
Ssoam 656
ibcNogJowwwzYnpOqvj 869
JRgUxqAAS 385
KXUjcPfXlkCk 136
BaKvykbVHvJfuMG 479
Bhp 987
TZijlaAWwsxjI 178
Group A:
Bhp
BaKvykbVHvJfuMG
BHTXvn
TZijlaAWwsxjI
Group B:
wHlWHmGdIQTFhT
Ssoam
sadCjjoxTbLK
KXUjcPfXlkCk
Group C:
ibcNogJowwwzYnpOqvj
EaEQYJeFNxURmNCxj
HPb
McusDFgpzhV
Group D:
sOxtRopt
JRgUxqAAS
zMXuyAKOOwPASEBPgdHT
BzAcZuVA
`
	runSample(t, s)
}
