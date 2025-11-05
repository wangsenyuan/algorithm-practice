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
	n, m := readTwoNums(reader)
	s := readString(reader)
	return solve(n, m, s)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(n int, m int, s string) int {
	k := n - m
	dp := make([][]int, k+1)
	for i := range k + 1 {
		dp[i] = make([]int, k+1)
	}

	dp[0][0] = 1

	for i := 1; i <= k; i++ {
		for bal := 0; bal < i; bal++ {
			// 如果这里使用左括号
			dp[i][bal+1] = add(dp[i][bal+1], dp[i-1][bal])
			if bal > 0 {
				dp[i][bal-1] = add(dp[i][bal-1], dp[i-1][bal])
			}
		}
	}

	var bal int
	var right_bal int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			bal--
		} else {
			bal++
		}
		right_bal = min(right_bal, bal)
	}

	bal = 0
	var left_bal int
	for i := range s {
		if s[i] == '(' {
			bal++
		} else {
			bal--
		}
		left_bal = min(left_bal, bal)
	}

	var ans int
	for x := 0; x <= k; x++ {
		// 前面x个，后面y个
		y := k - x
		for w := 0; w <= x; w++ {
			if w < -left_bal {
				// w必须保证比中间最大的右括号bal多
				continue
			}
			// w + bal + suf = 0
			v := (w + bal)
			if v > y || v < -right_bal {
				// 中间左括号多，所有后面需要更多的右括号
				continue
			}
			ans = add(ans, mul(dp[x][w], dp[y][v]))
		}
	}

	return ans
}
