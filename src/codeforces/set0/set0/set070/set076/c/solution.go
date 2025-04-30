package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	n, k, TT := readThreeNums(reader)
	s := readString(reader)[:n]
	T := readNNums(reader, k)
	A := make([][]int, k)
	for i := range A {
		A[i] = readNNums(reader, k)
	}
	return solve(s, T, A, TT)
}

func solve(s string, T []int, A [][]int, TT int) int {
	n := len(s)
	k := len(T)
	dp := make([]int, k)
	for i := range k {
		dp[i] = -1
	}
	fp := make([]int, 1<<k)
	for i := range k {
		fp[1<<i] = T[i]
	}

	var all int

	for i := range n {
		x := int(s[i] - 'A')
		all |= 1 << x
		for j := range k {
			if dp[j] >= 0 {
				if (dp[j]>>j)&1 == 0 && (dp[j]>>x)&1 == 0 {
					fp[dp[j]] += A[j][x]
					fp[dp[j]|1<<j] -= A[j][x]
					fp[dp[j]|1<<x] -= A[j][x]
					fp[dp[j]|1<<j|1<<x] += A[j][x]
				}
				dp[j] |= 1 << x
			}
		}
		dp[x] = 0
	}

	for i := range k {
		for j := range 1 << k {
			if (j>>i)&1 == 1 {
				fp[j] += fp[j^(1<<i)]
			}
		}
	}
	var ans int

	for i := range 1 << k {
		if i&all == i && fp[i] <= TT && i != all {
			ans++
		}
	}
	return ans
}
