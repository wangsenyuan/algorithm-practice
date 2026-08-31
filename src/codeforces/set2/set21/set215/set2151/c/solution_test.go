package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int64) {
	t.Helper()

	reader := bufio.NewReader(strings.NewReader(input))
	if res := drive(reader); !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
32 78
`, []int64{46})
}

func TestSample2(t *testing.T) {
	runSample(t, `2
4 5 6 9
`, []int64{4, 6})
}

func TestSample3(t *testing.T) {
	runSample(t, `4
6149048 26582657 36124499 43993239 813829899 860114890 910238130 913669539
`, []int64{78018749, 1737022233, 1845329695, 3385003015})
}

func TestAlternatingLargeGaps(t *testing.T) {
	runSample(t, `3
0 1 10 11 20 21
`, []int64{3, 39, 41})
}
