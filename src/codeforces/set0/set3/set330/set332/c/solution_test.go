package main

import (
	"bufio"
	"cmp"
	"reflect"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res, orders, k := process(reader)

	expect := readNNums(reader, len(res))

	check := func(arr []int) []int {
		var tmp [][]int
		for _, i := range arr {
			tmp = append(tmp, orders[i-1])
		}
		slices.SortFunc(tmp, func(u []int, v []int) int {
			return cmp.Or(u[1]-v[1], v[0]-u[0])
		})
		res := make([]int, 2)
		for i := range arr {
			if i < len(arr)-k {
				res[1] += tmp[i][1]
			} else {
				res[0] += tmp[i][0]
			}
		}

		return res
	}

	x := check(expect)
	y := check(res)

	if !reflect.DeepEqual(x, y) {
		t.Errorf("Sample expect %v(%v), but got %v(%v)", expect, x, res, y)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 3 2
5 6
5 8
1 3
4 3
4 11
3 1 2`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 3 3
10 18
18 17
10 20
20 18
20 18
2 4 5`)
}

func TestSample3(t *testing.T) {
	runSample(t, `10 7 4
4 3
5 3
5 5
4 3
4 5
3 5
4 5
4 4
3 5
4 5
1 4 8 3 5 7 10`)
}
