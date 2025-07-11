package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Printf("%.10f\n", res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) float64 {
	n := readNum(reader)
	x := readNNums(reader, n)
	v := readNNums(reader, n)
	return solve(x, v)
}

const eps = 1e-9

func solve(x []int, v []int) float64 {

	type pair struct {
		first  float64
		second float64
	}

	n := len(x)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{
			first:  float64(x[i]),
			second: float64(v[i]),
		}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Compare(a.first, b.first)
	})

	check := func(expect float64) bool {
		// 前面的往后面，后面的往前面
		hi := arr[0].first + expect*arr[0].second
		var l int
		for l < n && arr[l].first <= hi {
			hi = min(hi, arr[l].first+expect*arr[l].second)
			l++
		}

		lo := arr[n-1].first - expect*arr[n-1].second
		r := n - 1

		for r >= 0 && arr[r].first >= lo {
			lo = max(lo, arr[r].first-expect*arr[r].second)
			r--
		}

		return (lo < hi || math.Abs(lo-hi) < eps)
	}

	var lo, hi float64 = 0, 1e10

	for range 100 {
		mid := (lo + hi) / 2
		if check(mid) {
			hi = mid
		} else {
			lo = mid
		}
	}

	return (lo + hi) / 2
}
