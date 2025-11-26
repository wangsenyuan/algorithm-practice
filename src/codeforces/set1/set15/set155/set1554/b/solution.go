package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		a := readNNums(reader, n)
		res := solve(a, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

const inf = 1 << 60

type pair struct {
	first  int
	second int
}

func solve(a []int, k int) int {
	n := len(a)

	S := 2 * n

	// dp[s] = (i, j) where s contains a[i] and a[j]
	dp := make([]pair, S+1)

	for i, v := range a {
		i++
		dp[v].second = dp[v].first
		dp[v].first = i
	}

	merge := func(a pair, b pair) pair {
		if b.first > a.first {
			a.second = a.first
			a.first = b.first
		} else if b.first > a.second && b.first < a.first {
			a.second = b.first
		}

		if b.second > a.second && b.second < a.first {
			a.second = b.second
		}

		return a
	}

	H := bits.Len(uint(S))

	best := -inf

	for s := 1; s <= S; s++ {
		for i := range H {
			if (s>>i)&1 == 1 {
				dp[s] = merge(dp[s], dp[s^(1<<i)])
			}
		}
		if dp[s].first > 0 && dp[s].second > 0 {
			i, j := dp[s].first, dp[s].second
			tmp := i*j - k*(a[i-1]|a[j-1])
			best = max(best, tmp)
		}
	}
	return best
}

func solve1(a []int, k int) int {
	n := len(a)

	s := max(0, n-2*k)

	res := -inf
	for i := s; i < n; i++ {
		for j := i + 1; j < n; j++ {
			tmp := (i+1)*(j+1) - k*(a[i]|a[j])
			res = max(res, tmp)
		}
	}

	return res
}
