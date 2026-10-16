package main

import (
	"bufio"
	"os"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %s, expect %v, but got %v", s, expect, res)
	}
}
func TestSample1(t *testing.T) {
	s := `1
2
1 4
1 2`
	expect := []int{1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1
2
2 3
1 2`
	expect := []int{2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1
3
1 4
2 3
1 2`
	expect := []int{2}
	runSample(t, s, expect)
}

func estSampleData(t *testing.T) {

	dir, _ := os.Getwd()

	dat, _ := os.Open(dir + "/sample/P11457_3.in")

	defer dat.Close()

	reader := bufio.NewReader(dat)

	res := process(reader)

	ans, _ := os.Open(dir + "/sample/P11457_3.out")
	defer ans.Close()

	ansReader := bufio.NewReader(ans)

	for i := 0; i < len(res); i++ {
		expect := readNum(ansReader)
		if res[i] != expect {
			t.Fatalf("Sample expect %d, but got %d[@%d]", expect, res[i], i)
		}
	}

}
