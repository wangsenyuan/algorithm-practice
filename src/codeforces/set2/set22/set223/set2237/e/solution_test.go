package main

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expectOk bool, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	var tc int
	fmt.Fscan(reader, &tc)
	ok, res := drive(reader)
	if ok != expectOk {
		t.Fatalf("Sample expect ok=%v, but got ok=%v (res=%v)", expectOk, ok, res)
	}
	if !expectOk {
		return
	}
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
3
2 3 1
-1 -1 -1
`, true, []int{1, 2, 3})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
4
2 1 4 3
-1 -1 4 -1
`, true, []int{1, 2, 4, 3})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
4
2 1 4 3
3 1 -1 -1
`, false, nil)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
4
2 1 4 3
1 -1 -1 2
`, false, nil)
}

func TestSample5(t *testing.T) {
	runSample(t, `1
5
2 3 1 5 4
2 -1 -1 -1 -1
`, true, []int{2, 3, 1, 4, 5})
}

func TestSample6(t *testing.T) {
	runSample(t, `1
5
2 3 1 5 4
4 -1 -1 -1 -1
`, false, nil)
}

func TestSample7(t *testing.T) {
	runSample(t, `1
6
2 3 1 5 6 4
4 -1 -1 -1 -1 -1
`, true, []int{4, 5, 6, 1, 2, 3})
}

func TestSample8(t *testing.T) {
	runSample(t, `1
6
2 1 4 3 6 5
-1 3 -1 -1 -1 -1
`, true, []int{4, 3, 1, 2, 5, 6})
}

func TestSample9(t *testing.T) {
	runSample(t, `1
6
3 5 6 2 1 4
-1 -1 -1 3 6 -1
`, false, nil)
}

func TestSample10(t *testing.T) {
	runSample(t, `1
7
2 3 1 5 4 6 7
-1 -1 -1 -1 -1 7 -1
`, true, []int{1, 2, 3, 4, 5, 7, 6})
}

func TestSample11(t *testing.T) {
	runSample(t, `1
8
2 3 4 1 6 7 8 5
5 7 -1 -1 -1 -1 -1 -1
`, false, nil)
}

func TestSample12(t *testing.T) {
	runSample(t, `1
8
2 3 4 1 6 7 8 5
5 -1 -1 -1 -1 -1 -1 -1
`, true, []int{5, 6, 7, 8, 1, 2, 3, 4})
}
