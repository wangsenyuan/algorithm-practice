package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) bool {
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(k, a)
}

func solve(k int, a []int) bool {
	n := len(a)
	arr := make([]int, n)

	suf := make([]int, n+1)
	next := make([]int, n)
	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, 2)
	}
	check := func(a []int) bool {
		for i, v := range a {
			if v <= k {
				arr[i] = -1
			} else {
				arr[i] = 1
			}
		}

		dp[n][0] = n
		dp[n][1] = n

		for i := n - 1; i >= 0; i-- {
			next[i] = n
			suf[i] = arr[i]
			if i+1 < n {
				suf[i] += suf[i+1]
				next[i] = next[i+1]
			}
			if suf[i] < 0 || suf[i] == 0 && (n-i)&1 == 0 {
				next[i] = i
			}
			copy(dp[i], dp[i+1])
			if suf[i] > suf[dp[i][(n-i)&1]] {
				dp[i][(n-i)&1] = i
			}
		}

		var sum int

		for l := 0; l < n; l++ {
			sum += arr[l]
			if l&1 == 0 && sum < 0 || l&1 == 1 && sum <= 0 {
				if l+2 < n && next[l+2] < n {
					return true
				}
				p := (n - (l + 1)) & 1
				r0 := dp[l+1][p]
				r1 := dp[l+1][p^1]
				if l+1 < r0 && r0 < n && suf[l+1]-suf[r0] <= 0 {
					return true
				}
				if l+1 < r1 && r1 < n && suf[l+1]-suf[r1] < 0 {
					return true
				}
			}
		}

		return false
	}

	if check(a) {
		return true
	}
	reverse(a)

	return check(a)
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
