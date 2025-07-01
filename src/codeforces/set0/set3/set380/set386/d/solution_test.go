package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans, res, pos, grid := process(reader)
	expect := readNum(reader)
	if ans != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, ans)
	}
	if ans < 0 {
		return
	}
	a, b, c := pos[0]-1, pos[1]-1, pos[2]-1
	for i := range res {
		u, v := res[i][0]-1, res[i][1]-1
		if u == a {
			if grid[u][v] != grid[b][c] {
				t.Fatalf("Sample result %v, not correct", res)
			}
			a = v
		} else if u == b {
			if grid[u][v] != grid[a][c] {
				t.Fatalf("Sample result %v, not correct", res)
			}
			b = v
		} else if u == c {
			if grid[u][v] != grid[a][b] {
				t.Fatalf("Sample result %v, not correct", res)
			}
			c = v
		} else {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}

	k := newKey(a, b, c)
	if k.a != 0 || k.b != 1 || k.c != 2 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
2 3 4
*aba
a*ab
ba*b
abb*
1
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
2 3 4
*abc
a*ab
ba*b
cbb*
-1
	`)
}

func TestSample3(t *testing.T) {
	runSample(t, `10
1 4 7
*cbcbbaacd
c*bbcaaddd
bb*ababdcc
cba*aabcdb
bcba*cdaac
baaac*caab
aabbdc*cab
addcaac*bb
cdcdaaab*b
ddcbcbbbb*
3
	`)
}

func TestSample4(t *testing.T) {
	runSample(t, `10
6 9 5
*cabaccbbc
c*bcccbcac
ab*bacaaca
bcb*caccba
acac*caccb
cccac*ccac
cbacac*ccb
bcacccc*cb
bacbcacc*a
ccaabcbba*
3
	`)
}

