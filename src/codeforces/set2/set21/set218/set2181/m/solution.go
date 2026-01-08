package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	tc := readNum(reader)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNum(reader *bufio.Reader) int {
	s := readString(reader)
	v, _ := strconv.Atoi(s)
	return v
}

func drive(reader *bufio.Reader) int {
	s := readString(reader)
	t := readString(reader)
	return solve(s, t)
}

func solve(x string, y string) int {
	// dp[i][0] 表示到i为止，c[i] % 2 == 0 时的最优解
	n := len(x)
	dp := []int{0, n}
	for i := range n {
		ndp := []int{n, n}

		for d := range 2 {
			for d2 := range 2 {
				cnt := dp[d]
				if (d+int(x[i]-'0'))%2 != d2 {
					cnt++
				}
				if d2 != int(y[i]-'0') {
					cnt++
				}
				ndp[d2] = min(ndp[d2], cnt)
			}
		}

		copy(dp, ndp)
	}

	return min(dp[0], dp[1])
}
