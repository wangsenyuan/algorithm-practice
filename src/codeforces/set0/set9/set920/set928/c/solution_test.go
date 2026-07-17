package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []proj) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res == nil {
		res = []proj{}
	}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
a 3
2
b 1
c 1

b 2
0

b 1
1
b 2

c 1
1
b 2
`, []proj{{"b", 1}, {"c", 1}})
}

func TestSample2(t *testing.T) {
	runSample(t, `9
codehorses 5
3
webfrmk 6
mashadb 1
mashadb 2

commons 2
0

mashadb 3
0

webfrmk 6
2
mashadb 3
commons 2

extra 4
1
extra 3

extra 3
0

extra 1
0

mashadb 1
1
extra 3

mashadb 2
1
extra 1
`, []proj{{"commons", 2}, {"extra", 1}, {"mashadb", 2}, {"webfrmk", 6}})
}

func TestSample3(t *testing.T) {
	runSample(t, `3
abc 1
2
abc 3
cba 2

abc 3
0

cba 2
0
`, []proj{{"cba", 2}})
}
