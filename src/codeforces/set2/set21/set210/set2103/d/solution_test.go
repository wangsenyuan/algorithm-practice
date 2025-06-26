package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, res := process(reader)

	n := len(a)
	b := make([]int, n)
	id := 1
	flag := make([]bool, n)
	arr := make([]int, n)
	for i := range n {
		arr[i] = i
	}
	for len(arr) > 1 {
		for j, i := range arr {
			if id&1 == 1 {
				if j > 0 && res[i] > res[arr[j-1]] {
					flag[j] = true
				}
				if j+1 < len(arr) && res[i] > res[arr[j+1]] {
					flag[j] = true
				}
			} else {
				if j > 0 && res[i] < res[arr[j-1]] {
					flag[j] = true
				}
				if j+1 < len(arr) && res[i] < res[arr[j+1]] {
					flag[j] = true
				}
			}
		}

		var next []int
		for j := range arr {
			if flag[j] {
				b[arr[j]] = id
				flag[j] = false
			} else {
				next = append(next, arr[j])
			}
		}

		arr = next

		id++
	}

	b[arr[0]] = n

	if !reflect.DeepEqual(b, a) {
		t.Errorf("Sample expect %v, but got %v", a, b)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3
1 1 -1
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1 -1 1 2 1
	`)
}

func TestSample3(t *testing.T) {
	runSample(t, `8
3 1 2 1 -1 1 1 2
	`)
}

func TestSample4(t *testing.T) {
	runSample(t, `7
1 1 1 -1 1 1 1
	`)
}

func TestSample5(t *testing.T) {
	runSample(t, `5
1 1 1 1 -1
	`)
}

func TestSample6(t *testing.T) {
	runSample(t, `5
-1 1 1 1 1
	`)
}

func TestSample7(t *testing.T) {
	runSample(t, `5
-1 1 2 1 2
	`)
}

func TestSample8(t *testing.T) {
	runSample(t, `4
2 1 1 -1
	`)
}

func TestSample9(t *testing.T) {
	runSample(t, `4
1 1 1 -1
	`)
}
