package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
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

func drive(reader *bufio.Reader) int {
	firstLine := readNNums(reader, 4)
	a, b, T := firstLine[1], firstLine[2], firstLine[3]
	s := readString(reader)
	return solve(a, b, T, s)
}

func solve(a int, b int, T int, s string) int {
	n := len(s)
	// s2 := s + s
	n2 := 2 * n
	dp := make([]int, n2)

	for i := range n2 {
		if s[i%n] == 'w' {
			dp[i] = b
		}
		if i > 0 {
			dp[i] += dp[i-1]
		}
	}

	check := func(x int) bool {
		if x > n {
			return false
		}
		// 能够观看x个照片
		for l := n; l > 0 && l+x-1 >= n; l-- {
			sum := dp[l+x-1] - dp[l-1]
			// 先移动到l，然后再移动到 l + x - 1的位置
			sum += a * (n - l)
			sum += a * (l + x - 1 - l)
			if sum <= T-x {
				return true
			}
		}

		for r := n; r < n2 && r-x+1 <= n; r++ {
			// 先移动到r, 再移动到 r - x + 1的位置
			sum := dp[r] - dp[r-x]
			sum += a * (r - n)
			sum += a * (r - (r - x + 1))
			if sum <= T-x {
				return true
			}
		}

		return false
	}

	l, r := 1, n+1

	for l < r {
		mid := (l + r) >> 1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return r - 1
}
