package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2
1 1 1 1
1 1 1 1
3
home
1
car
1
brother
1
`
	expect := [][]string{
		{"home", "car"},
		{"brother"},
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 4
1 2 2 1
2 3 2 2
4
armchair
3
quetzalcoatl
10
pilotage
5
defibrillator
7
`
	expect := [][]string{
		{"armchair", "quetzalcoatl"},
		{"pilotage", "defibrillator"},
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 10
7 4 2 1
3 8 4 3
10 1 8 2
1 10 8 1
7 2 8 5
20
awhiwqom
97
bupsrkw
67
cdcle
46
dbw
20
hypayjdabm
96
igclyxfkshgbhrf
78
ldcyjtkldc
73
mgrdcybqbvor
49
njzddrjxafjrrw
63
ozucivhhggjtfkbojuqw
11
qgzsrwl
95
rz
28
tperpgd
66
vilawpd
16
vnujx
8
wadgiywkj
23
ydisqmwtprbiexkxy
16
yejvg
29
yusivpcoucy
43
zdnpcaybyupwnrqjln
69
`
	expect := [][]string{
		{"ydisqmwtprbiexkxy", "ldcyjtkldc"},
		{"dbw", "tperpgd", "qgzsrwl", "awhiwqom"},
		{"vilawpd", "vnujx", "rz", "mgrdcybqbvor", "bupsrkw", "zdnpcaybyupwnrqjln", "hypayjdabm"},
		{"cdcle", "njzddrjxafjrrw"},
		{"ozucivhhggjtfkbojuqw", "yejvg", "wadgiywkj", "yusivpcoucy", "igclyxfkshgbhrf"},
	}
	runSample(t, s, expect)
}
