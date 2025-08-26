package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	var buf bytes.Buffer
	for range tc {
		res := drive(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([][]int, 2)
	for i := 0; i < 2; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(a)
}

func solve(a [][]int) int {
	sum := make([]int, 2)
	cnt := make([]int, 2)
	n := len(a[0])

	for i := 0; i < n; i++ {
		if a[0][i] > a[1][i] {
			sum[0] += a[0][i]
		} else if a[0][i] < a[1][i] {
			sum[1] += a[1][i]
		} else {
			if a[0][i] == 1 {
				cnt[1]++
			} else if a[0][i] == -1 {
				cnt[0]++
			}
		}
	}
	if sum[0] > sum[1] {
		sum[0], sum[1] = sum[1], sum[0]
	}
	if sum[0]+cnt[1] <= sum[1]-cnt[0] {
		return sum[0] + cnt[1]
	}
	return (sum[0] + sum[1] + cnt[1] - cnt[0]) >> 1
}
