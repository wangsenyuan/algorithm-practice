package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3 2
2 2
14 6 9
4 9 20
17 15 7
`, 28)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 4 1
1 1
5 10 1 1
10 1 1 1
1 1 1 1
`, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, `8 10 2
1 5
388 130 971 202 487 924 247 286 237 316
117 166 918 106 336 928 493 391 235 398
124 280 425 955 212 988 227 222 307 226
336 302 478 246 950 368 291 236 170 101
370 200 204 141 287 410 388 314 205 460
291 104 348 337 404 399 416 263 415 339
105 420 302 334 231 481 466 366 401 452
119 432 292 403 371 417 351 231 482 184
`, 1343)
}
