package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int64 {
	var n int
	var a, b string
	fmt.Fscan(reader, &n, &a, &b)
	return solve(a, b)
}

type item struct {
	diff  int
	ones  int
	zeros int
}

func solve(a string, b string) int64 {
	n := len(a)
	x := make([]item, n)
	y := make([]item, n)

	var ones int
	for i := range n {
		if a[i] == '1' {
			ones++
		}
		zeros := i + 1 - ones
		x[i] = item{2*ones - i - 1, ones, zeros}
	}

	ones = 0
	for i := range n {
		if b[i] == '1' {
			ones++
		}
		zeros := i + 1 - ones
		y[i] = item{2*ones - i - 1, ones, zeros}
	}

	sort.Slice(y, func(i, j int) bool {
		return y[i].diff < y[j].diff
	})

	sumOnes := make([]int64, n+1)
	sumZeros := make([]int64, n+1)
	for i, cur := range y {
		sumOnes[i+1] = sumOnes[i] + int64(cur.ones)
		sumZeros[i+1] = sumZeros[i] + int64(cur.zeros)
	}

	var ans int64
	for _, cur := range x {
		cnt := sort.Search(n, func(i int) bool {
			return y[i].diff > -cur.diff
		})
		ans += int64(cnt*cur.ones) + sumOnes[cnt]
		ans += int64((n-cnt)*cur.zeros) + sumZeros[n] - sumZeros[cnt]
	}

	return ans
}
