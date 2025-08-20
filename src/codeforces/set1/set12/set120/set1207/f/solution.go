package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	queries := make([][]int, n)
	for i := 0; i < n; i++ {
		queries[i] = readNNums(reader, 3)
	}

	res := solve(queries)
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}

	fmt.Print(buf.String())
}

const N = 500000

func solve(queries [][]int) []int {
	SN := int(math.Sqrt(float64(N/4))) + 1
	a := make([]int, N+1)
	dp := make([][]int, SN+1)
	for x := 1; x <= SN; x++ {
		dp[x] = make([]int, x)
	}

	add := func(i int, v int) {
		a[i] += v
		for x := 1; x <= SN; x++ {
			dp[x][i%x] += v
		}
	}

	find := func(x int, y int) int {
		if x > SN {
			var res int
			for i := y; i <= N; i += x {
				res += a[i]
			}
			return res
		}
		return dp[x][y]
	}

	var ans []int

	for _, cur := range queries {
		x, y := cur[1], cur[2]
		if cur[0] == 1 {
			add(x, y)
		} else {
			ans = append(ans, find(x, y))
		}
	}

	return ans
}
