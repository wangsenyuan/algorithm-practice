package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%.8f\n", process(reader))
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
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(a, k)
}

const eps = 1e-8

func solve(a []int, k int) float64 {
	sort.Ints(a)
	n := len(a)
	f := 1.0 - float64(k)/100
	check := func(x float64) bool {
		// x是希望达到的量
		sum := []float64{0, 0}
		var i int
		for i < n && float64(a[i]) < x {
			sum[0] += x - float64(a[i])
			i++
		}

		for i < n {
			sum[1] += float64(a[i]) - x
			i++
		}

		return sum[1]*f >= sum[0] || math.Abs(sum[1]*f-sum[0]) < eps
	}
	var l, r float64 = 0, float64(slices.Max(a))

	for range 100 {
		mid := (l + r) / 2
		if check(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	return (l + r) / 2
}
