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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 4
1 2 3
2 4 5
1 3 4
3 4 7
`, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 3
1 2 1
2 3 2
3 4 4
`, 7)
}

func TestSample3(t *testing.T) {
	runSample(t, `7 10
1 2 726259430069220777
1 4 988687862609183408
1 5 298079271598409137
1 6 920499328385871537
1 7 763940148194103497
2 4 382710956291350101
3 4 770341659133285654
3 5 422036395078103425
3 6 472678770470637382
5 7 938201660808593198
`, 186751192333709144)
}
