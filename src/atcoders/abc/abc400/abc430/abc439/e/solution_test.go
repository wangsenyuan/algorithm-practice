package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
3 5
1 4
2 6
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1 2
1 3
1 4
1 5
1 6
`, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, `10
440423913 766294629
725560240 59187619
965580535 585990756
550925213 623321125
549392044 122410708
21524934 690874816
529970099 244587368
757265587 736247509
576136367 993115118
219853537 21553211
`, 4)
}
